require_relative '../spec_helper'

RSpec.describe 'Study Sessions API' do
  describe '/api/study_sessions' do
    it 'returns a successful response and a list of study sessions' do
      response = HTTParty.get('http://localhost:8080/api/study_sessions') # Adjust URL if needed

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)

      expect(json_body).to be_an(Array)
      if json_body.any?
        first_session = json_body.first
        expect(first_session).to be_a(Hash)
        expect(first_session).to have_key('id')
        expect(first_session).to have_key('group_id')
        expect(first_session).to have_key('study_activity')
        # ... Add more assertions based on your study session object structure
      end
    end
  end

  describe '/api/study_sessions/{id}' do # Example for getting a specific study session by ID
    it 'returns a successful response and study session details for a valid ID' do
      session_id = 1 # Replace with a valid session ID
      response = HTTParty.get("http://localhost:8080/api/study_sessions/#{session_id}") # Adjust URL

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)
      expect(json_body).to be_a(Hash)
      expect(json_body).to have_key('id')
      expect(json_body).to have_key('group_id')
      expect(json_body).to have_key('study_activity')
      # ... Add more assertions for specific study session details
    end

    it 'returns a 404 Not Found for an invalid study session ID' do
      invalid_session_id = 9999 # Replace with an invalid session ID
      response = HTTParty.get("http://localhost:8080/api/study_sessions/#{invalid_session_id}") # Adjust URL
      expect(response.code).to eq(404) # Or the appropriate error status code
    end
  end

  describe '/api/study_sessions/{sessionId}/words' do # Example for getting words in a study session
    it 'returns a successful response and a list of words reviewed in a session' do
      session_id = 1 # Replace with a valid session ID
      response = HTTParty.get("http://localhost:8080/api/study_sessions/#{session_id}/words") # Adjust URL

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)
      expect(json_body).to be_an(Array)
      # ... Add assertions to check the words in the study session
    end
  end

  # Add more 'describe' blocks for POST, PUT, DELETE study sessions endpoints if applicable, and for ReviewWordInStudySession endpoint
end 