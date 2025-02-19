from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import Optional
import os
from transformers import AutoModelForCausalLM, AutoTokenizer
import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI()

class TranslationRequest(BaseModel):
    text: str
    max_tokens: Optional[int] = 50
    temperature: Optional[float] = None

# Initialize at startup
logger.info("Loading Mistral model...")
model_name = "mistralai/Mistral-7B-Instruct-v0.3"  # Latest instruction-tuned version

try:
    logger.info(f"Using HF token: {'token exists' if os.getenv('HF_TOKEN') else 'no token found'}")
    tokenizer = AutoTokenizer.from_pretrained(model_name, token=os.getenv('HF_TOKEN'))
    model = AutoModelForCausalLM.from_pretrained(
        model_name,
        token=os.getenv('HF_TOKEN'),
        device_map="cpu",
        torch_dtype="auto",
        low_cpu_mem_usage=True,
        quantization_config=None
    )
    logger.info("Model loaded successfully")
except Exception as e:
    logger.error(f"Failed to load model: {str(e)}")
    raise

@app.post("/v1/translate")
async def translate(request: TranslationRequest):
    """
    Translate English text to Ancient Greek
    
    Parameters:
    - text: The English text to translate
    - max_tokens: Maximum length of translation
    - temperature: Controls randomness (0.0-1.0)
    
    Returns:
    - translation: The Ancient Greek translation
    - input_text: The original English text
    """
    if not request.text.strip():
        raise HTTPException(status_code=400, detail="Text cannot be empty")
    
    if len(request.text.split()) > 100:
        raise HTTPException(status_code=400, detail="Text too long")
    
    try:
        # Format prompt for Mistral
        prompt = f"""<s>[INST] Translate to Ancient Greek:
1. 'Wisdom is virtue.' → Σοφία ἐστὶν ἀρετή.
Now translate: {request.text} [/INST]"""

        inputs = tokenizer(prompt, return_tensors="pt")
        outputs = model.generate(
            **inputs,
            max_new_tokens=30,  # Keep only this, remove max_length
            do_sample=False,
            num_beams=4,        # Add beam search
            early_stopping=True, # Now valid with num_beams>1
            pad_token_id=tokenizer.eos_token_id
        )
        translation = tokenizer.decode(outputs[0], skip_special_tokens=True)
        
        return {
            "translation": translation.strip(),
            "input_text": request.text
        }
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    import uvicorn
    port = int(os.getenv("PORT", "8000"))
    uvicorn.run(app, host="0.0.0.0", port=port) 