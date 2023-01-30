FROM golang:1.19

RUN apt update
RUN apt upgrade -y
RUN apt install sqlite3 -y

WORKDIR /usr/src/graphql-go

ENTRYPOINT go mod tidy && go run cmd/server/server.go

# docker build -t stephano1234/graphql-go .
# docker run --rm -it -v "$(pwd)"/:/usr/src/graphql-go -p 8080:8080 stephano1234/graphql-go