

# **Building an OPEA Mega Service from Scratch**

## **1. Set Up Your Project Directory**
First, create a project folder and navigate into it:
```sh
mkdir mega_service_project
cd mega_service_project
```

Inside this folder, create the following subdirectories:
```sh
mkdir mega_service_new
mkdir mega_service_new/bin
```

---

## **2. Create the Main Mega Service File**
Create a Python file for your Mega Service inside `mega_service_new`:
```sh
touch mega_service_new/chat.py
```
Open `chat.py` and define the basic structure:
```python
import fastapi
from comps import ServiceOrchestrator, Microservice

class Chat:
    def __init__(self):
        self.service_orchestrator = ServiceOrchestrator()
        self.endpoint = "chat_service"

    def add_remote_services(self):
        pass

    def start(self):
        pass

if __name__ == "__main__":
    chat_service = Chat()
    chat_service.start()
```

---

## **3. Install Required Dependencies**
Create a `requirements.txt` file:
```sh
touch requirements.txt
```
Add the following dependencies:
```
fastapi
comps
uvicorn
```
Then install the dependencies:
```sh
pip install -r requirements.txt
```

---

## **4. Define the Mega Service**
Modify `chat.py` to define the Mega Service:

```python
import fastapi
from comps import ServiceOrchestrator, Microservice

class Chat:
    def __init__(self):
        self.service_orchestrator = ServiceOrchestrator()
        self.endpoint = "chat_service"
        self.port = 8888
        self.host = "0.0.0.0"

    def add_remote_services(self):
        # Define microservices
        embedding_service = Microservice(endpoint="v1/embedding")
        retriever_service = Microservice(endpoint="v1/retriever")
        rerank_service = Microservice(endpoint="v1/rerank")
        llm_service = Microservice(endpoint="v1/chat/completions")

        # Register microservices
        self.service_orchestrator.register_microservice(embedding_service)
        self.service_orchestrator.register_microservice(retriever_service)
        self.service_orchestrator.register_microservice(rerank_service)
        self.service_orchestrator.register_microservice(llm_service)

        # Define service execution flow
        self.service_orchestrator.flow_to(embedding_service, retriever_service)
        self.service_orchestrator.flow_to(retriever_service, rerank_service)
        self.service_orchestrator.flow_to(rerank_service, llm_service)

    def start(self):
        # Define the Mega Service entry point
        mega_service = Microservice(service_type="mega_service", endpoint=self.endpoint)

        # Attach a request handler
        @mega_service.app.post(f"/{self.endpoint}")
        async def handle_request(request: fastapi.Request):
            data = await request.json()
            runtime_graph = self.service_orchestrator.schedule(data)
            return runtime_graph

        # Start the Mega Service
        mega_service.start(self.host, self.port)

if __name__ == "__main__":
    chat_service = Chat()
    chat_service.add_remote_services()
    chat_service.start()
```

---

## **5. Create a Dockerfile**
Create a `Dockerfile`:
```sh
touch Dockerfile
```
Add the following content:
```dockerfile
FROM python:3.9

# Set working directory
WORKDIR /app

# Copy dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy application code
COPY . .

# Expose port
EXPOSE 8888

# Start application
CMD ["python", "mega_service_new/chat.py"]
```

---

## **6. Create a Docker Compose File**
Create a `docker-compose.yml`:
```sh
touch docker-compose.yml
```
Add the following content:
```yaml
version: "3.8"

services:
  mega_service:
    build: .
    container_name: mega_service
    ports:
      - "8888:8888"
    environment:
      - OPEA_ENV=production
    networks:
      - app_network

networks:
  app_network:
    driver: bridge
```

---

## **7. Run the Mega Service**
Build and run the Mega Service using Docker:
```sh
docker-compose up --build
```

---

## **8. Test the Mega Service**
Create a `bin/message` script to send a request:
```sh
touch bin/message
```
Add the following content:
```sh
#!/bin/bash
curl -X POST "http://localhost:8888/chat_service" \
     -H "Content-Type: application/json" \
     -d '{
         "messages": [{"role": "user", "content": "Hello, how are you?"}],
         "model": "llama-3-7b",
         "stream": false
     }'
```
Make it executable:
```sh
chmod +x bin/message
```
Run the script:
```sh
bin/message
```

---

## **9. Debugging and Troubleshooting**
If you run into errors:
- **Check logs**: `docker-compose logs -f`
- **Stop all running containers**: `docker-compose down`
- **Rebuild if necessary**: `docker-compose up --build`

---

### **Final Thoughts**
We now have a fully functional Mega Service built using OPEA. This service can orchestrate multiple microservices for tasks like embeddings, retrieval, re-ranking, and LLM processing.

