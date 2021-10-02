# shurl

Short URL.

Simple practice url shortener to prepare for an interview.

## What

Single route:

- /shurl/

### PUT

`PUT /shurl/?url=https://github.com/Karitham/shurl` => Returns a 6 letter word that's mapped to a URL

### GET

`GET /shurl/ABCD` => Redirects to whatever was provided with PUT
