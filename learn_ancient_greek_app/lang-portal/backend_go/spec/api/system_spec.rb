require_relative '../spec_helper'

RSpec.describe 'System API' do
  let(:base_url) { 'http://localhost:8080/api' }

  describe 'POST /reset_history' do
    it 'resets study history' do
      response = HTTParty.post("#{base_url}/reset_history")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to include(
        'success' => true,
        'message' => 'Study history has been reset'
      )
    end
  end

  describe 'POST /full_reset' do
    it 'performs full system reset' do
      response = HTTParty.post("#{base_url}/full_reset")
      body = JSON.parse(response.body)

      expect(response.code).to eq(200)
      expect(body).to include(
        'success' => true,
        'message' => 'System has been fully reset'
      )
    end
  end
end 