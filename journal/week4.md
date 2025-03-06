## Table of Contents

- [Table of Contents](#table-of-contents)
- [AIPC - Multimodal Visual Language Experience Generator with Ria](#aipc---multimodal-visual-language-experience-generator-with-ria)



> 2025-03-05

## AIPC - Multimodal Visual Language Experience Generator with Ria

I will attempt to install on my Ubuntu WSL since I have already managed to install Openvino on Windows. 

1. Downloaded the repository https://github.com/openvinotoolkit/openvino_build_deploy.git and copied it into my free-genai-bootcamp-2025 repository
2. Changed directory to multimodal_ai_visual_generator
```bash
cd ai_ref_kits/multimodal_ai_visual_generator
```
3. Created a python environment and installed requirements
```bash
conda create --name run_env python=3.9 -y && conda activate run_env && pip install -r requirements.txt
```
I had to change sherpa-onnx version from 1.10.41 to 1.10.43 because 1.10.41 was not found.

4. Ran script to download models
```bash
python3 download_and_prepare_models.py
```
At this point I realised that I needed access to llama models and also to add my huggingface token
```bash
huggingface-cli login
python3 download_and_prepare_models.py
```
I identified some warnings after the script ended.

- **Whisper and Torch JIT Tracer Warnings:**  
   - Several warnings indicate that tensors are being converted to Python booleans and that using functions like `len()` on tensors might lead to tracing inaccuracies. These warnings suggest that the model’s tracing (used for converting parts of the model to a static computation graph) might not generalize correctly to other inputs.  
   - There’s a specific Torch JIT trace warning where the output of the traced function doesn’t match the Python function output—this mismatch (with 100% of elements differing) could indicate potential issues in performance or reliability of the traced model.

- **Configuration Warnings:**  
   - A warning notes that a parameter (`loss_type=None`) in the model configuration isn’t recognized. As a result, it defaults to using `ForCausalLMLoss`. While this might be acceptable, it could lead to unexpected behavior if a different loss function was intended.
   - Additionally, there’s a notice about moving certain attributes from the main config to the generation config, which is informational but might be worth checking to ensure that generation behavior is as expected.

5. Here's where things get tricky because main.py needs to be written from scratch.

6. And although this was done, I found a limitation with WSL
```bash
Successfully imported openvino_genai
Using custom implementation for SuperResolutionPipeline
Using custom implementation for DepthEstimationPipeline
Successfully imported openvino_genai
Running in console mode due to WSL display limitations
```
- WSL doesn't have a display server configured
- The Qt libraries needed for the GUI couldn't connect to an X server

7. Switching to Powershell terminal

