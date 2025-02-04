# Free GenAI Bootcamp 2025 - Week0

## Table of Contents

- [GenAI Architecting](#genai-architecting)
- [](#)
  - [](#)
  - [](#)

> 2025-02-02

## GenAI Architecting

Today, I focused on creating a **conceptual diagram** in Lucidchart following along the GenAI Architecting lecture.  

As a Solution Architect, my goal is to design architectural diagrams that help stakeholders understand the key components of GenAI workloads. 
These diagrams should guide discussions on **infrastructure choices, integration patterns, and system dependencies** without directly prescribing solutions.  

#### Conceptual Diagram  
**Purpose**  
- Provides a **high-level overview** of the business solution.  
- Helps non-technical stakeholders understand the general architecture.  
- Focuses on **what** is being built, not **how** it will be implemented.  

[View Diagram](https://lucid.app/lucidchart/86d79ef3-7e66-4ae4-aa6b-e58b7ef201c2/edit?viewport_loc=-77%2C175%2C2992%2C1391%2C0_0&invitationId=inv_86590b77-e099-4abc-bdb6-96519ed9ea43)


> 2025-02-03

Completed a conceptual diagram in Lucidchart based on the instructor's work to illustrate key components of a GenAI workload. I also documented architectural considerations in a README file to align with the business and technical goals.

> 2025-02-04

I realised that Granite 3.0 might need further finetuning to be used for Ancient Greek. So I am trying to identify which LLM model would be more appropriate for my use case. 

| **Model**                 | **Pricing** | **License**                           | **Who Can Use It**                                              | **AI PC Specs Needed**                                         | **Use Case & Business Model Alignment**                                                                 |
|---------------------------|-------------|--------------------------------------|-----------------------------------------------------------------|----------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| **[Mistral 7B Instruct](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.3)**  | Free        | Apache 2.0                           | Individuals, researchers, commercial use (free)                 | 32GB RAM, GPU with 16GB VRAM (quantization possible for CPU)    | Text generation, language learning activities, aligns well with interactive lessons                      |
| **[LLaMA 2-7B-Chat](https://huggingface.co/meta-llama/Llama-2-7b-chat-hf)**       | Free        | Meta LLaMA 2 Community License       | Free for research and commercial use with conditions            | 32GB RAM, GPU with 14GB VRAM (quantization for CPU use)         | Conversational AI, chatbot features, great for language practice scenarios                              |
| **[MarianMT (English-Ancient Greek)](https://huggingface.co/docs/transformers/en/model_doc/marian#marianmt)** | Free        | CC-BY-4.0                             | Free for academic, research, and commercial use                 | 32GB RAM, runs efficiently on CPU without GPU                   | Translation tasks, vocabulary exercises, aligns with translation-focused modules                        |


**IBM Granite 3.0 8B Instruct** is designed as a **general-purpose instruction-following model**, similar to models optimized for **task completion**, **text generation**, **question answering**, and **conversational AI**.

---

### 🚀 **Comparison of IBM Granite 3.0 with Ancient Greek translating Models**

| **Model**                   | **Functionality Compared to IBM Granite** | **Key Differences** |
|-----------------------------|-------------------------------------------|---------------------|
| **Mistral 7B Instruct**     |  **Most Similar**                        | Both are **instruction-tuned LLMs** focused on text generation, task completion, and conversational tasks. Mistral is lightweight but powerful, while Granite may offer more enterprise-grade reliability. |
| **LLaMA 2-7B-Chat**         |  **Similar for Conversational AI**       | Optimized specifically for **chat-based interactions**. While Granite handles broader tasks, LLaMA 2-7B-Chat excels in **dialogue flow** and natural conversation. |
| **MarianMT (English-Ancient Greek)**|  **Not Similar**                        | **Task-specific translation model** focused solely on language translation (e.g., English ↔ Ancient Greek). It lacks general instruction-following capabilities. |

---

###  **Final Verdict**

-  **Closest Match:** **Mistral 7B Instruct**  
   - **Why:** Both Mistral and IBM Granite are versatile, capable of handling **instruction-following tasks**, **text generation**, and **multi-purpose NLP applications** without being restricted to conversational or translation-specific roles.  
   - Mistral is also highly efficient for deployment, especially with my AI PC setup.  

-  **Runner-up:** **LLaMA 2-7B-Chat** (if conversational AI is the primary focus).


### Considerations:
- Deploy using **Hugging Face Inference API** or **AWS SageMaker** for scalable infrastructure in case my AI PC cannot handle it.

### Next steps

Since my local machine is the Intel AI PC Development Kit, equipped with an Intel® Core™ Ultra 7 Processor 155H, 32 GB of DDR5 RAM, and 512 GB storage, 
I will use conda to set up an environment to install required libraries and load Mistral 7B with 4-bit Quantization.

**Quantization** is the process of reducing the precision of the model's weights from 32-bit (full precision) to lower bit representations like 8-bit, 4-bit, or even 2-bit.
It's importand because:
- **Reduces Memory Usage**: Makes large models like Mistral fit into your system's RAM.
- **Speeds Up Inference**: Less data to process means faster results, especially with NPUs.
Optionally, I could also focus on **Model Optimization**

**Optimization** involves fine-tuning how the model runs to make it faster and more efficient without changing its accuracy.

- **Techniques Include:**  
  - **Operator Fusion:** Combines operations to reduce computation time.  
  - **ONNX Runtime:** Converts models into an optimized format for faster inference.  
  - **Using NPUs:** Offloads specific tasks to your Intel NPU for acceleration.


### Create and Activate the Conda Environment

```bash
# Create a Conda environment with Python 3.10
conda create -n mistral_env python=3.10.0 ipykernel -y

# Activate the environment
conda activate mistral_env
```


### Install JupyterLab and Required Libraries

```bash
# Install JupyterLab
conda install -c conda-forge jupyterlab -y

# Install PyTorch (CPU version; adjust for GPU if needed)
conda install pytorch torchvision torchaudio cpuonly -c pytorch -y

# Install Transformers, Accelerate, and bitsandbytes for quantization
pip install transformers accelerate bitsandbytes
```


### Launch JupyterLab

```bash
jupyter lab --no-browser --allow-root --ip 0.0.0.0
```

### Import Mistral 7B in Jupyter Notebook


```python
from transformers import AutoModelForCausalLM, AutoTokenizer

# Load Mistral 7B Instruct with 4-bit quantization
model_name = "mistralai/Mistral-7B-Instruct-v0.3"

# Load model with quantization for better performance
model = AutoModelForCausalLM.from_pretrained(
    model_name,
    device_map="auto",
    load_in_4bit=True  # Applies quantization
)

tokenizer = AutoTokenizer.from_pretrained(model_name)

# Test prompt for Ancient Greek translation
prompt = "Translate this sentence into Ancient Greek: 'Knowledge is power.'"
inputs = tokenizer(prompt, return_tensors="pt").to("cuda" if model.device.type == "cuda" else "cpu")
outputs = model.generate(**inputs, max_length=50)

# Display the output
response = tokenizer.decode(outputs[0], skip_special_tokens=True)
print(response)
```
