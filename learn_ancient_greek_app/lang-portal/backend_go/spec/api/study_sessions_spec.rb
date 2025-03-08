require_relative '../spec_helper'

RSpec.describe 'Study Sessions API' do
  let(:base_url) { 'http://localhost:8080/api' }

  describe 'GET /study_sessions' do
    it 'returns paginated study sessions' do
      response = HTTParty.get("#{base_url}/study_sessions?page=1")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to include(
        'items',
        'current_page',
        'total_pages',
        'total_items',
        'items_per_page'
      )
    end
  end

  describe 'GET /study_sessions/:id' do
    it 'returns study session details' do
      response = HTTParty.get("#{base_url}/study_sessions/1")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to include(
        'id',
        'group_id',
        'study_activity_id',
        'created_at'
      )
    end
  end

  describe 'GET /study_sessions/:id/words' do
    it 'returns paginated words for a study session' do
      response = HTTParty.get("#{base_url}/study_sessions/1/words?page=1")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to include('items', 'pagination')
      expect(body['items'].first).to include(
        'id',
        'ancient_greek',
        'greek',
        'english'
      )
    end
  end
end 