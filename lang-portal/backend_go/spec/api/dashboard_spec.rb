require_relative '../spec_helper'

RSpec.describe 'Dashboard API' do
  describe '/api/dashboard/quick-stats' do
    it 'returns a successful response and quick stats data' do
      response = HTTParty.get('http://localhost:8080/api/dashboard/quick-stats') # Adjust URL if needed

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)

      expect(json_body).to be_a(Hash)
      expect(json_body).to have_key('totalWords')
      expect(json_body).to have_key('groupsCount')
      expect(json_body).to have_key('wordsLearned')
      expect(json_body).to have_key('wordsInProgress')

      expect(json_body['totalWords']).to be_a(Integer) if json_body.key?('totalWords')
      expect(json_body['wordsLearned']).to be_a(Integer) if json_body.key?('wordsLearned')
      expect(json_body['wordsInProgress']).to be_a(Integer) if json_body.key?('wordsInProgress')
      # ... Add more assertions based on your quick stats response structure
    end
  end

  describe '/api/dashboard/last_study_session' do
    it 'returns a successful response and the last study session data' do
      response = HTTParty.get('http://localhost:8080/api/dashboard/last_study_session') # Adjust URL

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)

      expect(json_body).to be_a(Hash) # Assuming it returns a single study session object
      expect(json_body).to have_key('id')
      expect(json_body).to have_key('group_id')
      expect(json_body).to have_key('study_activity_id')
      # ... Add more assertions for the last study session response
    end
  end

  describe '/api/dashboard/study_progress' do
    it 'returns a successful response and study progress data' do
      response = HTTParty.get('http://localhost:8080/api/dashboard/study_progress') # Adjust URL

      expect(response.code).to eq(200)
      expect(response.headers['Content-Type']).to include('application/json')

      json_body = JSON.parse(response.body)

      expect(json_body).to be_a(Hash) # Or Array, depending on your response structure
      # ... Add assertions for study progress data
    end
  end

  # Add more 'describe' blocks for other dashboard endpoints if any
end 