FROM golang:1.18.3-alpine3.16 as builder

ENV GOBIN /go/bin
RUN mkdir /go/src/app /go/src/app/bin
WORKDIR /go/src/app
ADD . /go/src/app

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go mod vendor
RUN go build -mod=vendor -a -installsuffix cgo -o /go/src/app/bin/main /go/src/app/cmd/service
RUN go build -mod=vendor -a -installsuffix cgo -o /go/src/app/bin/migrator /go/src/app/cmd/migrator

# Run stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/

COPY --from=builder /go/src/app/bin ./bin
COPY --from=builder /go/src/app/db/migrations /app/db/migrations

EXPOSE 8080
CMD ["./bin/main"]