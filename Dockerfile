FROM golang:1.17-alpine as build

WORKDIR /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=linux CGO_ENABLED=0 go build -o /bit-driver-matching-service ./cmd/

FROM alpine
COPY --from=build /bit-driver-matching-service ./app

COPY ./service_config.yaml ./
EXPOSE 8080

CMD [ "./app" ]