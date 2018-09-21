FROM golang:alpine

COPY src/main.go src/

WORKDIR src/

CMD ["go", "run", "main.go"]

