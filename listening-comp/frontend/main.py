import streamlit as st
import requests
import io
from audio_recorder_streamlit import audio_recorder

# Backend API URL
API_URL = "http://localhost:5000/api"

st.title("Ancient Greek Learning App")

# Initialize session state for storing the text
if 'current_text' not in st.session_state:
    st.session_state.current_text = None

# Get random sequence button
if st.button("Get Random Greek Sequence"):
    response = requests.get(f"{API_URL}/random-sequence")
    if response.status_code == 200:
        data = response.json()
        st.session_state.current_text = data["text"]  # Store the text
        
        # Display the 7-word sequence
        st.write("📜 Selected Sequence:")
        st.write(data["text"])
        
        # Show position info
        st.write("ℹ️ " + data["sequence_position"])
        
        # Option to show full passage
        if st.checkbox("Show full passage"):
            st.write("Full passage:")
            st.write(data["full_passage"])
    else:
        st.error("Failed to get random sequence")

# Separate Listen button
if st.session_state.current_text and st.button("🔊 Listen"):
    audio_response = requests.post(
        f"{API_URL}/generate-audio",
        json={"text": st.session_state.current_text}
    )
    if audio_response.status_code == 200:
        st.audio(audio_response.content, format="audio/wav")
    else:
        st.error(f"Failed to generate audio: {audio_response.text}")

# Record user's pronunciation
st.write("🎤 Record your pronunciation:")
audio_bytes = audio_recorder()
if audio_bytes:
    st.audio(audio_bytes, format="audio/wav") 