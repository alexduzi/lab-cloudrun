# Lab Cloud Run - API de Temperatura por CEP

Sistema em Go que recebe um CEP brasileiro, identifica a localização e retorna o clima atual em diferentes escalas de temperatura (Celsius, Fahrenheit e Kelvin). O sistema está publicado no Google Cloud Run.

## 📋 Descrição do Desafio

Este projeto foi desenvolvido como parte do desafio Full Cycle para criar uma API que:
- Recebe um CEP válido de 8 dígitos
- Consulta a localização através da API ViaCEP
- Busca as informações climáticas através da API WeatherAPI
- Retorna as temperaturas convertidas em Celsius, Fahrenheit e Kelvin

## 🌐 API Pública

A API está disponível publicamente no Google Cloud Run:
```
🔗 URL: https://lab-cloudrun-729219189762.us-central1.run.app
```

### Exemplos de Uso
```bash
# Consultar temperatura por CEP
curl https://lab-cloudrun-729219189762.us-central1.run.app/01310100

# Health check
curl https://lab-cloudrun-729219189762.us-central1.run.app/health

# Documentação Swagger
https://lab-cloudrun-729219189762.us-central1.run.app/swagger/index.html
```

## 🚀 Funcionalidades

- ✅ Validação de CEP no formato brasileiro (8 dígitos)
- ✅ Consulta de localização via ViaCEP
- ✅ Consulta de temperatura via WeatherAPI
- ✅ Conversão automática de temperaturas (°C, °F, K)
- ✅ Documentação Swagger/OpenAPI
- ✅ Health checks e readiness probes
- ✅ Graceful shutdown
- ✅ Docker e Docker Compose
- ✅ Deploy no Google Cloud Run

## 📊 Respostas da API

### Sucesso (200 OK)
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

### CEP Inválido (422 Unprocessable Entity)
```json
{
  "message": "invalid zipcode"
}
```

### CEP Não Encontrado (404 Not Found)
```json
{
  "message": "can not find zipcode"
}
```

## 🔧 Tecnologias Utilizadas

- **Go 1.25.1** - Linguagem de programação
- **Gin** - Framework web
- **Viper** - Gerenciamento de configurações
- **Swagger** - Documentação da API
- **Docker** - Containerização
- **Google Cloud Run** - Hospedagem serverless

### APIs Externas

- **ViaCEP** - Consulta de endereços por CEP
- **WeatherAPI** - Consulta de informações climáticas

## 📦 Pré-requisitos

