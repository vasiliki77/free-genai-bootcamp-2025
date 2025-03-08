import json
import sqlite3
from typing import Dict, List, Optional, Any
import os
from dotenv import load_dotenv
import logging
import random

# Load environment variables
load_dotenv()

# Set up database path
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
DB_PATH = os.path.join(BASE_DIR, 'backend', 'listening-learning.db')

logger = logging.getLogger(__name__)

class GreekQuestionGenerator:
    def __init__(self, db_path: str):
        self.db_path = db_path
        logger.info("Initialized Greek Question Generator")
        
        # Define common question patterns
        self.question_patterns = [
            {
                "type": "word_meaning",
                "template": "What does '{word}' mean?",
                "words": [
                    {"word": "μῆνιν", "meaning": "wrath", "options": ["wrath", "love", "peace", "joy"]},
                    {"word": "θεὰ", "meaning": "goddess", "options": ["goddess", "god", "hero", "mortal"]},
                    {"word": "ἄειδε", "meaning": "sing", "options": ["sing", "dance", "write", "speak"]},
                    {"word": "οὐλομένην", "meaning": "destructive", "options": ["destructive", "beautiful", "peaceful", "wise"]},
                    {"word": "ἄλγε", "meaning": "pains", "options": ["pains", "joys", "dreams", "hopes"]}
                ]
            },
            {
                "type": "grammar_form",
                "template": "What type of word is '{word}'?",
                "words": [
                    {"word": "μῆνιν", "answer": "noun (accusative)", "options": ["noun (accusative)", "verb", "adjective", "adverb"]},
                    {"word": "ἄειδε", "answer": "verb (imperative)", "options": ["verb (imperative)", "noun", "pronoun", "particle"]},
                    {"word": "θεὰ", "answer": "noun (vocative)", "options": ["noun (vocative)", "verb", "adjective", "pronoun"]}
                ]
            },
            {
                "type": "phrase_meaning",
                "template": "What does the phrase '{phrase}' mean?",
                "phrases": [
                    {
                        "phrase": "μῆνιν ἄειδε θεὰ",
                        "meaning": "sing, goddess, of the wrath",
                        "options": [
                            "sing, goddess, of the wrath",
                            "praise the immortal gods",
                            "tell me about the war",
                            "speak of ancient times"
                        ]
                    },
                    {
                        "phrase": "Διὸς δ᾽ ἐτελείετο βουλή",
                        "meaning": "and the will of Zeus was accomplished",
                        "options": [
                            "and the will of Zeus was accomplished",
                            "the gods were pleased",
                            "the heroes fought bravely",
                            "the battle was won"
                        ]
                    }
                ]
            }
        ]

    def get_random_passage(self) -> Optional[str]:
        try:
            conn = sqlite3.connect(self.db_path)
            cursor = conn.cursor()
            cursor.execute('SELECT passage FROM greek_texts ORDER BY RANDOM() LIMIT 1')
            result = cursor.fetchone()
            conn.close()
            
            if result:
                return result[0]
            else:
                logger.error("No passages found in database")
                return None
                
        except Exception as e:
            logger.error(f"Error getting random passage: {str(e)}")
            return None

    def generate_question(self, passage: str, difficulty: str = 'intermediate', 
                         question_type: str = 'comprehension') -> Optional[Dict[str, Any]]:
        try:
            # Choose a random question pattern
            pattern = random.choice(self.question_patterns)
            
            if pattern["type"] == "word_meaning":
                # Generate a word meaning question
                word_data = random.choice(pattern["words"])
                question_data = {
                    "question": pattern["template"].format(word=word_data["word"]),
                    "options": word_data["options"],
                    "correct_answer": 1,  # First option is always correct in our word list
                    "explanation": f"'{word_data['word']}' means '{word_data['meaning']}' in English.",
                    "word": word_data["word"],
                    "translation": word_data["meaning"]
                }
            
            elif pattern["type"] == "grammar_form":
                # Generate a grammar form question
                word_data = random.choice(pattern["words"])
                question_data = {
                    "question": pattern["template"].format(word=word_data["word"]),
                    "options": word_data["options"],
                    "correct_answer": 1,  # First option is always correct in our word list
                    "explanation": f"'{word_data['word']}' is a {word_data['answer']}.",
                    "word": word_data["word"],
                    "translation": word_data["answer"]
                }
            
            else:  # phrase_meaning
                # Generate a phrase meaning question
                phrase_data = random.choice(pattern["phrases"])
                question_data = {
                    "question": pattern["template"].format(phrase=phrase_data["phrase"]),
                    "options": phrase_data["options"],
                    "correct_answer": 1,  # First option is always correct in our phrase list
                    "explanation": f"The phrase translates to: {phrase_data['meaning']}",
                    "word": phrase_data["phrase"],
                    "translation": phrase_data["meaning"]
                }
            
            return {
                "passage": passage,
                "question": json.dumps(question_data),
                "type": pattern["type"],
                "difficulty": difficulty
            }
            
        except Exception as e:
            logger.error(f"Error generating question: {str(e)}")
            return None

    def _create_prompt(self, passage: str, difficulty: str, question_type: str) -> str:
        return ""  # Not used anymore

    def generate_comprehension_question(self, difficulty: str = "intermediate") -> Dict:
        passage = self.get_random_passage()
        if not passage:
            return None
            
        response = self.generate_question(passage, difficulty, 'comprehension')
        if not response:
            return None

        try:
            question_data = json.loads(response['question'])
            question_data["greek_text"] = passage
            return question_data
        except:
            return None

    def get_grammar_question(self, topic: str = "verb forms") -> Dict:
        """Generate a grammar-focused question from a random Greek passage"""
        passage = self.get_random_passage()
        if not passage:
            return None
            
        prompt = f"""Given this Ancient Greek passage, create a grammar question about {topic}.
        
        Passage: {passage}
        
        Create a question that:
        1. Identifies a specific grammar construct in the text
        2. Asks about its form, usage, or meaning
        3. Provides four possible answers
        4. Includes a detailed explanation of the correct answer
        
        Return the response in this JSON format:
        {{
            "grammar_topic": "{topic}",
            "identified_text": "the specific Greek text being asked about",
            "question": "Question about the grammar",
            "options": ["Option 1", "Option 2", "Option 3", "Option 4"],
            "correct_answer": number,
            "explanation": "Grammatical explanation"
        }}
        
        Ensure the response is valid JSON and all fields are present."""

        response = self.generate_question(passage, 'grammar', topic)
        if not response:
            return None

        try:
            question_data = json.loads(response['question'])
            question_data["greek_text"] = passage
            return question_data
        except:
            return {
                "greek_text": passage,
                "grammar_topic": topic,
                "identified_text": "",
                "question": "Error generating grammar question",
                "options": [
                    "Unable to generate options",
                    "Please try again",
                    "System error",
                    "Contact support"
                ],
                "correct_answer": 1,
                "explanation": "Error generating explanation"
            }

    def get_vocabulary_question(self, word: Optional[str] = None) -> Dict:
        """Generate a vocabulary question, optionally focused on a specific word"""
        passage = self.get_random_passage()
        if not passage:
            return None
            
        prompt = f"""Given this Ancient Greek passage, create a vocabulary question.
        {f'Focus on the word: {word}' if word else 'Choose an interesting word from the passage.'}
        
        Passage: {passage}
        
        Create a question that:
        1. Identifies a specific word
        2. Tests understanding of its meaning and usage
        3. Provides four possible translations or usages
        4. Includes etymology or memorable context if relevant
        
        Return the response in this JSON format:
        {{
            "target_word": "the Greek word being asked about",
            "question": "Question about the word",
            "options": ["Option 1", "Option 2", "Option 3", "Option 4"],
            "correct_answer": number,
            "explanation": "Explanation including etymology or usage notes"
        }}
        
        Ensure the response is valid JSON and all fields are present."""

        response = self.generate_question(passage, 'vocabulary', word)
        if not response:
            return None

        try:
            question_data = json.loads(response['question'])
            question_data["greek_text"] = passage
            return question_data
        except:
            return {
                "greek_text": passage,
                "target_word": word or "",
                "question": "Error generating vocabulary question",
                "options": [
                    "Unable to generate options",
                    "Please try again",
                    "System error",
                    "Contact support"
                ],
                "correct_answer": 1,
                "explanation": "Error generating explanation"
            } 