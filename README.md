# eth-keygen-api

This is a serveless service that signs ethereum transactions received via SQS.

## Usage

### Installation

1. Clone the project repository.

```bash
git clone https://github.com/Gabriel-Jeronimo/eth-keygen-api.git
```

```bash
cd eth-keygen-api
```

2. Download terraform dependencies.

```bash
make init
```

3. Deploy using terraform.

```bash
make apply
```

Your API is now ready for use on AWS

![Architecture picture](https://github.com/Gabriel-Jeronimo/eth-keygen-api/assets/55462130/351cb4b8-47bb-4444-b2d3-a6b52001b20b)
