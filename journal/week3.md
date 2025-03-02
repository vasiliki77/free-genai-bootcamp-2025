
## Table of Contents

- [Table of Contents](#table-of-contents)
- [English to Ancient Greek Translation Microservice](#english-to-ancient-greek-translation-microservice)
  - [1. Create the environment variables script:](#1-create-the-environment-variables-script)
  - [2. Create a Dockerfile for the service:](#2-create-a-dockerfile-for-the-service)
  - [3. Create a Python microservice file (opea\_translation\_microservice.py):](#3-create-a-python-microservice-file-opea_translation_microservicepy)
  - [4. Create a docker-compose.yaml file:](#4-create-a-docker-composeyaml-file)
  - [5. Create a deployment script (similar to the Text-to-SQL one):](#5-create-a-deployment-script-similar-to-the-text-to-sql-one)
  - [6. Create a requirements.txt file for the service:](#6-create-a-requirementstxt-file-for-the-service)
  - [7. Directory Structure](#7-directory-structure)
  - [8. Usage](#8-usage)
  - [9. Tests and conclusion](#9-tests-and-conclusion)
- [New approach](#new-approach)
- [Conclusion](#conclusion)

## English to Ancient Greek Translation Microservice

Starting with the https://github.com/opea-project/GenAIComps. Going backwards though. I know that the best model for translation to ancient Greek is Mistral-7B-Instruct-v0.3 so I am trying to find examples of where my model has been used to see which component is a good match.
One of the examples that is using the model is [text2sql](https://github.com/opea-project/GenAIComps/tree/main/comps/text2sql) so I will try to create a similar mega-service.

I'll need to:

1. Set up a similar infrastructure (using Docker)
2. Use the same Mistral model
3. Create a translation-specific API endpoint
4. Modify the prompt engineering to handle translation instead of SQL generation


### 1. Create the environment variables script:

```bash
#!/bin/bash
# Set environment variables for English to Ancient Greek Translation service

export TAG='comps'
export TRANSLATION_PORT=11700
export LLM_ENDPOINT_PORT=11710
export LLM_MODEL_ID="mistralai/Mistral-7B-Instruct-v0.3"
export HUGGINGFACEHUB_API_TOKEN=${HUGGINGFACEHUB_API_TOKEN}
export service_name="eng2greek"
```

### 2. Create a Dockerfile for the service:

```dockerfile
FROM python:3.10-slim

WORKDIR /app

# Copy necessary files
COPY comps/eng2greek/src/requirements.txt /app/
COPY comps/eng2greek/src/ /app/

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Expose the port
EXPOSE 8080

# Command to run the service
CMD ["python", "opea_translation_microservice.py"]
```

### 3. Create a Python microservice file (opea_translation_microservice.py):

```python
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import uvicorn
import os
from langchain.llms import HuggingFaceTextGenInference
from langchain.prompts import PromptTemplate

app = FastAPI()

# Configure LLM endpoint
llm_endpoint_url = os.environ.get("llm_endpoint_url")
if not llm_endpoint_url:
    llm_endpoint_url = "http://localhost:8008"

# Initialize LLM
llm = HuggingFaceTextGenInference(
    inference_server_url=llm_endpoint_url,
    max_new_tokens=512,
    temperature=0.1,
    timeout=120,
)

# Translation prompt template
TRANSLATION_TEMPLATE = """You are an expert translator specializing in Ancient Greek. 
Translate the following English text into Ancient Greek (Koine Greek).
Make sure to use proper Ancient Greek grammar, vocabulary, and sentence structure.

English text: {input_text}

Ancient Greek translation:"""

translation_prompt = PromptTemplate(
    template=TRANSLATION_TEMPLATE,
    input_variables=["input_text"],
)

class TranslationRequest(BaseModel):
    input_text: str

class TranslationResponse(BaseModel):
    input_text: str
    output: str

@app.get("/health")
def health_check():
    return {"status": "healthy"}

@app.post("/v1/translate", response_model=TranslationResponse)
def translate_text(request: TranslationRequest):
    try:
        # Format the prompt
        prompt = translation_prompt.format(input_text=request.input_text)
        
        # Get translation from LLM
        result = llm(prompt)
        
        # Return the result
        return TranslationResponse(
            input_text=request.input_text,
            output=result
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)
```

### 4. Create a docker-compose.yaml file:

```yaml
version: '3'

services:
  tgi-server:
    image: ghcr.io/huggingface/text-generation-inference:2.1.0
    container_name: tgi-server
    ports:
      - "${LLM_ENDPOINT_PORT}:80"
    environment:
      - HF_TOKEN=${HUGGINGFACEHUB_API_TOKEN}
      - model=${LLM_MODEL_ID}
    command: --model-id ${LLM_MODEL_ID}
    volumes:
      - ../../../../data:/data
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: 1
              capabilities: [gpu]

  eng2greek:
    image: opea/eng2greek:${TAG}
    container_name: eng2greek-server
    depends_on:
      - tgi-server
    ports:
      - "${TRANSLATION_PORT}:8080"
    environment:
      - llm_endpoint_url=${TGI_LLM_ENDPOINT}
```

### 5. Create a deployment script (similar to the Text-to-SQL one):

```bash
#!/bin/bash
# Deployment script for English to Ancient Greek translation service

set -x

WORKPATH=$(dirname "$PWD")
LOG_PATH="$WORKPATH/tests"
ip_address=$(hostname -I | awk '{print $1}')
DATA_PATH=$WORKPATH/data

export TAG='comps'
export TRANSLATION_PORT=11700
export LLM_ENDPOINT_PORT=11710
export LLM_MODEL_ID="mistralai/Mistral-7B-Instruct-v0.3"
export HUGGINGFACEHUB_API_TOKEN=${HUGGINGFACEHUB_API_TOKEN}
export service_name="eng2greek"

function build_docker_images() {
  cd $WORKPATH
  docker build --no-cache -t opea/eng2greek:$TAG -f comps/eng2greek/src/Dockerfile .
}

check_tgi_connection() {
  url=$1
  timeout=1200
  interval=10

  local start_time=$(date +%s)

  while true; do
    if curl --silent --head --fail "$url" > /dev/null; then
      echo "Success"
      return 0
    fi
    echo
    local current_time=$(date +%s)
    local elapsed_time=$((current_time - start_time))

    if [ "$elapsed_time" -ge "$timeout" ]; then
      echo "Timeout,$((timeout / 60))min can't connect $url"
      return 1
    fi
    echo "Waiting for service for $elapsed_time seconds"
    sleep "$interval"
  done
}

function start_service() {
  export TGI_LLM_ENDPOINT="http://${ip_address}:${LLM_ENDPOINT_PORT}"
  unset http_proxy

  cd $WORKPATH/comps/eng2greek/deployment/docker_compose
  docker compose -f compose.yaml up ${service_name} -d > ${LOG_PATH}/start_services_with_compose.log
  check_tgi_connection "${TGI_LLM_ENDPOINT}/health"
}

function validate_microservice() {
  result=$(http_proxy="" curl http://${ip_address}:${TRANSLATION_PORT}/v1/translate \
  -X POST \
  -d '{"input_text": "Hello, how are you today?"}' \
  -H 'Content-Type: application/json')

  if [[ $result == *"output"* ]]; then
    echo $result
    echo "Result correct."
  else
    echo "Result wrong. Received was $result"
    docker logs eng2greek-server > ${LOG_PATH}/eng2greek.log
    docker logs tgi-server > ${LOG_PATH}/tgi.log
    exit 1
  fi
}

function stop_docker() {
  cd $WORKPATH/comps/eng2greek/deployment/docker_compose
  docker compose -f compose.yaml down ${service_name} --remove-orphans
}

function main() {
  stop_docker
  build_docker_images
  start_service
  validate_microservice
  stop_docker
  echo y | docker system prune
}

main
```

### 6. Create a requirements.txt file for the service:

```
fastapi==0.104.1
uvicorn==0.24.0
langchain==0.0.335
requests==2.31.0
pydantic==2.4.2
```

### 7. Directory Structure

Set up your project with this directory structure:

```
comps/
└── eng2greek/
    ├── src/
    │   ├── Dockerfile
    │   ├── requirements.txt
    │   └── opea_translation_microservice.py
    └── deployment/
        └── docker_compose/
            └── compose.yaml
tests/
data/
```

### 8. Usage

Once deployed, test the translation service with:

```bash
curl http://${your_ip}:${TRANSLATION_PORT}/v1/translate \
  -X POST \
  -d '{"input_text": "Wisdom is the most valuable possession."}' \
  -H 'Content-Type: application/json'
```

This implementation follows the same pattern as the Text-to-SQL service but is focused on translating English to Ancient Greek. The service utilizes the Mistral-7B model, but the prompt and API are customized for translation tasks.

### 9. Tests and conclusion

There are several issues with the current setup but the main problem was that the TGI server is having trouble initializing with the Mistral-7B model and it's stuck in the "warming up"phase for over an hour.


> 02/03/2025

## New approach

1. Use a smaller, faster model instead of Mistral-7B (which seems to be causing startup issues)
2. Fix token authentication (create proper token file)
3. Fix the healthcheck configuration
4. Ensure the network setup is correct
5. Update the compose file with proper parameters

For the model, an option like facebook/opt-125m or facebook/opt-350m would be more suitable for a CPU deployment.

## Conclusion

This solution was tested and the infrastructure worked, but the translation was very poor even when example prompts were provided. 
It seems that the particular model, although lightweight compared to Mistral and capable of answering simple questions, it is not adequately trained to translate to ancient Greek. 
My final thought is that, although I made the OPEA mega service work, I cannot use it for translation to ancient greek. 

For a guide on how to build the mega service, please see the [README.md](../opea-comps/mega-service/README.md)