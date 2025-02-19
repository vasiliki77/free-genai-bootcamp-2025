#!/bin/bash

# Copyright (C) 2024 Intel Corporation
# SPDX-License-Identifier: Apache-2.0


# Set default values



default_port=8009
default_model="mistralai/Mistral-7B-Instruct-v0.3"
default_device="gpu"
swap_space=50
image="opea/vllm-arc:latest"

while getopts ":hm:p:d:" opt; do
  case $opt in
    h)
      echo "Usage: $0 [-h] [-m model] [-p port] [-d device]"
      echo "Options:"
      echo "  -h         Display this help message"
      echo "  -m model   Model (default: meta-llama/Llama-2-7b-hf for cpu"
      echo "             meta-llama/Llama-3.2-3B-Instruct for gpu)"
      echo "  -p port    Port (default: 8000)"
      echo "  -d device  Target Device (Default: cpu, optional selection can be 'cpu' and 'gpu')"
      exit 0
      ;;
    m)
      model=$OPTARG
      ;;
    p)
      port=$OPTARG
      ;;
    d)
      device=$OPTARG
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done

# Assign arguments to variables
model_name=${model:-$default_model}
port_number=${port:-$default_port}
device=${device:-$default_device}


# Set the Huggingface cache directory variable
HF_CACHE_DIR=$HOME/.cache/huggingface
if [ "$device" = "gpu" ]; then
  docker_args="-e VLLM_OPENVINO_DEVICE=GPU  --device /dev/dri -v /dev/dri/by-path:/dev/dri/by-path"
  vllm_args="--max_model_len=1024"
  model_name="mistralai/Mistral-7B-Instruct-v0.3"
  image="opea/vllm-arc:latest"
fi
# Start the model server using Openvino as the backend inference engine.
# Provide the container name that is unique and meaningful, typically one that includes the model name.

docker run --rm --name="vllm-openvino-server-gpu" \
  -p 8009:80 \
  --ipc=host \
  --device /dev/dri:/dev/dri \
  -e VLLM_OPENVINO_DEVICE=GPU \
  -e VLLM_OPENVINO_ENABLE_QUANTIZED_WEIGHTS=ON \
  -e VLLM_OPENVINO_KVCACHE_SPACE=8 \
  -e HTTPS_PROXY=$https_proxy \
  -e HTTP_PROXY=$https_proxy \
  -e HF_TOKEN=${HF_TOKEN} \
  -e LD_LIBRARY_PATH=/opt/intel/openvino/runtime/lib/intel64:$LD_LIBRARY_PATH \
  -e PYTHONPATH=/opt/intel/openvino/python:$PYTHONPATH \
  -v $HOME/.cache/huggingface:/root/.cache/huggingface \
  opea/vllm-arc:latest /bin/bash -c "\
    python3 -c 'import openvino; print(\"OpenVINO version:\", openvino.__version__)' && \
    python3 -m vllm.entrypoints.openai.api_server \
      --model mistralai/Mistral-7B-Instruct-v0.3 \
      --host 0.0.0.0 \
      --port 80 \
      --max-model-len=8192 \
      --tensor-parallel-size=1"