- Go 1.25.1 ou superior
- Docker e Docker Compose
- Conta no [WeatherAPI](https://www.weatherapi.com/) para obter API Key
- (Opcional) Google Cloud SDK para deploy

## ⚙️ Configuração

### 1. Configurar Variáveis de Ambiente
```bash
# Copiar arquivo de exemplo
cp .env.example .env

# Editar o arquivo .env e adicionar sua chave da WeatherAPI
# WEATHER_API_KEY=sua_chave_aqui
```

### 2. Obter API Key do WeatherAPI

1. Acesse [WeatherAPI](https://www.weatherapi.com/)
2. Crie uma conta gratuita
3. Copie sua API Key
4. Cole no arquivo `.env`

### Variáveis de Ambiente

| Variável | Descrição | Padrão | Obrigatória |
|----------|-----------|--------|-------------|
| `APP_PORT` | Porta da aplicação | `8080` | Não |
| `WEATHER_API_KEY` | Chave da API WeatherAPI | - | **Sim** |
| `GIN_MODE` | Modo do Gin (debug/release/test) | `debug` | Não |
| `VIA_CEP_BASE_URL` | URL base da API ViaCEP | `https://viacep.com.br/ws/{cep}/json/` | Não |
| `WEATHER_BASE_URL` | URL base da API Weather | `http://api.weatherapi.com/v1/current.json` | Não |

## 🚀 Como Executar

### Opção 1: Usando Make (Recomendado)
```bash
# Configuração inicial (criar .env e instalar dependências)
make setup

# Executar a aplicação localmente
make run

# Build da aplicação
make build

# Executar com Docker Compose
make docker-compose-up

# Parar Docker Compose
make docker-compose-down
```

### Opção 2: Usando Go Diretamente
```bash
# Baixar dependências
go mod download

# Executar
go run cmd/api/main.go
```

### Opção 3: Usando Docker
```bash
# Build da imagem
docker build -t lab-cloudrun-api .

# Executar container
docker run -p 8080:8080 --env-file .env lab-cloudrun-api
```

### Opção 4: Usando Docker Compose
```bash
# Iniciar aplicação
docker-compose up -d

# Ver logs
docker-compose logs -f

# Parar aplicação
docker-compose down
```

## 🧪 Testes
```bash
# Executar todos os testes
make test

# Executar apenas testes unitários
make test-unit

# Executar testes com cobertura
make test-coverage

# Gerar relatório HTML de cobertura
make test-coverage-html
```

### Status dos Testes
```
[TODO: Adicionar status dos testes após implementação completa]

- [ ] Testes unitários do conversor de temperatura ✅
- [ ] Testes unitários dos handlers
- [ ] Testes unitários dos clients (CEP e Weather)
- [ ] Testes de integração end-to-end
- [ ] Cobertura mínima: 80%
```

## 📝 Comandos do Makefile

### Configuração e Setup
```bash
make setup               # Configuração inicial do projeto
make deps                # Baixar e atualizar dependências
```

### Desenvolvimento Local
```bash
make run                 # Executar aplicação localmente
make build               # Compilar aplicação
make swagger             # Gerar documentação Swagger
```

### Testes e Qualidade
```bash
make test                # Executar todos os testes
make test-unit           # Executar apenas testes unitários
make test-integration    # Executar testes de integração
make test-coverage       # Executar testes com relatório de cobertura
make test-coverage-html  # Gerar relatório HTML de cobertura
make lint                # Executar análise de código (golangci-lint)
```

### Docker
```bash
make docker-build            # Build da imagem Docker
make docker-run              # Executar container Docker
make docker-stop             # Parar e remover container
make docker-logs             # Ver logs do container
make docker-clean            # Limpar recursos Docker
```

### Docker Compose
```bash
make docker-compose-up       # Iniciar com Docker Compose
make docker-compose-up-build # Build e iniciar com Docker Compose
make docker-compose-down     # Parar Docker Compose
make docker-compose-logs     # Ver logs do Docker Compose
make docker-compose-restart  # Reiniciar serviços
```

### Utilitários
```bash
make clean               # Limpar artefatos de build e cache de testes
make all                 # Executar setup, build, test e lint
make help                # Exibir ajuda com todos os comandos
```

## 🌐 Endpoints da API

### Weather

#### GET /{cep}
Retorna a temperatura atual para o CEP informado.

**Parâmetros:**
- `cep` (path) - CEP brasileiro com 8 dígitos (com ou sem hífen)

**Exemplos:**
```bash
curl http://localhost:8080/01310100
curl http://localhost:8080/01310-100
```

### Health Checks

#### GET /health
Verifica se o serviço está saudável.
```bash
curl http://localhost:8080/health
```

#### GET /readiness
Verifica se o serviço está pronto para receber tráfego.
```bash
curl http://localhost:8080/readiness
```

### Documentação

#### GET /swagger/index.html
Documentação interativa da API (Swagger UI).
```bash
# Acessar no navegador
http://localhost:8080/swagger/index.html
```

## 🏗️ Estrutura do Projeto
```
.
├── cmd/
│   └── api/
│       └── main.go                 # Ponto de entrada da aplicação
├── internal/
│   ├── client/
│   │   ├── cep.go                  # Cliente da API ViaCEP
│   │   └── weather.go              # Cliente da API WeatherAPI
│   ├── config/
│   │   └── config.go               # Gerenciamento de configurações
│   ├── conversor/
│   │   ├── temperature_conversor.go
│   │   └── temperature_conversor_test.go
│   ├── http/
│   │   ├── error/
│   │   │   └── http_errors.go      # Definição de erros HTTP
│   │   ├── middleware/
│   │   │   ├── error.go            # Middleware de tratamento de erros
│   │   │   └── error_test.go
│   │   ├── get_temperature.go      # Handler principal
│   │   ├── handler.go              # Setup do handler
│   │   ├── health.go               # Endpoints de health check
│   │   └── router.go               # Configuração de rotas
│   └── model/
│       └── model.go                # Estruturas de dados
├── docs/
│   ├── swagger.json                # Especificação OpenAPI (JSON)
│   ├── swagger.yaml                # Especificação OpenAPI (YAML)
│   └── docs.go                     # Código gerado do Swagger
├── test/
│   └── integration/                # [TODO] Testes de integração
├── .env.example                    # Exemplo de variáveis de ambiente
├── .dockerignore
├── Dockerfile                      # Configuração Docker (multi-stage)
├── docker-compose.yml              # Orquestração Docker
├── Makefile                        # Comandos de automação
├── go.mod                          # Dependências Go
└── README.md                       # Este arquivo
```

## 🐳 Deploy no Google Cloud Run

### Pré-requisitos

1. Instalar [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)
2. Fazer login:
```bash
gcloud auth login
```

3. Configurar projeto:
```bash
gcloud config set project [SEU-PROJECT-ID]
```

### Passos para Deploy

#### 1. Build e Push da Imagem
```bash
# Build usando Cloud Build
gcloud builds submit --tag gcr.io/[SEU-PROJECT-ID]/lab-cloudrun-api

# Ou build local e push
docker build -t gcr.io/[SEU-PROJECT-ID]/lab-cloudrun-api .
docker push gcr.io/[SEU-PROJECT-ID]/lab-cloudrun-api
```

#### 2. Deploy no Cloud Run
```bash
gcloud run deploy lab-cloudrun-api \
  --image gcr.io/[SEU-PROJECT-ID]/lab-cloudrun-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=[SUA-API-KEY] \
  --set-env-vars GIN_MODE=release \
  --port 8080 \
  --memory 256Mi \
  --cpu 1 \
  --max-instances 10 \
  --timeout 60
```

#### 3. Verificar Deploy
```bash
# Obter URL do serviço
gcloud run services describe lab-cloudrun-api \
  --platform managed \
  --region us-central1 \
  --format 'value(status.url)'

# Testar endpoint
curl [URL-DO-CLOUD-RUN]/health
```

### Configurações Adicionais (Opcional)

#### Configurar Domínio Customizado
```bash
gcloud run domain-mappings create \
  --service lab-cloudrun-api \
  --domain seu-dominio.com \
  --region us-central1
```

#### Configurar Secrets (para API Keys sensíveis)
```bash
# Criar secret
echo -n "sua-api-key" | gcloud secrets create weather-api-key --data-file=-

# Deploy com secret
gcloud run deploy lab-cloudrun-api \
  --image gcr.io/[SEU-PROJECT-ID]/lab-cloudrun-api \
  --update-secrets WEATHER_API_KEY=weather-api-key:latest
```

## 🔬 Fórmulas de Conversão

### Celsius para Fahrenheit
```
F = C × 1.8 + 32
```

### Celsius para Kelvin
```
K = C + 273.15
```

**Nota:** A aplicação utiliza `273.15` (valor cientificamente preciso) ao invés de `273` mencionado no desafio.

## 📚 Documentação da API

A documentação completa da API está disponível através do Swagger UI:

- **Local:** http://localhost:8080/swagger/index.html
- **Produção:** https://lab-cloudrun-729219189762.us-central1.run.app/swagger/index.html

### Regenerar Documentação
```bash
# Instalar swag (se necessário)
go install github.com/swaggo/swag/cmd/swag@latest

# Gerar documentação
make swagger
```

## 🐛 Troubleshooting

### Erro: "WEATHER_API_KEY is not set"

**Solução:** Configure a variável de ambiente no arquivo `.env`:
```bash
WEATHER_API_KEY=sua_chave_aqui
```

### Erro: "can not find zipcode" para CEP válido

**Possíveis causas:**
1. CEP não existe na base do ViaCEP
2. Problema de conectividade com a API ViaCEP
3. CEP muito recente (ainda não cadastrado)

### Container Docker não inicia

**Solução:** Verificar logs:
```bash
docker logs lab-cloudrun-api
# ou
make docker-logs
```

### Porta 8080 já em uso

**Solução:** Alterar porta no `.env`:
```bash
APP_PORT=8081
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto foi desenvolvido como parte do desafio Full Cycle.

## 👤 Autor

**Alex Duzi**
- Email: duzihd@gmail.com
- GitHub: [@alexduzi](https://github.com/alexduzi)

---

⭐ Se este projeto foi útil para você, considere dar uma estrela no repositório!
