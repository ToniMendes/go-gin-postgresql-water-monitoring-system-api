# 💧 Water Monitoring System API 🚀

![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=for-the-badge&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-008080?style=for-the-badge&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/Status-Em%20Desenvolvimento-yellow?style=for-the-badge)

Uma API robusta e escalável desenvolvida em **Go** para o monitoramento inteligente de consumo de água residencial. O sistema integra geolocalização via CEP, persistência em PostgreSQL e um worker de alta performance para simulação e registro de consumo.

---

## 📋 Índice
- [Funcionalidades](#-funcionalidades)
- [Tecnologias](#-tecnologias)
- [Arquitetura](#-arquitetura)
- [Configuração e Execução](#-configuração-e-execução)
- [Endpoints da API](#-endpoints-da-api)
- [Testes](#-testes)
- [Worker de Monitoramento](#-worker-de-monitoramento)

---

## ✨ Funcionalidades

- ✅ **Gestão de Residências:** Cadastro e atualização de proprietários.
- 📍 **Integração ViaCEP:** Validação e preenchimento automático de endereço através do CEP.
- 🤖 **Monitoramento Automatizado:** Worker rodando em background para processar medições.
- ⚡ **Alta Concorrência:** Processamento paralelo utilizando Goroutines e Channels.
- 🗄️ **Persistência Segura:** Camada de dados utilizando PostgreSQL com pool de conexões.

---

## 🛠️ Tecnologias

As seguintes ferramentas foram usadas na construção do projeto:

- **[Go](https://go.dev/):** Linguagem principal.
- **[Gin Gonic](https://gin-gonic.com/):** Framework web HTTP de alta performance.
- **[PostgreSQL](https://www.postgresql.org/):** Banco de dados relacional.
- **[pgx](https://github.com/jackc/pgx):** Driver PostgreSQL e toolkit para Go.
- **[ViaCEP API](https://viacep.com.br/):** Serviço de consulta de endereços.

---

## 🏗️ Arquitetura

O projeto segue princípios de **Clean Architecture** e **S.O.L.I.D.**, organizado da seguinte forma:

```text
internal/
├── domain/           # Entidades e interfaces de repositório
├── usecase/          # Regras de negócio (Create, Update, DTOs)
├── infra/            # Implementações de DB, Serviços Externos e Injeção de Dependência
├── web/              # Handlers HTTP (Gin)
└── worker/           # Lógica de processamento em background
```

---

## ⚙️ Como Executar

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/go-gin-postgresql-water-monitoring-system-api.git
   ```

2. **Configure as variáveis de ambiente:**
   Crie um arquivo `.env` ou configure sua string de conexão no `configs/env`.

3. **Execute a aplicação:**
   ```bash
   go run cmd/api/main.go
   ```

---

## 📡 Endpoints da API

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| `POST` | `/residences` | Cadastra uma nova residência e proprietário. |
| `PUT` | `/residences/:id` | Atualiza os dados do proprietário de uma residência. |

---

## 🧪 Testes

O projeto conta com testes de unidade para garantir a integridade das regras de negócio e dos contratos da API.

Para executar todos os testes, utilize o comando:

```bash
go test ./... -v
```

---

## 🤖 Worker de Monitoramento

O sistema possui um worker interno (`WaterMonitoring`) que executa ciclos de medição a cada **90 segundos**. 

- **Concorrência:** Utiliza um pool de **100 workers** (Goroutines).
- **Fluxo:** Busca todos os IDs no banco via Channels e distribui o processamento de "faturas" de consumo de forma assíncrona.

---

## 🛠️ Próximos Passos (WIP)
- [ ] Implementação de Logs estruturados.
- [x] Adição de testes unitários para Entidades e Handlers.
- [ ] Dockerização completa da aplicação.