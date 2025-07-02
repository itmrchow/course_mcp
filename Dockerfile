# build
FROM golang:1.24.3 AS builder
WORKDIR /app
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x
COPY . /app
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target="/root/.cache/go-build" \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o course-mcp
# run
FROM alpine:3.19.2
WORKDIR /app
COPY --from=builder /app/course-mcp /app
COPY --from=builder /app/config.yaml /app  
ENTRYPOINT ["./course-mcp"]