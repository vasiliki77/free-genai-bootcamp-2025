require_relative '../spec_helper'

RSpec.describe 'Words API' do
  let(:base_url) { 'http://localhost:8080/api' }

  describe 'GET /words' do
    it 'returns paginated words list matching spec format' do
      response = HTTParty.get("#{base_url}/words?page=1")
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
          'items_per_page' => 100  # Spec requires 100 items per page
        )
      )
    end

    it 'handles invalid page parameters' do
      response = HTTParty.get("#{base_url}/words?page=0&per_page=-1")
      expect(response.code).to eq(400)
    end
  end

  describe 'GET /words/:id' do
    it 'returns word details matching spec format' do
      response = HTTParty.get("#{base_url}/words/1")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to match(
        'id' => 1,
        'ancient_greek' => be_kind_of(String),
        'greek' => be_kind_of(String),
        'english' => be_kind_of(String),
        'parts' => include(
          'present' => be_kind_of(String),
          'future' => be_kind_of(String),
          'aorist' => be_kind_of(String),
          'perfect' => be_kind_of(String)
        ),
        'stats' => match(
          'correct_count' => be_kind_of(Integer),
          'wrong_count' => be_kind_of(Integer)
        ),
        'groups' => array_including(
          match(
            'id' => be_kind_of(Integer),
            'name' => be_kind_of(String)
          )
        )
      )
    end

    it 'returns 404 for non-existent word' do
      response = HTTParty.get("#{base_url}/words/999")
      expect(response.code).to eq(404)
    end

    it 'handles invalid word ID format' do
      response = HTTParty.get("#{base_url}/words/invalid")
      expect(response.code).to eq(400)
    end
  end
end 