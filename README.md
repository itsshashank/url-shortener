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

## How to Run in Docker

Build Docker from file

```bash
docker build -f hack/Dockerfile . -t urlshort
```

Run 
```bash
docker run -p 8000:8080 urlshort
```

Can access on localhost:8000 or at containerip:8080
```bash
docker inspect 480cf1308d0d | grep "IPAddress"
```
```bash
➜  url-shortener git:(main) ✗ docker inspect c2499030d285 | grep "IPAddress"
            "SecondaryIPAddresses": null,
            "IPAddress": "172.17.0.3",
                    "IPAddress": "172.17.0.3",

```