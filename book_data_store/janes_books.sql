CREATE DATABASE janes_books_db;
\connect janes_books_db;
CREATE TABLE authors (
    author_id SERIAL PRIMARY KEY,
    password character varying(255),
    author_name character varying(255) UNIQUE
);
CREATE UNIQUE INDEX name ON authors (author_id);
CREATE TABLE books (
    books_id SERIAL PRIMARY KEY,
    author_id SERIAL REFERENCES authors,
    title character varying(255),
    created_at timestamp without time zone NOT NULL default current_timestamp,
    description character varying(255),
    image_s3_url character varying(255), --should store in s3
    price numeric(15,4)
);

--Test values
INSERT INTO authors(password, author_name) VALUES('eyJkYXRhIjoidGVzdCJ9', 'Lucy score');

INSERT INTO authors(password, author_name) VALUES('eyJkYXRhIjoidGVzdCJ9', 'Dave Gerad');
INSERT INTO authors(password, author_name) VALUES('eyJkYXRhIjoidGVzdCJ9', 'Jane');
