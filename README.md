# microsservice-go

    O foco desse repositorio é gerar uma comunicacao entre microsservicos independemente se eles sao ou nao robustos. 
A principio, os servicos foram desenvolvidos em `go`. Mas minha ideia é expandir isso utilizando outras linguagens para servicos posteriores e que facam sentido para o projeto.

O fluxo se resume da seguinte maneira:
    - `broker` recebe as requisicoes e gera as mensagens para ser consumidas em mensageria pelos outros apps.
    - `listener` Ouvinte das mensagens que sao postadas e as consume de acordo com oque cada app espera. +
    

Para filas é utlizando o servico `RabbitMQ` e esse servico foi implementado no codigo utilizando um lib do proprio go.

## Clone
    git clone https://github.com/smarters/microsservice-go.git

## Run the apps
    go run .
    
Por enquanto os servicos precisam ser executados usando esse comando. Mas posteriormente uma imagem de cada servico sera gerada e assim
poderemos fazer o build completo com `docker-compose up -d`.

