# Handling gRPC Errors in a REST API Gateway(POC)

Learning how to handle errors sent out by a gRPC server in a REST Gateway and sending meaningful error responses to clients

The REST API Gateway is running on `localhost:8082` and the gRPC server is running on `localhost:8081`.

## How to run

1. Start the gRPC server

```bash
make start_grpc
```

2. Start the Gateway

```bash
make start_gateway
```

## gRPC Server

```proto
service CalculatorService {
    rpc DivideNumbers(DivideNumbersRequest) returns (DivideNumbersResponse);
}

// Divide number_one by number_two
message DivideNumbersRequest {
    float number_one = 1;
    float number_two = 2;
}

message DivideNumbersResponse {
    int32 code = 1;
    float answer = 2;
}
```

## REST API Gateway

Endpoint: `/calculator/divide`

Method: `POST`

Request Body:

```json
{
  "number_one": 10,
  "number_two": 0
}
```

Success Response:

```json
{
  "code": 200,
  "answer": 0.01
}
```

Error Response:

```json
// Errors out when number_two is 0 (Division by zero)
{
  "error": true,
  "message": "division by zero not possible"
}
```
