## OPEA Components summary

1. Choosing vLLM Instead of TGI
Initially, I needed a serving solution for running large language models (LLMs) on my AI PC. Andrew planned to use TGI (Text Generation Inference) and then Ollama, but I wanted an alternative to work with MistralAI.

### Why vLLM?

- Supports OpenVINO, which is optimized for Intel CPUs and GPUs.
- More flexible than TGI in handling model parallelism.
- Efficient memory usage with techniques like PagedAttention.
- Can be containerized, making deployment easier.

### Why OpenVINO?

OpenVINO is Intel’s AI inference engine, optimized for Intel hardware (CPUs, Arc GPUs, NPUs).
Enables faster inference on Intel architecture without needing NVIDIA GPUs.
Compatible with vLLM, allowing us to use MistralAI efficiently on Intel hardware.

2. Setting Up Docker and WSL
Before running vLLM with OpenVINO, I needed to ensure Docker was correctly installed inside WSL.

- Installed Docker on WSL according to [the documentation](https://docs.docker.com/engine/install/ubuntu/)
```bash
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
```
- Verified Docker inside WSL (docker ps, docker info).

3. Running vLLM with OpenVINO
   
After setting up Docker, I followed the official vLLM OpenVINO guide from GenAIComps:

Cloned the GenAIComps repo:

```bash
git clone https://github.com/opea-project/GenAIComps.git
cd GenAIComps/comps/third_parties/vllm
```
Then copied `GenAIComps/comps/third_parties/vllm` into my `opea-comps` folder.

Exported the env parameters:
```bash
export LLM_ENDPOINT_PORT=8008
export host_ip=$(hostname -I | awk '{print $1}')
export HF_TOKEN="my-token"
export LLM_ENDPOINT="http://${host_ip}:${LLM_ENDPOINT_PORT}"
export LLM_MODEL_ID="mistralai/Mistral-7B-Instruct-v0.3"
```



```bash 
./build_docker_vllm_openvino.sh gpu
```
- Initially, I ran it on GPU but faced some issues.
- Later, I switched to CPU mode.
- Ran the container (GPU mode first, then CPU mode):


4. Debugging Container Issues
After launching the container, it kept exiting immediately. I troubleshooted by:

Checking running containers:
```bash
docker ps -a
```
Found that the container started and then stopped.

Checking logs for errors:
```bash
docker logs vllm-openvino-server
```
Saw warnings about OpenVINO IR model conversion:

`Provided model id Intel/neural-chat-7b-v3-3 does not contain OpenVINO IR`


5. Realizing I wasn’t Using MistralAI
   
Then I realised that although I had exported env parameters, the script had some default values. So I had to edit the script to use the env parameters.

```bash
curl http://172.25.72.192:8009/v1/chat/completions \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"model": "mistralai/Mistral-7B-Instruct-v0.3", "messages": [{"role": "user", "content": "What is Deep Learning?"}], "max_tokens": 17}'
```
 I am using port 8008 for the container running on CPU and port 8009 for the container running on GPU.

 6. Confirming MistralAI is Running

I ran the same curl command on both ports and and waited for the response.
But I was waiting for too long. 

Then I realised that both images were created for CPU because my system GPU was 0%.

I connected to the container that was supposed to be running on GPU.
```bash
docker exec -it vllm-openvino-server-gpu /bin/bash
root@8cb29cf583f8:/workspace# docker exec -it vllm-openvino-server-gpu /bin/bash
root@8cb29cf583f8:/workspace# python3 -c "from openvino import Core; print(Core().available_devices)"
['CPU']
```

So I tweaked the scripts to have two separate images that would create two separate containers.
CPU on port 8008 and GPU on port 8009.

Verified I had two images from two repositories. 
```bash
REPOSITORY      TAG       IMAGE ID   
vllm-openvino   latest    ddf49df221fc  
opea/vllm-arc   latest    7e9808c6a238  
```

I also noticed a typo in the launch scripts. The default image to be used should be `vllm-openvino`, not `vllm:openvino`.


Eventually, after ensuring that the containers were created using the correct images for each device, I couldn't make the GPU container work. I connected and checked, it wasn't using GPU but CPU. 

```bash
docker exec -it vllm-openvino-server-gpu /bin/bash
root@8cb29cf583f8:/workspace# python3 -c "from openvino import Core; print(Core().available_devices)"
['CPU']
```

and after a few seconds, the container was stopping. 

## Further troubleshooting

```bash
Inside the container:
root@6fcb2399cc98:/workspace# ls /dev/dri
by-path  card0  renderD128
```
This confirms that the GPU is accessible inside the container.
```bash
 docker exec -it vllm-openvino-server-gpu /bin/bash
root@6649f5497a8b:/workspace# python3 -c "from openvino.runtime import Core; print(Core().available_devices)"
/usr/lib/python3.10/importlib/util.py:247: DeprecationWarning: The `openvino.runtime` module is deprecated and will be removed in the 2026.0 release. Please replace `openvino.runtime` with `openvino`.
  self.__spec__.loader.exec_module(self)
['CPU']
```
But the container doesn't see it.
#### Bold step
Modifying the Dockerfile.intel_gpu to include OpenVINO installation didn't work.


> 2025-02-19


### Goal
Today I attempted to deploy Mistral-7B-Instruct-v0.3 using vLLM with Intel Arc GPU acceleration in WSL2.

### System Configuration
- OS: Windows 11 Pro with WSL2 Ubuntu
- GPU: Intel Arc Graphics (Driver version: 32.0.101.6129)
- Environment: Docker container with OpenVINO support

### Steps Taken

1. **Initial Approach - OpenVINO GPU**
   - Tried using vLLM with OpenVINO GPU acceleration
   - Used Dockerfile.intel_gpu for container build
   - Attempted to run with `launch_vllm_service_openvino_gpu.sh`

2. **Issues Encountered**
   ```
   [GPU] Can't get PERFORMANCE_HINT property as no supported devices found
   [GPU] Please check OpenVINO documentation for GPU drivers setup guide
   ```

3. **Attempted Solutions**
   - Tried installing Intel GPU drivers and dependencies
   - Attempted to use OpenVINO 2023 instead of 2025
   - Explored WSL2 GPU passthrough configuration

## Challenges
1. WSL2 GPU passthrough for Intel Arc is not straightforward
2. OpenVINO GPU support in container environment needs specific setup
3. Version compatibility issues between vLLM and OpenVINO

### Available Options

1. **CPU-Only Deployment**
   - Pros: Works immediately, simpler setup
   - Cons: Slower inference time

2. **Native Windows Deployment**
   - Pros: Direct access to Intel Arc GPU
   - Cons: Requires different setup approach

3. **Alternative Frameworks**
   - Intel's optimized transformers
   - Intel Neural Chat
   - Direct PyTorch with Intel extensions

### Next Steps

1. Short term: Deploy using CPU-only version for testing even though it's slower and doesn't return an output.
2. Long term: Either
   - Set up native Windows environment for GPU acceleration
   - Or implement using Intel's optimized frameworks

### Final decisions

I will first try to find another OPEA component that I can use to deploy Mistral-7B-Instruct-v0.3.
If I cannot find one, I will consider skipping OPEA components and Set up native Windows environment for GPU acceleration

### References
- [vLLM Documentation](https://github.com/vllm-project/vllm)
- [OpenVINO GPU Support](https://docs.openvino.ai/latest/openvino_docs_install_guides_installing_openvino_docker.html)
- [Intel Arc GPU Setup](https://www.intel.com/content/www/us/en/developer/articles/guide/getting-started-with-intel-oneapi-base-toolkit-in-wsl-2.html) 