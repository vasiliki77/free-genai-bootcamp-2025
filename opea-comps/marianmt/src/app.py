from fastapi import FastAPI
from transformers import MarianMTModel, MarianTokenizer
import uvicorn
from pydantic import BaseModel
import logging
import sys
import os

# Set up logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI()

try:
    # Log environment variables (excluding sensitive data)
    logger.info("Starting application...")
    logger.info(f"HF_TOKEN present: {'HF_TOKEN' in os.environ}")
    
    # Load model and tokenizer
    logger.info("Loading model and tokenizer...")
    model_name = "Helsinki-NLP/opus-mt-en-grc"
    tokenizer = MarianTokenizer.from_pretrained(model_name, use_auth_token=os.environ.get('HF_TOKEN'))
    model = MarianMTModel.from_pretrained(model_name, use_auth_token=os.environ.get('HF_TOKEN'))
    logger.info("Model and tokenizer loaded successfully")

except Exception as e:
    logger.error(f"Error during startup: {str(e)}")
    sys.exit(1)

class TranslationRequest(BaseModel):
    text: str

@app.post("/translate")
async def translate(request: TranslationRequest):
    try:
        # Tokenize and translate
        inputs = tokenizer(request.text, return_tensors="pt", padding=True)
        outputs = model.generate(**inputs)
        translation = tokenizer.decode(outputs[0], skip_special_tokens=True)
        
        return {"translation": translation}
    except Exception as e:
        logger.error(f"Translation error: {str(e)}")
        return {"error": str(e)}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000) 