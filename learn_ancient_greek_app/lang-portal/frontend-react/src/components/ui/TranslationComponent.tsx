import React, { useState } from 'react';
import axios from 'axios';

interface TranslationResponse {
  original: string;
  translated: string;
  error?: string;
}

const TranslationComponent: React.FC = () => {
  const [inputText, setInputText] = useState<string>('');
  const [translation, setTranslation] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');

  const handleTranslate = async (): Promise<void> => {
    if (!inputText.trim()) return;
    
    setIsLoading(true);
    setError('');
    
    try {
      // This should point to your Go backend
      const response = await axios.post<TranslationResponse>('http://localhost:8080/api/translate', {
        text: inputText
      });
      
      if (response.data.error) {
        setError(response.data.error);
      } else {
        setTranslation(response.data.translated);
      }
    } catch (err) {
      setError('Failed to get translation. Please try again later.');
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="translation-container">
      <h2>English to Ancient Greek Translator</h2>
      
      <div className="input-section">
        <textarea
          value={inputText}
          onChange={(e) => setInputText(e.target.value)}
          placeholder="Enter English text to translate..."
          rows={4}
        />
        <button 
          onClick={handleTranslate} 
          disabled={isLoading || !inputText.trim()}
        >
          {isLoading ? 'Translating...' : 'Translate'}
        </button>
      </div>
      
      {error && <div className="error-message">{error}</div>}
      
      {translation && (
        <div className="result-section">
          <h3>Translation:</h3>
          <div className="translation-result">{translation}</div>
        </div>
      )}
    </div>
  );
};

export default TranslationComponent;