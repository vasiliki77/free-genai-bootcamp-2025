require_relative '../spec_helper'

RSpec.describe 'Study Activities API' do
  describe '/api/study_activities' do
    it 'returns a successful response and a list of study activities' do
      response = HTTParty.get('http://localhost:8080/api/study_activities') # Adjust URL if needed

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)

      expect(json_body).to be_an(Array)
      if json_body.any?
        first_activity = json_body.first
        expect(first_activity).to be_a(Hash)
        expect(first_activity).to have_key('id')
        # ... Add more assertions based on your study activity object structure
      end
    end
  end

  describe '/api/study_activities/{id}' do # Example for getting a specific study activity by ID
    it 'returns a successful response and study activity details for a valid ID' do
      activity_id = 1 # Replace with a valid activity ID
      response = HTTParty.get("http://localhost:8080/api/study_activities/#{activity_id}") # Adjust URL

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)
      expect(json_body).to be_a(Hash)
      expect(json_body).to have_key('id')
      # ... Add more assertions for specific study activity details
    end

    it 'returns a 404 Not Found for an invalid study activity ID' do
      invalid_activity_id = 9999 # Replace with an invalid activity ID
      response = HTTParty.get("http://localhost:8080/api/study_activities/#{invalid_activity_id}") # Adjust URL
      expect(response.code).to eq(404) # Or the appropriate error status code
    end
  end

  describe '/api/study_activities/{activityId}/study-sessions' do # Example for getting study sessions for an activity
    it 'returns a successful response and a list of study sessions for an activity' do
      activity_id = 1 # Replace with a valid activity ID
      response = HTTParty.get("http://localhost:8080/api/study_activities/#{activity_id}/study-sessions") # Adjust URL

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)
      expect(json_body).to be_an(Array)
      # ... Add assertions to check the study sessions for the activity
    end
  end

  describe '/api/study_activities/session' do # Example for creating a study activity session (POST)
    it 'returns a successful response when creating a new study activity session' do
      # Assuming you need to send a request body for creating a session
      request_body = { group_id: 1, activity_id: 1 }.to_json # Adjust request body as needed
      headers = { 'Content-Type' => 'application/json' }
      response = HTTParty.post('http://localhost:8080/api/study_activities/session', body: request_body, headers: headers) # Adjust URL and method

      expect(response.code).to be_between(200, 299) # Check for successful status codes (2xx)
      expect(response.headers['Content-Type']).to include('application/json')
      json_body = JSON.parse(response.body)
      expect(json_body).to be_a(Hash) # Or appropriate response type
      # ... Add assertions to check the created study activity session
    end
  end

  # Add more 'describe' blocks for other study activity endpoints (POST, PUT, DELETE if applicable)
end 