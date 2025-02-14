require_relative '../spec_helper'

RSpec.describe 'Study Activities API' do
  let(:base_url) { 'http://localhost:8080/api' }

  describe 'GET /study_activities/:id' do
    it 'returns study activity details' do
      response = HTTParty.get("#{base_url}/study_activities/1")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to include(
        'id' => 1,
        'name' => be_kind_of(String),
        'description' => be_kind_of(String),
        'thumbnail_url' => be_kind_of(String)
      )
    end

    it 'returns 404 for non-existent activity' do
      response = HTTParty.get("#{base_url}/study_activities/999")
      expect(response.code).to eq(404)
    end
  end

  describe 'GET /study_activities/:id/study_sessions' do
    it 'returns paginated study sessions for an activity' do
      response = HTTParty.get("#{base_url}/study_activities/1/study_sessions?page=1")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to include(
        'items',
        'current_page',
        'total_pages',
        'total_items',
        'items_per_page'
      )
      expect(body['items'].first).to include(
        'id',
        'group_id',
        'created_at'
      )
    end
  end
end 