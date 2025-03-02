from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import uvicorn
import os
import time
import logging

# Configure logging
logging.basicConfig(level=logging.INFO, 
                    format='%(asctime)s - %(levelname)s - %(message)s')
logger = logging.getLogger(__name__)

# Initialize FastAPI app
app = FastAPI()

# Configure LLM endpoint
llm_endpoint_url = os.environ.get("llm_endpoint_url", "http://tgi-server:80/generate")
logger.info(f"Using LLM endpoint: {llm_endpoint_url}")

# Initialize connection to the LLM
llm = None
try:
    from langchain.llms import HuggingFaceTextGenInference
    from langchain.prompts import PromptTemplate
    
    # Initialize LLM
    llm = HuggingFaceTextGenInference(
        inference_server_url=llm_endpoint_url,
        max_new_tokens=256,
        temperature=0.1,
        timeout=60,
    )
    logger.info("LLM client initialized successfully")
except Exception as e:
    logger.error(f"Error initializing LLM client: {str(e)}")

# Translation prompt template - simplified for smaller model
TRANSLATION_TEMPLATE = """You are translating from English to Ancient Greek. 
Translate this English phrase to Ancient Greek: {input_text}

Ancient Greek Translation:"""

translation_prompt = PromptTemplate(
    template=TRANSLATION_TEMPLATE,
    input_variables=["input_text"],
)

class TranslationRequest(BaseModel):
    input_text: str

class TranslationResponse(BaseModel):
    input_text: str
    output: str

@app.get("/health")
def health_check():
    return {"status": "healthy"}

@app.post("/v1/translate", response_model=TranslationResponse)
def translate_text(request: TranslationRequest):
    try:
        global llm
        if not llm:
            # Recreate LLM client if it was not initialized
            from langchain.llms import HuggingFaceTextGenInference
            llm = HuggingFaceTextGenInference(
                inference_server_url=llm_endpoint_url,
                max_new_tokens=256,
                temperature=0.1,
                timeout=60,
            )
            
        # Format the prompt
        prompt = translation_prompt.format(input_text=request.input_text)
        logger.info(f"Processing translation request for: {request.input_text}")
        
        # Get translation from LLM
        start_time = time.time()
        result = llm(prompt)
        end_time = time.time()
        logger.info(f"Translation completed in {end_time - start_time:.2f} seconds")
        
        # Extract just the translation text (sometimes models output additional text)
        cleaned_result = result.strip()
        
        # Return the result
        return TranslationResponse(
            input_text=request.input_text,
            output=cleaned_result
        )
    except Exception as e:
        logger.error(f"Translation error: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080) 