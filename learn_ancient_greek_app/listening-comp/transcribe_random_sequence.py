import sqlite3
import requests
import random

def get_random_sequence(db_path="listening-learning.db", num_words=7):
    """Get a random sequence of words from the database."""
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    
    # Get all passages
    cursor.execute("SELECT passage FROM greek_texts")
    passages = cursor.fetchall()
    
    if not passages:
        print("❌ No passages found in database")
        return None
        
    # Take a random passage and split into words
    passage = random.choice(passages)[0]
    words = [word for word in passage.split() if len(word) > 1]  # Filter out single chars
    
    # Select random sequence of words
    if len(words) >= num_words:
        start_idx = random.randint(0, len(words) - num_words)
        selected_words = words[start_idx:start_idx + num_words]
    else:
        selected_words = words
        
    conn.close()
    return " ".join(selected_words)

def text_to_speech(text, output_file="random_sequence.wav"):
    """Convert text to speech using the TTS server."""
    response = requests.post(
        "http://localhost:5002/api/tts",
        data={"text": text}
    )
    
    if response.status_code == 200:
        with open(output_file, "wb") as f:
            f.write(response.content)
        print(f"✅ Audio saved as {output_file}")
    else:
        print("❌ Error generating audio")
        print("Status code:", response.status_code)
        print("Response:", response.text)

def main():
    # Get random sequence
    print("🔎 Getting random sequence from database...")
    sequence = get_random_sequence()
    
    if sequence:
        print("\n📜 Selected text:")
        print(sequence)
        print("\n🔊 Converting to speech...")
        text_to_speech(sequence)

if __name__ == "__main__":
    main() 