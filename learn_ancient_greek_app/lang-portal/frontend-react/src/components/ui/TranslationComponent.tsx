import React, { useState } from 'react';
import axios from 'axios';

const TranslationComponent: React.FC = () => {
  const [inputText, setInputText] = useState('');
  const [translation, setTranslation] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleTranslate = async (): Promise<void> => {
    if (!inputText.trim()) return;
    
    setIsLoading(true);
    setError('');
    setTranslation('');
    
    try {
      const response = await axios.post(
        `/api/translate?sentence=${encodeURIComponent(inputText)}`,
        null,
        {
          headers: {
            Authorization: 'Bearer KLEIDI'
          }
        }
      );

      // Parse the full response to extract the content between the last <START> and <END> tags
      const fullResponse = response.data.full_response;
      const startTagIndex = fullResponse.lastIndexOf('<START>');
      const endTagIndex = fullResponse.lastIndexOf('<END>');
      
      if (startTagIndex !== -1 && endTagIndex !== -1 && startTagIndex < endTagIndex) {
        const extractedTranslation = fullResponse.substring(startTagIndex + 7, endTagIndex);
        setTranslation(extractedTranslation);
      } else {
        // Fallback to the translation field if available, otherwise show an error
        if (response.data.translation) {
          setTranslation(response.data.translation);
        } else {
          setError('Could not extract translation from response');
        }
      }
    } catch (err) {
      setError('Failed to translate: ' + (err instanceof Error ? err.message : String(err)));
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="translation-container p-6 max-w-2xl mx-auto">
      <h2 className="text-2xl font-bold mb-4">English to Ancient Greek Translation</h2>
      
      <div className="input-section mb-6">
        <label htmlFor="english-input" className="block mb-2 font-medium">
          Enter English Text:
        </label>
        <textarea
          id="english-input"
          className="w-full p-3 border rounded-md min-h-[100px]"
          value={inputText}
          onChange={(e) => setInputText(e.target.value)}
          placeholder="Type English text here..."
        />
      </div>
      
      <div className="button-section mb-6">
        <button
          className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:bg-gray-400"
          onClick={handleTranslate} 
          disabled={isLoading || !inputText.trim()}
        >
          {isLoading ? 'Translating...' : 'Translate'}
        </button>
      </div>
      
      {error && <div className="error-message text-red-600 mb-4">{error}</div>}
      
      {translation && (
        <div className="result-section bg-gray-100 p-4 rounded-md">
          <h3 className="text-xl font-semibold mb-2">Ancient Greek Translation:</h3>
          <div className="translation-result text-lg font-medium">{translation}</div>
        </div>
      )}
    </div>
  );
};

export default TranslationComponent;