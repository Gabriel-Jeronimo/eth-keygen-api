Sure, here's the translation:

# eth-keygen-api

This is a Go project that provides an API for generating Ethereum keys. It includes a Dockerfile for easy deployment in containers.

You can find an implementation of this service on AWS in the `feat/aws` branch.

## How to Run

1. Clone the project repository:

```bash
git clone https://github.com/Gabriel-Jeronimo/eth-keygen-api.git
```

```bash
cd eth-keygen-api
```

2. Build the Docker image:

```bash
docker build -t eth-keygen-api .
```

3. Run the container:

```bash
docker run -p 8080:8080 eth-keygen-api
```

Now, the API will be available at http://localhost:8080.

## API Usage

The API provides endpoints for generating Ethereum keys. You can make HTTP requests to these endpoints to generate keys.

Example request to generate a key:

```bash
curl --location --request POST 'localhost:8080/keypair'
```

Example request to get the Ethereum address from a key:

```bash
curl --location 'localhost:8080/address?publicKey=044f26d656319657f42a81f9d580365d447acc176e22b283b955b108b8380bcffbe06713fa49534048da5a18cc95a46e7c53b8b61ee15b6f04fccdabec564741a5'
```
