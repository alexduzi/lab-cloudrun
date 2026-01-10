# Documentação da API

Este diretório contém a documentação OpenAPI/Swagger para a API de Clima.

## Arquivos

- `swagger.json` - Especificação OpenAPI em formato JSON
- `swagger.yaml` - Especificação OpenAPI em formato YAML
- `docs.go` - Código Go gerado para integração com Swagger

## Acessando a Documentação

Quando a aplicação estiver em execução, você pode acessar a interface interativa Swagger UI em:

```
http://localhost:8080/swagger/index.html
```

## Regenerando a Documentação

Para regenerar a documentação Swagger após fazer alterações nas anotações da API:

```bash
make swagger
```

Ou manualmente:

```bash
swag init -g cmd/api/main.go -o docs
```

## Endpoints da API

### Clima

- `GET /api/v1/temperature/{cep}` - Obtém a temperatura pelo CEP (Código de Endereçamento Postal) brasileiro
  - Sucesso: 200 - Retorna a temperatura em Celsius, Fahrenheit e Kelvin
  - Erro: 404 - CEP não encontrado
  - Erro: 422 - Formato de CEP inválido

### Saúde

- `GET /health` - Endpoint de verificação de saúde
- `GET /readiness` - Endpoint de verificação de prontidão
