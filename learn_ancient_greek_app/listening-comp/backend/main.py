from flask import Flask, request, jsonify
import sqlite3
import requests
import random
import os
from question_generator import GreekQuestionGenerator
from flask_cors import CORS  # Add CORS support
import logging

app = Flask(__name__)
CORS(app)  # Enable CORS for all routes

# Set up logging
logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

# Set up database path
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
DB_PATH = os.path.join(BASE_DIR, 'backend', 'listening-learning.db')

def get_db_connection():
    try:
        conn = sqlite3.connect(DB_PATH)
        conn.row_factory = sqlite3.Row
        return conn
    except Exception as e:
        logger.error(f"Database connection error: {str(e)}")
        return None

def check_database():
    """Check if database exists and has data"""
    if not os.path.exists(DB_PATH):
        return False, f"Database file not found at {DB_PATH}. Please run init_db.py first."
    
    try:
        conn = get_db_connection()
        if conn is None:
            return False, "Could not connect to database"
        
        cur = conn.cursor()
        cur.execute("SELECT COUNT(*) FROM greek_texts")
        count = cur.fetchone()[0]
        conn.close()
        
        if count == 0:
            return False, "Database is empty. Please run init_db.py to populate it."
        return True, f"Database contains {count} passages"
    except sqlite3.OperationalError:
        return False, "Database table not found. Please run init_db.py to initialize the database."
    except Exception as e:
        return False, f"Database error: {str(e)}"

# Initialize components
db_ok, db_message = check_database()
if not db_ok:
    print(f"⚠️ {db_message}")
else:
    print(f"✅ {db_message}")

# Initialize question generator with DB_PATH
question_generator = GreekQuestionGenerator(DB_PATH)

# TTS API URL (Docker container)
TTS_API_URL = "http://localhost:5002/api/tts"

# Add a root route for testing
@app.route('/')
def home():
    try:
        conn = get_db_connection()
        if conn is None:
            return jsonify({"status": "error", "message": "Could not connect to database"}), 500
        
        cur = conn.cursor()
        cur.execute('SELECT COUNT(*) as count FROM greek_texts')
        count = cur.fetchone()['count']
        conn.close()
        
        return jsonify({
            "status": "ok",
            "database_path": DB_PATH,
            "passages_count": count,
            "endpoints": [
                "/api/generate-question",
                "/api/random-sequence"
            ]
        })
    except Exception as e:
        logger.error(f"Error in home route: {str(e)}")
        return jsonify({"status": "error", "message": str(e)}), 500

@app.route('/api/generate-question', methods=['POST'])
def generate_question():
    try:
        data = request.get_json()
        difficulty = data.get('difficulty', 'intermediate')
        question_type = data.get('type', 'comprehension')
        
        logger.debug(f"Generating question with difficulty: {difficulty}, type: {question_type}")
        
        # Get random passage and generate question
        passage = question_generator.get_random_passage()
        if not passage:
            logger.error("No passage retrieved from database")
            return jsonify({"error": "No passages available in database"}), 500
        
        logger.debug(f"Selected passage: {passage[:100]}...")
        
        question = question_generator.generate_question(passage, difficulty, question_type)
        if not question:
            logger.error("Failed to generate question")
            return jsonify({"error": "Failed to generate question"}), 500
        
        return jsonify(question)
        
    except Exception as e:
        logger.error(f"Error generating question: {str(e)}")
        return jsonify({"error": f"Failed to generate question: {str(e)}"}), 500

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
        conn = get_db_connection()
        if conn is None:
            return jsonify({"error": "Database connection failed"}), 500
            
        cur = conn.cursor()
        cur.execute('SELECT passage FROM greek_texts ORDER BY RANDOM() LIMIT 1')
        result = cur.fetchone()
        conn.close()
        
        if result:
            return jsonify({"sequence": result['passage']})
        else:
            return jsonify({"error": "No sequences available"}), 500
            
    except Exception as e:
        logger.error(f"Error getting random sequence: {str(e)}")
        return jsonify({"error": str(e)}), 500

@app.route('/api/check-answer', methods=['POST'])
def check_answer():
    """Check if the submitted answer is correct"""
    try:
        data = request.get_json()
        if not data:
            return jsonify({"error": "No data provided"}), 400
            
        question_data = data.get('question')
        selected_answer = data.get('selected_answer')
        
        if not question_data or selected_answer is None:
            return jsonify({"error": "Missing question data or answer"}), 400
            
        correct = question_data.get('correct_answer') == selected_answer
        
        return jsonify({
            "correct": correct,
            "explanation": question_data.get('explanation'),
            "correct_answer": question_data.get('correct_answer')
        })
    except Exception as e:
        return jsonify({"error": f"Server error: {str(e)}"}), 500

if __name__ == '__main__':
    print("🚀 Server starting on http://localhost:5000")
    app.run(host='0.0.0.0', port=5000, debug=True) 