## Выбор используемой памяти

### Из терминала

`in-memory`:

Добавить параметр `USE_MEMORY=true`

`postgres`:

Добавить параметр `USE_MEMORY=false`

### Через `docker-compose`

Выставить нужное значение переменной `USE_MEMORY` в файле `.env `

## REST

### POST

`http://localhost:8080/api/generate`

Пример запроса:
```
POST http://localhost:8080/api/generate
Content-Type: application/json

{
    "link": "test1"
}
```

Пример ответа:
```
{
    "token": "Qh0WldYzMN"
}
```

### GET

`http://localhost:8080/:token`

Пример запроса:
```
GET http://localhost:8080/Qh0WldYzMN
```

Пример ответа:
```
{
    "link": "test1"
}
```


## gRPC

`localhost:9090`

### GetToken
 ```
GRPC localhost:9090/api.URLShortener/GetToken

{
"link" : "grpc test1"
}
```
Ответ:
 ```
{
  "token": "geVSRubR3C"
}
```
### GetLinkByToken

 ```
GRPC localhost:9090/api.URLShortener/GetLinkByToken

{
  "token" : "geVSRubR3C"
}
```
Ответ:
 ```
{
  "link": "grpc test1"
}
```