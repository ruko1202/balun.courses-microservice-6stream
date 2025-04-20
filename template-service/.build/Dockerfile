FROM goreleaser/goreleaser:v2.8.2 as goreleaser
RUN ls -l /usr/bin | grep gore


FROM golang:1.23.7-alpine3.21 as builder

RUN apk update && \
    apk add --no-cache git make

# Mark source directory as safe to avoid problem with dubious ownership in git while building image.
RUN git config --global --add safe.directory /go/src/app

WORKDIR /app
COPY --from=goreleaser /usr/bin/goreleaser /app/bin/goreleaser

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN GOBIN=/app/bin make build


FROM alpine:3.21

EXPOSE 8080

WORKDIR /app
COPY --from=builder /app/bin/app ./
ENTRYPOINT ["./app"]
