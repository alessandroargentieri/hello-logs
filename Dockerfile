# build stage
FROM golang:1.21 AS build-stage

WORKDIR /build
COPY . /build/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello-logs .

# final stage ~~~~~~~~~~~~~~~~~~~
FROM busybox:latest

COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /app
COPY --from=build-stage /build/hello-logs ./

RUN addgroup -S appgroup && adduser -S --no-create-home appuser -G appgroup
USER appuser

CMD ["./hello-logs"]