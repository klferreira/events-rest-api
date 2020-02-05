# REST API de Eventos
Implementação do [desafio de back-end](https://github.com/ingresse/backend-test/blob/master/readme.md) da [Ingresse](https://www.ingresse.com/).

O teste consiste em escrever uma API REST que disponibiliza endpoints de buscar, criar, atualizar e deletar eventos em um banco de dados.

# Executando

## Makefile

Utilize o seguinte comando do Makefile contido na raiz do projeto para construir as imagens dos serviços
```
    make build
```

E então, para subir os containers:
```
    make up
```

Após isso a API estará rodando na porta `:3000`.
Os endpoints para os eventos estão disponívels como (`GET`, `POST`, `PUT`, `DELETE`) em `/v1/events`.
