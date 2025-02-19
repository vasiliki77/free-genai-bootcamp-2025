
1. **Create and Activate OpenVINO Environment**
````powershell
# Create environment
python -m venv openvino_env

# Activate it
openvino_env\Scripts\activate

# Upgrade pip
python -m pip install --upgrade pip
````


2. **Install OpenVINO**
````powershell
python -m pip install openvino

# Verify installation
python -c "from openvino import Core; print(Core().available_devices)"
# Output shows: ['CPU', 'GPU', 'NPU']
````


3. **Install Dependencies**
````powershell
# Install PyTorch and Intel Extensions
pip install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cpu
# Install required packages
pip install transformers
pip install accelerate
pip install --upgrade huggingface_hub
pip install protobuf
pip install safetensors
# Install vLLM
pip install vllm==0.7.2
````


4. **Create Test Files**
````powershell
# Create test_inference.py and run it
python test_inference.py
# Successfully showed GPU detection and computation
````


5. **Set up Hugging Face Access**
````powershell
# Set environment variable for HF token
$env:HF_TOKEN="my_token"
````


Current blocker: Need to install sentencepiece package to run the Mistral model.