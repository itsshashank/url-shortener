# url-shortener


## How to use

Create new short url:

```bash
curl -X POST http://localhost:8080/short \
  -H 'Content-Type: application/json' \
  -d '{
        "url": "https://github.com"
}'
```

The response:

```json
{"ShortUrl":"localhost:8080/a3ebZ2l","message":"short url created successfully"}
```


Get redirected:

```bash
 curl -X GET localhost:8080/a3ebZ2l
```

The response:

```bash
<a href="https://github.com">Permanent Redirect</a>.
```


Get ShortURL info
```bash
curl -X GET localhost:8080/a3ebZ2l/info
```

The response:
```json
{"ShortUrl":"localhost:8080/a3ebZ2l","longUrl":"https://github.com"}
```