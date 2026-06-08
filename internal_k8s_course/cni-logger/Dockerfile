# Dockerfile (create this in the current cni-logger/ directory)
FROM golang:1.26 AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /cni-logger main.go

FROM busybox:1.36
COPY --from=build /cni-logger /cni-logger