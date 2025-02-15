require_relative '../spec_helper'

RSpec.describe 'Dashboard API' do
  let(:base_url) { 'http://localhost:8080/api' }

  describe 'GET /dashboard/last_study_session' do
    it 'returns last study session with correct format' do
      response = HTTParty.get("#{base_url}/dashboard/last_study_session")
      body = JSON.parse(response.body)

      # Match exact format from specs
      expect(response.code).to eq(200)
      expect(body).to match(
        'id' => be_kind_of(Integer),
        'group_id' => be_kind_of(Integer),
        'group_name' => be_kind_of(String),
        'study_activity_id' => be_kind_of(Integer),
        'created_at' => match(/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$/) # ISO8601 format
      )
    end

    it 'returns 404 when no sessions exist' do
      # First, clear all study sessions
      HTTParty.post("#{base_url}/reset_history")
      
      response = HTTParty.get("#{base_url}/dashboard/last_study_session")
      expect(response.code).to eq(404)
    end
  end

  describe 'GET /dashboard/study_progress' do
    it 'returns study progress with correct format' do
      response = HTTParty.get("#{base_url}/dashboard/study_progress")
      body = JSON.parse(response.body)

      # Match exact format from specs
      expect(response.code).to eq(200)
      expect(body).to match(
        'total_words_studied' => be_kind_of(Integer),
        'total_available_words' => be_kind_of(Integer)
      )
    end
  end

  describe 'GET /dashboard/quick-stats' do
    it 'returns quick stats with correct format' do
      response = HTTParty.get("#{base_url}/dashboard/quick-stats")
      body = JSON.parse(response.body)

      # Match exact format from specs
      expect(response.code).to eq(200)
      expect(body).to match(
        'success_rate' => be_kind_of(Float),
        'total_study_sessions' => be_kind_of(Integer),
        'total_active_groups' => be_kind_of(Integer),
        'study_streak_days' => be_kind_of(Integer)
      )

      # Additional validations from specs
      expect(body['success_rate']).to be_between(0.0, 100.0)
    end
  end
end 