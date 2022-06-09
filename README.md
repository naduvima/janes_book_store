# Jane's Book Store
Working directory janes_book_store
Jane's local book store open to the world for publishing &amp; reading

Do it in Docker way!

Note: THE FILES SHOULD BE DOWNLADED TO DIRECTORY NAMED janes_book_store
git clone https://github.com/naduvima/janes_book_store


How to install Application
1. Must:- you should have docker on your desktop/machine

2. All files in directory named janes_book_store

3. Create all image, build the application
docker-compose up --build

4. Run setup script by following command
docker exec -it janes-books_api psql -f  book_data_store/janes_books.sql

5. Use curl commands to test the application

CURL commands
get status
curl localhost:8080/status

get all books for author Jane
curl localhost:8080/books/author/Jane

[{"book":{"title":"Fire in the Winter","author_id":0,"book_id":0,"description":"","image_s3_url":"","price":32},"author":{"author_id":0,"password":"","author_name":"Jane"}},{"book":{"title":"Fire in the Winter","author_id":0,"book_id":0,"description":"","image_s3_url":"","price":32},"author":{"author_id":0,"password":"","author_name":"Jane"}}]


Creating a jwt token using ruby and jwt gem 
Use cliet/client.rb to send publish or unpublish routes, changing th url.

Encryption Technique , using ruby and jwt
payload = { }
token = JWT.encode payload, 'eyJkYXRhIjoidGVzdCJ9', 'HS256'

USing curl:-
curl localhost:8000/books/publish --header "author: jane" --header "token: $token"  -d "$payload"
payload = { }

To view details of a title
localhost:8080/books/details?title=Fire in the Winter


Authentication Strategy is encrypt and decrypt by users own password ( non expiring)

(1) Encrypt data in certain way with Author's(user) own given password using JWT
(2) System retreives users stored password and try to decrypt , unsuccessful decrypt is Unauthorized attempt



