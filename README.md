# eth-keygen-api

Este é um projeto Go que fornece uma API para gerar chaves Ethereum. Ele inclui um Dockerfile para fácil implantação em containers.

## Como Executar

1. Clone o repositório do projeto:

```bash
git clone https://github.com/Gabriel-Jeronimo/eth-keygen-api.git
```

```bash
cd eth-keygen-api
```

2. Construa a imagem Docker:

```bash
docker build -t eth-keygen-api .
```

3. Execute o container:

```bash
docker run -p 8080:8080 eth-keygen-api
```

Agora, a API estará disponível em http://localhost:8080.

## Uso da API

A API oferece endpoints para gerar chaves Ethereum. Você pode fazer solicitações HTTP para esses endpoints para gerar chaves.

Exemplo de solicitação para gerar uma chave:

```bash
curl --location --request POST 'localhost:8080/keypair'
```

Exemplo de solicitação para pegar o endereço ethereum de uma chave:

```bash
curl --location 'localhost:8080/address?publicKey=044f26d656319657f42a81f9d580365d447acc176e22b283b955b108b8380bcffbe06713fa49534048da5a18cc95a46e7c53b8b61ee15b6f04fccdabec564741a5'
```
