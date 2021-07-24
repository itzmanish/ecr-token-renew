FROM golang:alpine as build

WORKDIR /app
COPY . .
RUN go build -o ecr-token-renew ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /app/ecr-token-renew .
CMD ["./ecr-token-renew"]
