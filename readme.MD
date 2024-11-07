## Fazendo chamadas rest

Para fazer chamadas rest existe na pasta **api** exemplos de chamadas que utilizam o **plugin REST Client do VS Code**, ou então tendo o **curl** instalado em sua máquina basta digitar os seguintes comandos: 

Para criar uma nova order:
```
curl -X POST http://localhost:8000/order \
  -H "Content-Type: application/json" \
  -d '{
    "id": "nova2",
    "price": 100.5,
    "tax": 0.5
  }'
```

Para listar as orders existentes:
```
curl -X GET http://localhost:8000/order \
  -H "Content-Type: application/json"
```

## Fazendo chamadas graphQL

Abra o navegador no endereço (http://localhost:50051/query) 

Para criar uma nova order: 
```
mutation {
  createOrder(input: {
	id: "1999",
    Price: 120.00,
    Tax: 10.0
  }) {
    id,
    Tax,
  	Price,
    FinalPrice
  }
}
```


Para listar as ordens execute a seguinte query: 

```
query listOrder {
  listOrder {
    id,
    Tax,
    Price,
    FinalPrice
  }
}
``` 

## Fazendo chamadas gRPC

Com o client [evans](https://github.com/ktr0731/evans) instalado digite os seguintes comandos:

```
$ evans -r repl 
```

Após abrir o terminal do **evans** selecione o package **pb**

```
package pb
```

Selecione então o serviço **OrderService** 

```
service OrderService
```

Fazendo chamadas para criação de order ou listagem de orders

```
# Para criar uma nova order
call CreateOrder

# Para listar as orders existentes
call ListOrder
```

