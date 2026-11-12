FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o /app .

FROM alpine:3.21
RUN adduser -D -u 10001 app
COPY --from=build /app /app
EXPOSE 8080
USER app
ENTRYPOINT ["/app"]
