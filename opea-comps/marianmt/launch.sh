#!/bin/bash

# Build the container
docker build -t translation-service .

# Run the container
docker run -d \
    --name translation-service \
    -p 8000:8000 \
    -e HF_TOKEN=${HF_TOKEN} \
    translation-service 