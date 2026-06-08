# Dockerfile
# Multi-stage build: compile in a Go image, run in a minimal Alpine image.
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o website-operator .

FROM alpine:3.21
COPY --from=builder /app/website-operator /website-operator
CMD ["/website-operator"]