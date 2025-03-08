require_relative '../spec_helper'

RSpec.describe 'Groups API' do
  let(:base_url) { 'http://localhost:8080/api' }

  describe 'GET /groups' do
    it 'returns paginated groups list matching spec format' do
      response = HTTParty.get("#{base_url}/groups?page=1")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to match(
        'items' => array_including(
          match(
            'id' => be_kind_of(Integer),
            'name' => be_kind_of(String),
            'word_count' => be_kind_of(Integer)
          )
        ),
        'pagination' => match(
          'current_page' => 1,
          'total_pages' => be_kind_of(Integer),
          'total_items' => be_kind_of(Integer),
          'items_per_page' => 100
        )
      )
    end
  end

  describe 'GET /groups/:id' do
    it 'returns group details matching spec format' do
      response = HTTParty.get("#{base_url}/groups/1")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to match(
        'id' => 1,
        'name' => be_kind_of(String),
        'stats' => match(
          'total_word_count' => be_kind_of(Integer)
        )
      )
    end

    it 'returns 404 for non-existent group' do
      response = HTTParty.get("#{base_url}/groups/999")
      expect(response.code).to eq(404)
    end
  end

  describe 'GET /groups/:id/words' do
    it 'returns paginated words for a group matching spec format' do
      response = HTTParty.get("#{base_url}/groups/1/words?page=1")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to match(
        'items' => array_including(
          match(
            'id' => be_kind_of(Integer),
            'ancient_greek' => be_kind_of(String),
            'greek' => be_kind_of(String),
            'english' => be_kind_of(String),
            'parts' => include(
              'present' => be_kind_of(String),
              'future' => be_kind_of(String),
              'aorist' => be_kind_of(String),
              'perfect' => be_kind_of(String)
            ),
            'correct_count' => be_kind_of(Integer),
            'wrong_count' => be_kind_of(Integer)
          )
        ),
        'pagination' => match(
          'current_page' => 1,
          'total_pages' => be_kind_of(Integer),
          'total_items' => be_kind_of(Integer),
          'items_per_page' => 100
        )
      )
    end
  end

  describe 'GET /groups/:id/study_sessions' do
    it 'returns paginated study sessions for a group' do
      response = HTTParty.get("#{base_url}/groups/1/study_sessions?page=1")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to include('items', 'pagination')
      expect(body['items'].first).to include(
        'id',
        'created_at',
        'study_activity_id'
      )
    end
  end
end 