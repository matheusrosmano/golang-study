# Golang web

Esse projeto visa aprender golang para web

## Requisitos

* Golang 1.18
* Docker
* Docker-compose

## Desenvolvimento

1. Crie o arquivo `.env` baseado no `.env.modelo` e ajuste os valores das variaveis de ambiente
1. Inicie o banco de dados
    ```
    docker-compose up -d
    ```
1. Execute o go
    ```
    go run main.go
    ```
1. Pronto, acesse http://localhost:8000

> Para administrar seu banco de dados , acesse http://localhost:8080