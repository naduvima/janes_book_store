require 'jwt'
require 'rest-client'
require 'json'
# gem install jwt first
author = 'jane' #Edit this
password = 'eyJkYXRhIjoidGVzdCJ9' #Edit this
# the above values should exist in author table
# get a token
payload = {"book": {"title": "Fire in the Winter","price": 32.0}, "author": {"author_name": "jane"}}
token = JWT.encode payload, password, 'HS256'

headers = { 'author' => author, 'token' => token }
response = RestClient.post 'http://localhost:8000/books/publish', payload.to_json, headers
puts response.code
puts response.body



