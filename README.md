# api-go-gin
Projeto onde crio uma API para fazer um crud de Livros


## ✔️ Tecnologias e Frameworks utilizadas

- GO
- Docker
- Postgres
- Gin
- Gorm
- GoDotEnv


## 📚 Detalhes do projeto

Projeto criado usando Gin para facilitar a necessidade de buscar informações na URL e utilizando o Gorm como ORM.

para executar o projeto:
- na raiz do projeto execute "docker-compose up -d" para criar o banco de dados.
- na raiz do projeto, acesse a pasta migration e execute "go run migrate.go" para criar a tabela no banco de dados.
- na raiz do projeto execute "go run main.go" para rodar a aplicação.
