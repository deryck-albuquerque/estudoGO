# estudoGO

Projeto de estudo da linguagem Go com foco em desenvolvimento backend e arquitetura em camadas.

A aplicação consiste em uma API HTTP simples para criação e listagem de usuários, utilizando persistência em memória (sem banco de dados) e um client separado que consome os endpoints da API.

---

# Objetivos do Projeto

Esse projeto foi desenvolvido com o objetivo de praticar conceitos fundamentais da linguagem Go, incluindo:

- Estruturação de projetos backend
- Arquitetura em camadas
- Manipulação de JSON
- Criação de APIs HTTP
- Organização de responsabilidades
- Uso de interfaces
- Injeção de dependência
- Tratamento de erros
- Uso de UUIDs
- Comunicação cliente-servidor
- Logging estruturado com slog

---

# Estrutura do Projeto

```bash
estudoGO/
│
├── cmd/
│   ├── api/
│   │   └── main.go
│   │
│   └── client/
│       └── main.go
│
├── internal/
│   ├── handlers/
│   ├── models/
│   ├── repositories/
│   │   └── users/
│   │       └── users.go
│   │
│   └── usecases/
│       └── usecases.go
│
├── go.mod
├── go.sum
└── README.md
```

---

# Arquitetura

O projeto foi dividido em camadas para separar responsabilidades e manter a aplicação organizada e desacoplada.

## Handlers

Responsáveis por:

- Receber requisições HTTP
- Validar payloads
- Serializar/deserializar JSON
- Retornar respostas HTTP

Camada de entrada da aplicação.

---

## UseCases

Camada de regra de negócio.

Responsável por:

- Validar regras da aplicação
- Orquestrar o fluxo de execução
- Controlar criação e listagem de usuários
- Garantir que e-mails duplicados não sejam cadastrados

---

## Repositories

Responsável pela persistência dos dados.

Nesse projeto os usuários são armazenados em memória utilizando slices.

Essa camada abstrai a origem dos dados e facilita futuras trocas para banco de dados real como PostgreSQL ou MongoDB.

---

## Models

Contém:

- Entidades da aplicação
- Estruturas de request
- Estruturas de response
- Modelos de erro

---

# Funcionalidades

## Criar usuário

Endpoint:

```http
POST /users
```

Payload:

```json
{
  "name": "Deryck",
  "email": "deryck.henrique22@gmail.com"
}
```

Resposta:

```json
{
  "newUserID": "uuid"
}
```

---

## Listar usuários

Endpoint:

```http
GET /users
```

Resposta:

```json
[
  {
    "ID": "uuid",
    "Name": "Deryck",
    "Email": "deryck.henrique22@gmail.com"
  }
]
```

---

# Como executar

## Clonar projeto

```bash
git clone https://github.com/deryck-albuquerque/estudoGO.git
```

---

## Entrar na pasta

```bash
cd estudoGO
```

---

## Instalar dependências

```bash
go mod tidy
```

---

## Executar API

```bash
go run ./cmd/api
```

Servidor iniciado em:

```bash
http://localhost:8080
```

---

## Executar Client

Em outro terminal:

```bash
go run ./cmd/client
```

---

# Fluxo da Aplicação

```text
Client
   ↓
Handlers
   ↓
UseCases
   ↓
Repositories
   ↓
Memória
```

---

# Tecnologias Utilizadas

- Go
- net/http
- encoding/json
- slog
- UUID

---

# Conceitos Praticados

- HTTP Server
- REST API
- JSON Encoder/Decoder
- Injeção de Dependência
- Repository Pattern
- Clean Architecture (simplificada)
- Organização modular
- Tratamento de erros
- Ponteiros em Go
- Interfaces
- Logging estruturado

---

# Aprendizados

Esse projeto foi desenvolvido para aprofundar conhecimentos em Go backend, explorando desde a criação de APIs HTTP até organização arquitetural baseada em responsabilidades.

Também serviu como prática para entender melhor:

- slices
- ponteiros
- interfaces
- métodos
- structs
- fluxo de aplicações backend em Go
- separação de responsabilidades
- comunicação HTTP cliente-servidor

---

# Autor

Desenvolvido por Deryck Henrique Albuquerque