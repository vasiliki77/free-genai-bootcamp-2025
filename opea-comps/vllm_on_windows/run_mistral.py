from transformers import AutoModelForCausalLM, AutoTokenizer
import openvino as ov
import os

# Initialize OpenVINO Core and check devices
core = ov.Core()
print(f"Available devices: {core.available_devices}")

# Model ID
model_id = "mistralai/Mistral-7B-Instruct-v0.3"

try:
    # Load tokenizer
    print("Loading tokenizer...")
    tokenizer = AutoTokenizer.from_pretrained(
        model_id,
        token=os.environ.get('HF_TOKEN')  # Use token from environment
    )

    # Load model
    print("Loading model...")
    model = AutoModelForCausalLM.from_pretrained(
        model_id,
        device_map="auto",
        trust_remote_code=True,
        token=os.environ.get('HF_TOKEN')  # Use token here too
    )

    # Prepare input
    prompt = "What is artificial intelligence?"
    print(f"\nPrompt: {prompt}")

    # Generate text
    print("Generating response...")
    inputs = tokenizer(prompt, return_tensors="pt")
    output = model.generate(
        **inputs,
        max_length=100,
        temperature=0.7,
        do_sample=True
    )

    # Decode and print response
    response = tokenizer.decode(output[0], skip_special_tokens=True)
    print(f"\nResponse: {response}")

except Exception as e:
    print(f"An error occurred: {str(e)}") 