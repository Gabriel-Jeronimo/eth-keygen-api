# eth-keygen-api

Este é um projeto Go que um serviço serveless que assina as transações de acordo com as mensagens recebidas no SQS.

## Como Executar

1. Clone o repositório do projeto:

```bash
git clone https://github.com/seu-usuario/eth-keygen-api.git
```

```bash
cd eth-keygen-api
```

2. Insira suas variáveis no template.yaml

3. Execute o deploy com o SAM:

```bash
sam build && sam deploy --capabilities CAPABILITY_NAMED_IAM
```

Agora, a API estará disponível na AWS.
