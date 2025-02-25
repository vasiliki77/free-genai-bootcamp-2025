import streamlit as st
import requests
import json
import time
from audio_recorder_streamlit import audio_recorder
import logging

# Configure logging
logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

# Constants
BACKEND_URL = "http://localhost:5000"
TTS_URL = "http://localhost:5002/api/tts"

# Configure the page
st.set_page_config(
    page_title="Ancient Greek Learning",
    page_icon="🏺",
    layout="wide"
)

# Initialize session state
if 'current_sequence' not in st.session_state:
    st.session_state.current_sequence = None
if 'current_audio' not in st.session_state:
    st.session_state.current_audio = None
if 'question' not in st.session_state:
    st.session_state.question = None
if 'feedback' not in st.session_state:
    st.session_state.feedback = None
if 'selected_answer' not in st.session_state:
    st.session_state.selected_answer = None
if 'show_answer' not in st.session_state:
    st.session_state.show_answer = False
if 'practice_mode' not in st.session_state:
    st.session_state.practice_mode = "Reading Practice"
if 'audio_played' not in st.session_state:
    st.session_state.audio_played = False
if 'question_type' not in st.session_state:
    st.session_state.question_type = "comprehension"
if 'difficulty' not in st.session_state:
    st.session_state.difficulty = "intermediate"

def reset_state():
    """Reset all session state variables"""
    st.session_state.current_sequence = None
    st.session_state.current_audio = None
    st.session_state.question = None
    st.session_state.feedback = None
    st.session_state.selected_answer = None
    st.session_state.show_answer = False
    st.session_state.audio_played = False

def generate_question(difficulty="intermediate", question_type="comprehension"):
    """Generate a question based on the current settings"""
    try:
        response = requests.post(
            f"{BACKEND_URL}/api/generate-question",
            json={
                "difficulty": difficulty,
                "type": question_type
            },
            timeout=10
        )
        logger.debug(f"Question generation response status: {response.status_code}")
        logger.debug(f"Response content: {response.text}")
        
        if response.status_code == 200:
            return response.json()
        else:
            st.error(f"Error generating question: {response.json().get('error', 'Unknown error')}")
            return None
    except Exception as e:
        st.error(f"Error connecting to backend: {str(e)}")
        return None

def get_random_sequence():
    """Fetch a random Greek sequence from the backend"""
    try:
        response = requests.get(f"{BACKEND_URL}/api/random-sequence")
        if response.status_code == 200:
            return response.json()["sequence"]
        else:
            st.error(f"Error fetching sequence: {response.json().get('error', 'Unknown error')}")
            return None
    except Exception as e:
        st.error(f"Error connecting to backend: {str(e)}")
        return None

def generate_audio(text):
    """Generate audio for the given text"""
    try:
        response = requests.post(
            f"{BACKEND_URL}/api/generate-audio",
            json={"text": text}
        )
        if response.status_code == 200:
            return response.content
        else:
            st.error("Error generating audio")
            return None
    except Exception as e:
        st.error(f"Error generating audio: {str(e)}")
        return None

# Main UI
st.title("🏺 Ancient Greek Learning")
st.write("Practice your Ancient Greek reading and listening skills!")

# Sidebar controls
with st.sidebar:
    st.header("Settings")
    
    # Practice mode selector
    st.subheader("Practice Mode")
    practice_mode = st.radio(
        "Practice Mode",
        ["Reading Practice", "Listening Practice"],
        key="practice_mode"
    )
    
    if practice_mode == "Reading Practice":
        # Question type selector
        st.subheader("Question Type")
        question_type = st.selectbox(
            "Question Type",
            ["comprehension"],
            key="question_type"
        )
        
        # Difficulty selector
        st.subheader("Difficulty")
        difficulty = st.selectbox(
            "Difficulty",
            ["intermediate"],
            key="difficulty"
        )

        st.markdown("---")
        
        # Generate question button
        if st.button("Generate New Question", type="primary", use_container_width=True):
            with st.spinner("Generating question..."):
                question = generate_question(difficulty, question_type)
                if question:
                    try:
                        # Parse the question data
                        question_data = json.loads(question.get('question', '{}'))
                        st.session_state.question = question_data
                        st.session_state.current_sequence = question.get('passage', '')
                    except json.JSONDecodeError:
                        st.error("Error parsing question data")
                        st.session_state.question = None
                        st.session_state.current_sequence = None

# Main content area
if practice_mode == "Reading Practice":
    if st.session_state.question:
        # Create columns for better layout
        col1, col2 = st.columns([2, 1])
        
        with col1:
            # Display the word or phrase being asked about
            st.subheader("Study This")
            st.markdown(f"### {st.session_state.question.get('word', '')}")
            
            # Show translation button
            if st.button("Show Translation"):
                st.info(st.session_state.question.get('translation', ''))
        
        with col2:
            # Display the question
            st.subheader("Question")
            st.write(st.session_state.question.get('question', 'Question not available'))
            
            # Display options
            options = st.session_state.question.get('options', [])
            if options:
                selected = st.radio("Choose your answer:", options, key="answer_selection")
                selected_index = options.index(selected)
                
                if st.button("Check Answer"):
                    correct_answer = st.session_state.question.get('correct_answer', 0)
                    if selected_index == correct_answer - 1:  # Adjust for 0-based indexing
                        st.success("Correct! 🎉")
                    else:
                        st.error("Incorrect. Try again!")
                    st.info(f"Explanation: {st.session_state.question.get('explanation', 'No explanation available')}")
    else:
        st.info("👈 Click 'Generate New Question' to start learning!")

elif practice_mode == "Listening Practice":
    if st.button("Get New Sequence"):
        st.session_state.current_sequence = get_random_sequence()
        if st.session_state.current_sequence:
            st.session_state.current_audio = generate_audio(st.session_state.current_sequence)

    if st.session_state.current_sequence:
        st.subheader("Listen and Practice")
        
        # Audio playback
        if st.session_state.current_audio:
            st.audio(st.session_state.current_audio, format='audio/wav')
        
        # Show text button
        if st.button("Show Text"):
            st.markdown(f"```{st.session_state.current_sequence}```")
        
        # Recording section
        st.subheader("Record Your Pronunciation")
        audio_bytes = audio_recorder()
        if audio_bytes:
            st.audio(audio_bytes, format="audio/wav")
    else:
        st.info("👆 Click 'Get New Sequence' to start practicing pronunciation!") 