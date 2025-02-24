
For additional context, read the [journal entry](../journal/week2.md).

1. First, run the container with bash entrypoint:
```bash
docker run --rm -it -p 5002:5002 --entrypoint /bin/bash ghcr.io/coqui-ai/tts-cpu
```

2. Inside the container, start the TTS server:
```bash
python3 TTS/server/server.py --model_name tts_models/el/cv/vits
```

3. In a new terminal, install the requirements:
```bash
cd listening-comp
pip install -r requirements.txt
```

4. Initialize the database:
```bash
python backend/init_db.py
```

5. Start the Flask backend (in a new terminal):
```bash
python backend/main.py
```

6. Start the Streamlit frontend (in another new terminal):
```bash
streamlit run frontend/main.py
```


