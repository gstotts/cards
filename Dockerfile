FROM golang:1.17-alpine AS GO_BUILD
COPY . /server
WORKDIR /server/server
RUN go build -o /go/bin/server/server

FROM alpine:latest
WORKDIR /app
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server