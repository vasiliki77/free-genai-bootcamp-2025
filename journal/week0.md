# Free GenAI Bootcamp 2025 - Week0

## Table of Contents

- [GenAI Architecting](#genai-architecting)
- [Choosing appropriate model](#choosing-appropriate-model)
  - [Final Verdict](#final-verdict)
  - [Next steps](#next-steps)
- [Testing Mistral AI on my Intel AI PC](#testing-mistral-ai-on-my-intel-ai-pc)
  - [Tested Modern Greek to Ancient Greek translation](#tested-modern-greek-to-ancient-greek-translation)

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

## Choosing appropriate model

I realised that Granite 3.0 might need further finetuning to be used for Ancient Greek. So I am trying to identify which LLM model would be more appropriate for my use case. 

| **Model**                 | **Pricing** | **License**                           | **Who Can Use It**                                              | **AI PC Specs Needed**                                         | **Use Case & Business Model Alignment**                                                                 |
|---------------------------|-------------|--------------------------------------|-----------------------------------------------------------------|----------------------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| **[Mistral 7B Instruct](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.3)**  | Free        | Apache 2.0                           | Individuals, researchers, commercial use (free)                 | 32GB RAM, GPU with 16GB VRAM (quantization possible for CPU)    | Text generation, language learning activities, aligns well with interactive lessons                      |
| **[LLaMA 2-7B-Chat](https://huggingface.co/meta-llama/Llama-2-7b-chat-hf)**       | Free        | Meta LLaMA 2 Community License       | Free for research and commercial use with conditions            | 32GB RAM, GPU with 14GB VRAM (quantization for CPU use)         | Conversational AI, chatbot features, great for language practice scenarios                              |
| **[MarianMT (English-Ancient Greek)](https://huggingface.co/docs/transformers/en/model_doc/marian#marianmt)** | Free        | CC-BY-4.0                             | Free for academic, research, and commercial use                 | 32GB RAM, runs efficiently on CPU without GPU                   | Translation tasks, vocabulary exercises, aligns with translation-focused modules                        |


**IBM Granite 3.0 8B Instruct** is designed as a **general-purpose instruction-following model**, similar to models optimized for **task completion**, **text generation**, **question answering**, and **conversational AI**.

---

### Comparison of IBM Granite 3.0 with Ancient Greek translating Models

| **Model**                   | **Functionality Compared to IBM Granite** | **Key Differences** |
|-----------------------------|-------------------------------------------|---------------------|
| **Mistral 7B Instruct**     |  **Most Similar**                        | Both are **instruction-tuned LLMs** focused on text generation, task completion, and conversational tasks. Mistral is lightweight but powerful, while Granite may offer more enterprise-grade reliability. |
| **LLaMA 2-7B-Chat**         |  **Similar for Conversational AI**       | Optimized specifically for **chat-based interactions**. While Granite handles broader tasks, LLaMA 2-7B-Chat excels in **dialogue flow** and natural conversation. |
| **MarianMT (English-Ancient Greek)**|  **Not Similar**                        | **Task-specific translation model** focused solely on language translation (e.g., English ↔ Ancient Greek). It lacks general instruction-following capabilities. |

---

### Final Verdict

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
from huggingface_hub import login
from transformers import AutoModelForCausalLM, AutoTokenizer

# Login to Hugging Face
login("YOUR_HF_TOKEN")

# Load Mistral 7B Instruct with 4-bit quantization
model_name = "mistralai/Mistral-7B-Instruct-v0.3"

# Load model with quantization for better performance
model = AutoModelForCausalLM.from_pretrained(
    model_name,
    device_map="auto",
    load_in_4bit=True  # Applies quantization
)

# Load the tokenizer
tokenizer = AutoTokenizer.from_pretrained(model_name)

# Test prompt for Ancient Greek translation
prompt = "Translate this sentence into Ancient Greek: 'Knowledge is power.'"
inputs = tokenizer(prompt, return_tensors="pt").to("cuda" if model.device.type == "cuda" else "cpu")
outputs = model.generate(**inputs, max_length=50)

# Display the output
response = tokenizer.decode(outputs[0], skip_special_tokens=True)
print(response)
```
---

> 2025-02-05

## Testing Mistral AI on my Intel AI PC

Today, trying to follow the above steps, I came across this error:
```
/home/myhome/miniconda3/envs/mistral_env/lib/python3.10/site-packages/tqdm/auto.py:21: TqdmWarning: IProgress not found. Please update jupyter and ipywidgets. See https://ipywidgets.readthedocs.io/en/stable/user_install.html
  from .autonotebook import tqdm as notebook_tqdm
```
The resolution was to stop jupyterlab, and run the following in mistral_env.
```bash
# Install ipywidgets
conda install -c conda-forge ipywidgets -y
# Update JupyterLab to ensure compatibility
conda update -c conda-forge jupyterlab -y
```

After JupyterLab is restarted:


```bash
pip install torch transformers accelerate bitsandbytes sentencepiece protobuf
```

Although my system looks like it has GPU, I came across some error about CUDA. So I had to proceed with **CPU-only** and didn't need `bitsandbytes` (which relies on CUDA).

```bash
pip uninstall bitsandbytes
```


Had to create READ token and log in to **Hugging Face**:  

```python
from huggingface_hub import login

login("YOUR_HF_TOKEN")
```


Since I am using CPU, **I avoid bitsandbytes quantization**:

```python
from transformers import AutoModelForCausalLM, AutoTokenizer

# Load the model without quantization
model_name = "mistralai/Mistral-7B-Instruct-v0.3"
model = AutoModelForCausalLM.from_pretrained(
    model_name,
    device_map={"": "cpu"},  # Force CPU usage
    torch_dtype="auto"       # Use the best dtype for CPU
)

# Load the tokenizer
tokenizer = AutoTokenizer.from_pretrained(model_name)
```

Below I had to apply few-shot prompting because the model was returning the translation in latin characters or without [Greek diacritics](https://en.wikipedia.org/wiki/Greek_diacritics), which are not used in modern Greek. 
```python
prompt = (
    "Translate the following sentences into Ancient Greek using the Greek alphabet:\n"
    "1. 'Wisdom is virtue.' → Σοφία ἐστὶν ἀρετή.\n"
    "2. 'Life is short.' → Ὁ βίος βραχύς ἐστιν.\n"
    "3. 'Knowledge is power.' →"
)

# Tokenize the input
inputs = tokenizer(prompt, return_tensors="pt")

# Generate the output with enough tokens for full completion
outputs = model.generate(
    **inputs,
    max_new_tokens=40,  # Ensure it has enough tokens to complete the response
    pad_token_id=tokenizer.eos_token_id
)

# Decode and print the result
response = tokenizer.decode(outputs[0], skip_special_tokens=True)
print(response)
```

I had to increase 'max_new_tokens' to 40, otherwise the output was incomplete. 

Now that I know the model can perform on my local machine, I can proceed with Sentence Constructor. 

### Tested Modern Greek to Ancient Greek translation

```python
prompt = (
    "Translate the following sentences into Ancient Greek using the Greek alphabet:\n"
    "1. 'Η σοφία είναι αρετή.' → Σοφία ἐστὶν ἀρετή.\n"
    "2. 'Η ζωή είναι μικρή.' → Ὁ βίος βραχύς ἐστιν.\n"
    "3. 'Η γνώση είναι δύναμη.' →"
)

# Tokenize the prompt
inputs = tokenizer(prompt, return_tensors="pt")

# Generate output using max_new_tokens instead of max_length
outputs = model.generate(
    **inputs,
    max_new_tokens=30, 
    pad_token_id=tokenizer.eos_token_id
)

# Decode and display the output
response = tokenizer.decode(outputs[0], skip_special_tokens=True)
print(response)
```

Output:
```
Translate the following sentences into Ancient Greek using the Greek alphabet:
1. 'Η σοφία είναι αρετή.' → Σοφία ἐστὶν ἀρετή.
2. 'Η ζωή είναι μικρή.' → Ὁ βίος βραχύς ἐστιν.
3. 'Η γνώση είναι δύναμη.' → Γνώσις ἐστὶν δύναμις.
```

This means that it can also handle Modern Greek to Ancient Greek translation.