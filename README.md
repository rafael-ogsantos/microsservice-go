# microsservice-go

   O foco desse repositório é gerar uma comunicação entre microserviços independemente se eles são ou não robustos. 
A princípio, os serviços foram desenvolvidos em `go`. Mas a idéia é expandir isso utilizando outras linguagens para serviços posteriores e que façam sentido para o projeto.

O fluxo se resume da seguinte maneira:
  - `broker` recebe as requisições e gera as mensagens para serem consumidas em mensageria pelos outros apps.
  - `listener` ouvinte das mensagens que sao postadas e as consume de acordo com oque cada app espera. 
  - `auth` a principio é o unico servico desse projeto.
    

Para filas é utilizado o serviço `RabbitMQ` e esse serviço foi implementado no código utilizando um lib do próprio go.

## Clone
    git clone https://github.com/smarters/microsservice-go.git

## Run the apps
    go run .
    
Por enquanto os serviços precisam ser executados usando esse comando. Mas posteriormente uma imagem de cada serviço será gerada e assim
poderemos fazer o build completo com `docker-compose up -d`.

