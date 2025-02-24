from flask import Flask, request, jsonify
import sqlite3
import requests
import random

app = Flask(__name__)

# TTS API URL (Docker container)
TTS_API_URL = "http://localhost:5002/api/tts"

# Add a root route for testing
@app.route('/')
def home():
    return jsonify({"status": "Server is running", "endpoints": ["/api/generate-audio", "/api/random-sequence"]})

@app.route('/api/generate-audio', methods=['POST'])
def generate_audio():
    text = request.json.get('text')
    if not text:
        return jsonify({"error": "No text provided"}), 400
    
    try:
        # Send to TTS server
        response = requests.post(
            TTS_API_URL,
            data={"text": text}
        )
        
        if response.status_code == 200:
            return response.content, 200, {'Content-Type': 'audio/wav'}
        else:
            return jsonify({"error": f"TTS generation failed: {response.text}"}), 500
    except Exception as e:
        return jsonify({"error": f"Server error: {str(e)}"}), 500

@app.route('/api/random-sequence', methods=['GET'])
def get_random_sequence():
    try:
        conn = sqlite3.connect("listening-learning.db")
        cursor = conn.cursor()
        cursor.execute("SELECT passage FROM greek_texts ORDER BY RANDOM() LIMIT 1")
        passage = cursor.fetchone()
        conn.close()
        
        if passage:
            # Split passage into sentences by commas
            sentences = [s.strip() for s in passage[0].split(',')]
            # Filter for sentences with at least 5 words
            valid_sentences = [s for s in sentences if len(s.split()) >= 5]
            
            if valid_sentences:
                # Choose a random valid sentence
                selected = random.choice(valid_sentences)
                return jsonify({
                    "text": selected,
                    "full_passage": passage[0]
                })
            else:
                return jsonify({"error": "No sentences with 5+ words found"}), 400
                
        return jsonify({"error": "No passages found"}), 404
        
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    print("🚀 Server starting on http://localhost:5000")
    app.run(host='0.0.0.0', port=5000, debug=True) 