#!/bin/bash
# Set environment variables for English to Ancient Greek Translation service

# Get the directory where this script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export TAG='comps'
export TRANSLATION_PORT=11700
export LLM_ENDPOINT_PORT=11710
# Using a much smaller model that will initialize quickly on CPU
export LLM_MODEL_ID="facebook/opt-350m"
# Ensure no newlines in token
export HUGGINGFACEHUB_API_TOKEN="my-hf-token"
export HF_TOKEN="${HUGGINGFACEHUB_API_TOKEN}"
export service_name="eng2greek"
export DATA_PATH="${SCRIPT_DIR}/data"
# Ensure max_input_tokens is less than max_total_tokens to avoid the parameter error
export MAX_INPUT_TOKENS=512
export MAX_TOTAL_TOKENS=1024
export ip_address=$(hostname -I | awk '{print $1}')
export host_ip="${ip_address}"