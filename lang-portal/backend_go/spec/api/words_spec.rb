require_relative '../spec_helper'

RSpec.describe 'Words API' do
  describe '/api/words' do
    it 'returns a successful response and a list of words' do
      response = HTTParty.get('http://localhost:8080/api/words') # Adjust URL if needed

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)

      expect(json_body).to be_an(Array)
      if json_body.any?
        first_word = json_body.first
        expect(first_word).to be_a(Hash)
        expect(first_word).to have_key('id')
        expect(first_word).to have_key('ancient_greek')
        expect(first_word).to have_key('greek')
        expect(first_word).to have_key('english')
        # ... Add more assertions based on your word object structure
      end
    end
  end

  describe '/api/words/{id}' do # Example for getting a specific word by ID
    it 'returns a successful response and word details for a valid ID' do
      word_id = 1 # Replace with a valid word ID
      response = HTTParty.get("http://localhost:8080/api/words/#{word_id}") # Adjust URL

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)
      expect(json_body).to be_a(Hash)
      expect(json_body).to have_key('id')
      expect(json_body).to have_key('ancient_greek')
      expect(json_body).to have_key('greek')
      expect(json_body).to have_key('english')
      # ... Add more assertions for specific word details
    end

    it 'returns a 404 Not Found for an invalid word ID' do
      invalid_word_id = 9999 # Replace with an invalid word ID
      response = HTTParty.get("http://localhost:8080/api/words/#{invalid_word_id}") # Adjust URL
      expect(response.code).to eq(404) # Or the appropriate error status code
    end
  end

  describe '/api/words?groupId={groupId}' do # Example for filtering words by group
    it 'returns a successful response and a list of words for a specific group' do
      group_id = 1 # Example group ID, adjust as needed
      response = HTTParty.get("http://localhost:8080/api/words?groupId=#{group_id}") # Adjust URL

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)

      expect(json_body).to be_an(Array)
      # You might add assertions to check if the words belong to the specified group
    end
  end

  # Add more 'describe' blocks for POST, PUT, DELETE words endpoints if applicable
end 