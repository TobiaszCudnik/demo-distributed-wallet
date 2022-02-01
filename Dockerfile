FROM golang:1.17-alpine
# deps
RUN apk add --no-cache curl git protobuf-dev
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin
# src
WORKDIR /app
COPY . .
# run
EXPOSE 8080
CMD ["task",  "demo"]
