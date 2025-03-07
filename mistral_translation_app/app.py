import os
import re
from fastapi import FastAPI, Depends, HTTPException, status, Header
from fastapi.middleware.cors import CORSMiddleware
from fastapi.security import APIKeyHeader  # Added this import
import time
import uvicorn
from transformers import AutoTokenizer, AutoModelForCausalLM
import torch

app = FastAPI()

# Add CORS Middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Update origins as needed for production
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Simple rate limiter
class RateLimiter:
    def __init__(self, calls: int, period: int):
        self.calls = calls  # max calls per period
        self.period = period  # period in seconds
        self.timestamps = []

    async def __call__(self):
        now = time.time()
        # Remove timestamps older than period
        self.timestamps = [t for t in self.timestamps if now - t < self.period]
        
        if len(self.timestamps) >= self.calls:
            raise HTTPException(
                status_code=status.HTTP_429_TOO_MANY_REQUESTS,
                detail=f"Rate limit exceeded: {self.calls} calls per {self.period} seconds"
            )
        
        self.timestamps.append(now)
        return True

# Create a rate limiter - 10 calls per minute
rate_limiter = RateLimiter(calls=10, period=60)

# Optional API key authentication
API_KEY = os.getenv("API_KEY")
api_key_header = APIKeyHeader(name="X-API-Key", auto_error=False)

async def verify_api_key(api_key: str = Header(None)):
    if API_KEY and api_key != API_KEY:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Invalid API key"
        )
    return True

# Load Hugging Face token from environment variables
HF_TOKEN = os.getenv("HF_TOKEN")
if HF_TOKEN is None:
    raise ValueError("Missing Hugging Face Token! Set HF_TOKEN environment variable.")

# Load model at startup
print("Loading model and tokenizer...")
model_name = "mistralai/Mistral-7B-Instruct-v0.3"
tokenizer = AutoTokenizer.from_pretrained(model_name, token=HF_TOKEN)
model = AutoModelForCausalLM.from_pretrained(
    model_name,
    torch_dtype=torch.float16,
    device_map="auto",
    token=HF_TOKEN
)
print("Model loaded successfully!")

@app.get("/")
def read_root():
    return {"status": "Mistral Translation API is running"}

@app.post("/translate")
async def translate(sentence: str, api_key: str = Header(None)):
    secret_api_key = os.getenv("SECRET_API_KEY")
    if api_key != secret_api_key:
        raise HTTPException(status_code=401, detail="Invalid API key")

    prompt = f"""
Translate the following English sentence accurately into Classical Ancient Greek (Attic dialect). 
Use correct grammar, vocabulary, and polytonic diacritics. 
Provide ONLY the Ancient Greek translation enclosed by <START> and <END> tags, nothing else.

Examples:
English: Wisdom is virtue. → <START>Σοφία ἐστὶν ἀρετή.<END>
English: Life is short. → <START>Ὁ βίος βραχύς ἐστιν.<END>
English: Know thyself. → <START>Γνῶθι σεαυτόν.<END>
English: Hello world. → <START>Χαῖρε, ὦ κόσμε!<END>
English: I love philosophy. → <START>Φιλοσοφίαν φιλῶ.<END>

Now translate accurately:
English: {sentence} → <START>"""


    inputs = tokenizer(prompt, return_tensors="pt").to("cuda")
    outputs = model.generate(
        **inputs,
        max_new_tokens=60,
        temperature=0.1,
        do_sample=True,
        pad_token_id=tokenizer.eos_token_id
    )

    translated_text = tokenizer.decode(outputs[0], skip_special_tokens=True)

    match = re.search(r"<START>\s*(.+?)\s*<END>", translated_text, re.DOTALL)
    greek_text = match.group(1).strip() if match else "Translation unclear."

    return {
        "original": sentence,
        "translation": greek_text,
        "full_response": translated_text
    }




if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)