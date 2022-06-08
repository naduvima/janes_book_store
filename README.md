# Jane's Book Store
Working directory janes_book_store
Jane's local book store open to the world for publishing &amp; reading


Do it in Docker way!


Commands

Do you have ruby installed?

curl localhost:8000/books/publish --header "author: jane" --header "token: eyJhbGciOiJIUzI1NiJ9.eyJ0aXRsZSI6IkZpZnR5IHNoYWRlcyBvZiBncmF5IiwiYXV0aG9yIjoiamFuZSJ9.Zx9rGgvYYlqojy38PI-kz_yE4pm0uvCPKdIxw7yvUrY"  -d '{"title": "Fifty shades of gray", "author": "jane", "book_id": 1234}'
payload = {"title": "Fifty shades of gray", "author": "jane", "book_id": 1234}
token = JWT.encode payload, 'eyJkYXRhIjoidGVzdCJ9', 'HS256'
=> "eyJhbGciOiJIUzI1NiJ9.eyJ0aXRsZSI6IkZpZnR5IHNoYWRlcyBvZiBncmF5IiwiYXV0aG9yIjoiamFuZSJ9.Zx9rGgvYYlqojy38PI-kz_yE4pm0uvCPKdIxw7yvUrY"
irb(main):037:0>


Authentication Strategy is encrypt and decrypt by users own password ( non expiring)

(1) Encrypt data in certain way with Author's(user) own given password using JWT
(2) System retreives users stored password and try to decrypt , unsuccessful decrypt is Unauthorized attempt


