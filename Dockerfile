FROM golang:latest as builder
ADD . /src/app
WORKDIR /src/app
RUN CGO_ENABLED=0 GOOS=linux go build -o service ./cmd/service/main.go

FROM alpine:edge
COPY --from=builder /src/app/service /service
COPY ./config/config.yaml /config/config.yaml
RUN chmod +x ./service
EXPOSE 8080
ENTRYPOINT ["/service"]