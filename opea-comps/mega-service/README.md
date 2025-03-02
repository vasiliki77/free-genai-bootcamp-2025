### Prerequisites

1. **Docker**: Ensure Docker is installed and running on the system.
2. **Git**: Ensure Git is installed to clone the repository.

### Setup Instructions

1. **Clone the Repository**

   First, clone the repository to your local machine:

   ```bash
   git clone git@github.com:vasiliki77/free-genai-bootcamp-2025.git
   cd free-genai-bootcamp-2025/opea-comps/mega-service
   ```

2. **Set Up Environment Variables**

   Create a file named `set_env.sh` in the root of the `mega-service` directory with the following content:

   ```bash
   #!/bin/bash
   # Set environment variables for English to Ancient Greek Translation service

   export TAG='comps'
   export TRANSLATION_PORT=11700
   export LLM_ENDPOINT_PORT=11710
   export LLM_MODEL_ID="facebook/opt-350m"
   export HUGGINGFACEHUB_API_TOKEN="your_huggingface_api_token"
   export HF_TOKEN="${HUGGINGFACEHUB_API_TOKEN}"
   export service_name="eng2greek"
   export DATA_PATH="$(pwd)/data"
   export MAX_INPUT_TOKENS=512
   export MAX_TOTAL_TOKENS=1024
   export ip_address=$(hostname -I | awk '{print $1}')
   export host_ip="${ip_address}"
   ```

   **Note**: Replace `"your_huggingface_api_token"` with a valid Hugging Face API token.

3. **Create Necessary Directories and Token File**

   ```bash
   # Create data directory
   mkdir -p data

   # Create token file (without newline)
   echo -n "your_huggingface_api_token" > data/token
   chmod 600 data/token
   ```

4. **Create Docker Network**

   ```bash
   # Remove network if it exists
   docker network rm translation_network 2>/dev/null || true

   # Create network
   docker network create translation_network
   ```

5. **Build Docker Images**

   Navigate to the source directory and build the Docker image:

   ```bash
   cd eng2greek/src
   docker build -t opea/eng2greek:comps .
   ```

6. **Run the TGI Server**

   Start the TGI server with the specified model:

   ```bash
   docker run -d --name tgi-server \
     --network translation_network \
     -v $(pwd)/data:/data \
     -p 11710:80 \
     -e HF_TOKEN="your_huggingface_api_token" \
     -e HUGGING_FACE_HUB_TOKEN="your_huggingface_api_token" \
     ghcr.io/huggingface/text-generation-inference:2.4.0-intel-cpu \
     --model-id facebook/opt-350m \
     --max-input-tokens 512 \
     --max-total-tokens 1024 \
     --huggingface-hub-cache /data
   ```

This will take around 20 minutes. Check the logs:
```bash
docker logs -f tgi-server
```
It will be ready when you will see `text_generation_router::server: router/src/server.rs:2354: Connected`

7. **Run the eng2greek Service**

   Start the eng2greek service:

   ```bash
   docker run -d --name eng2greek-server \
     --network translation_network \
     -p 11700:8080 \
     -e llm_endpoint_url="http://tgi-server:80/generate" \
     opea/eng2greek:comps
   ```

8. **Verify Services are Running**

   Check if both containers are running:

   ```bash
   docker ps
   ```

9. **Test the Translation Service**

   Use the following command to test the translation service with example prompts:

   ```bash
   curl -X POST http://localhost:11700/v1/translate \
     -H 'Content-Type: application/json' \
     -d '{"input_text": "Hello, friend"}'
   ```

####
Example prompts with output
```bash
curl http://localhost:11710/generate \host:11710/generate \
  -X POST \
  -d '{
    "inputs": "Translate the following English sentence to Ancient Greek:\n1. \"Wisdom is a virtue.\" → Σοφία ἐστὶν ἀρετή.\n2. \"Life is short.\" → Ὁ βίος βραχύς ἐστιν.\n3. \"Knowledge is power.\" → Γνώσις ἐστὶν δύναμις.\n4. \"Hello\" →",
    "parameters": {"max_new_tokens": 20}
  }' \
  -H 'Content-Type: application/json'
{"generated_text":" Γνώσις ἐστιν.\n5. \""}
```
```bash
curl http://localhost:11710/generate \
  -X POST \
  -d '{"inputs":"Translate to Ancient Greek: Hello","parameters":{"max_new_tokens":10}}' \
  -H 'Content-Type: application/json'
{"generated_text":", my name is Tzadik.\n"}
```
I think it meant Tzatziki 😂😂😂😂

### Conclusion

Although the mega service now works, it's not good for translation to ancient Greek. I could use the mega service for a different use case, but not for this bootcamps business goal. 

The initial approach to create a mega service using tgi server and mistralAI was not successful as the model was taking more that 90 minutes to load. For more details see [week3.md](../../journal/week3.md)

### Additional Notes

- **Error Handling**: If any errors occur during the setup, check the Docker logs for more information:
  ```bash
  docker logs tgi-server
  docker logs eng2greek-server
  ```

