# Ancient Greek Learning App

An interactive application for learning Ancient Greek with AI-powered question generation and audio features.

## Setup Instructions

### Prerequisites
- [Anaconda](https://www.anaconda.com/download) or [Miniconda](https://docs.conda.io/en/latest/miniconda.html)

### Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd listening-comp
```

1. Create the conda environment:
```bash
conda env create -f environment.yml
```
2. Activate the conda environment in all the terminals:
```bash
conda activate greek-learning
```


### Features

- AI-powered question generation for:
  - Reading comprehension
  - Grammar
  - Vocabulary
- Interactive UI with immediate feedback
- Audio support for Greek text


>For additional context, read the [journal entry](../journal/week2.md).

### Running the Application


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

The application should now be running at:
- Backend: http://localhost:5000
- Frontend: http://localhost:8501
- coquiTTS: http://localhost:5002/