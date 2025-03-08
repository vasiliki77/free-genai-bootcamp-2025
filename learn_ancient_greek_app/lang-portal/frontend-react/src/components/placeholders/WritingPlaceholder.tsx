import React, { useState } from 'react';

const WritingPlaceholder = () => {
  const [isLoading, setIsLoading] = useState(true);

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-3xl font-bold mb-4">Ancient Greek Writing Practice</h1>
      
      {isLoading && (
        <div className="w-full flex justify-center my-8">
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-purple-500"></div>
        </div>
      )}
      
      <div className="w-full h-screen-90 border rounded-lg overflow-hidden">
        <iframe 
          src="http://localhost:8502" 
          title="Ancient Greek Writing Practice"
          className="w-full h-full"
          onLoad={() => setIsLoading(false)}
          style={{ minHeight: '800px' }}
        />
      </div>
    </div>
  );
};

export default WritingPlaceholder;
