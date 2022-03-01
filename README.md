<h1 align="center">Bi Taksi Driver Matching Service</h1>

![example workflow](https://github.com/ashkan90/bit-driver-matching-service/actions/workflows/main.yml/badge.svg)

> This project aims to find nearest drivers across given customer location. To do that, It's using another service which called 'driver-location-service'

# Hexagonal Architecture
The idea of Hexagonal Architecture is to put inputs and outputs at the edges of our design. Business logic should not depend on whether we expose a REST or a GraphQL API, and it should not depend on where we get data from — a database, a microservice API exposed via gRPC or REST, or just a simple CSV file. <br>

The pattern allows us to isolate the core logic of our application from outside concerns. Having our core logic isolated means we can easily change data source details without a significant impact or major code rewrites to the codebase.

## Run on Local Machine

```shell
docker build -t bit-driver-matching-service .
docker run -p 8082:8080 bit-driver-matching-service
```

## Endpoints

```console
GET /find-nearest 
```

## Example Usages
```shell
# Request
curl --location --request GET 'http://localhost:8083/find-nearest' \
--header 'Content-Type: application/json' \
--data-raw '{
    "longitude": 40.94289771,
    "latitude": 28.0390297
}'
```

```json
// Response
{
  "coordinates": [
    40.946104,
    28.035588
  ]
}
```

## Introducing Dependencies
- For mocking I choose 'mocgken' package. It's official package and I can easily find a solution for that I'm facing a problem right now.
- For unit testing I choose 'testify' package. It suits my usecase very-well and It's very simple to use tho. Also popular enough to worry about community-things.
- For server framework I choose 'echo' package. It has great documentation and It's coming with built-in examples. I might use fasthttp over echo but fasthttp has some internal problems such as response object is not designed for concurrent usage.


## Deployment

> The deployment is not in scope but if I wanted to do it then surely I use `github actions for heroku`

## Author

👤 **Emirhan Ataman**


## 📝 License

Copyright © 2022 [Emirhan Ataman](https://github.com/ashkan90). <br />
