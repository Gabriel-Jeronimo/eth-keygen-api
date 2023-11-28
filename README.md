# eth-keygen-api

Este é um projeto Go que um serviço serveless que assina as transações de acordo com as mensagens recebidas no SQS.

## Como Executar

1. Clone o repositório do projeto:

```bash
git clone https://github.com/Gabriel-Jeronimo/eth-keygen-api.git
```

```bash
cd eth-keygen-api
```

2. Faça o download das dependências do Terraform

```bash
make init
```

3. Execute o deploy com o Terraform

```bash
make apply
```

Pronto, a API estará disponível na AWS.

![Foto da arquitetura](https://github.com/Gabriel-Jeronimo/eth-keygen-api/assets/55462130/351cb4b8-47bb-4444-b2d3-a6b52001b20b)
