# REST API de Eventos
Implementação do [desafio de back-end](https://github.com/ingresse/backend-test/blob/master/readme.md) da [Ingresse](https://www.ingresse.com/).

O teste consiste em escrever uma API REST que disponibiliza endpoints de buscar, criar, atualizar e deletar eventos em um banco de dados.

## Executando

### Makefile

Utilize o seguinte comando do Makefile contido na raiz do projeto para construir as imagens dos serviços
```
    make build
```

E então, para subir os containers:
```
    make up
```

Após isso a API estará rodando na porta `:3000`.

## Testando

Para rodar os testes de funcionalidade execute
```
    make test
```

## Endpoints

Os endpoints disponibilizados por esta API são os seguintes:

- `GET -> /v1/events` (Busca por todos os eventos)
    - Obs.: Este endpoint pode receber um ou mais dos seguintes parâmetros de URL: [`name`, `place`, `tags`, `interested`, `sessions`]. (As buscas serão sempre por valores identicos aos recebidos na URL).

- `POST -> /v1/events` (Cria um novo evento)

- `PUT -> /v1/events` (Atualiza um evento existente)

- `PATCH -> /v1/events/{id}/add-interest` (Incrementa o número de "interessados" em um evento)

- `DELETE -> /v1/events/{id}` (Remove um evento do banco de dados)
