require 'httparty'
require 'json'

RSpec.configure do |config|
  config.before(:suite) do
    # Check if server is running
    begin
      HTTParty.get('http://localhost:8080/api/health')
    rescue Errno::ECONNREFUSED
      puts "\nError: API server must be running on http://localhost:8080"
      exit 1
    end
  end

  config.expect_with :rspec do |expectations|
    expectations.include_chain_clauses_in_custom_matcher_descriptions = true
  end

  config.mock_with :rspec do |mocks|
    mocks.verify_partial_doubles = true
  end

  config.shared_context_metadata_behavior = :apply_to_host_groups
end 