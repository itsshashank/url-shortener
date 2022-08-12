FROM golang

WORKDIR /workspace/url-shortener

COPY ./ .

RUN go mod download && go mod verify

RUN go build -v -o url-shortener ./cmd/

EXPOSE 8080

ENV GIN_MODE=release

CMD ["./url-shortener"]