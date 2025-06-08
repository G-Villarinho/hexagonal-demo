# Makefile for the Hexagonal Demo project

# Carrega as variáveis do arquivo .env para que possamos usá-las aqui
# O traço na frente ignora o erro caso o .env não exista
-include .env
export

# --- Variáveis ---
# Define a URL do banco de dados usando as variáveis do .env
DATABASE_URL := postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable
MIGRATIONS_PATH := internal/adapter/repository/postgres/migrations

# --- Comandos Docker ---

.PHONY: up
up: ## Sobe os contêineres Docker em modo detached
	@echo "Iniciando contêineres Docker..."
	@docker-compose up --build -d

.PHONY: down
down: ## Para e remove os contêineres Docker
	@echo "Parando contêineres Docker..."
	@docker-compose down

.PHONY: down-v
down-v: ## Para os contêineres e REMOVE os volumes (limpeza total)
	@echo "Parando contêineres e removendo volumes..."
	@docker-compose down -v

.PHONY: logs
logs: ## Mostra os logs dos contêineres em tempo real
	@docker-compose logs -f

.PHONY: dev
dev: ## Sobe os contêineres em modo de desenvolvimento com live-reload (Air)
	@echo "Iniciando ambiente de desenvolvimento com Air..."
	@docker-compose up --build

# --- Comandos de Migração ---

.PHONY: migrate-create
migrate-create: ## Cria um novo arquivo de migração. Ex: make migrate-create name=add_products_table
	@echo "Criando migração: $(name)"
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

.PHONY: migrate-up
migrate-up: ## Aplica todas as migrações pendentes no banco de dados
	@echo "Aplicando migrações (up)..."
	@migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_PATH) up

.PHONY: migrate-down
migrate-down: ## Reverte a última migração aplicada
	@echo "Revertendo migração (down)..."
	@migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_PATH) down

# --- Ajuda ---

.PHONY: help
help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
