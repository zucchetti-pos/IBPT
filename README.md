Tabela IBPT
===========

Códigos referentes ao artigo sobre Go para a Revista iMasters

Compile and test
----------------

Para configurar o GOPATH, entrar no diretório raiz do projeto e executar:

```
    export GOPATH="$GOPATH:$(pwd)"

```

Para instalar o Glide

```
    curl https://glide.sh/get | sh

```

Instalar as dependências

```
    glide install
    go install -v

```

Para compilar

```
    go build -o artigo-imasters

```

Para rodar os testes

```
    go test

```

## Executar

```
    ./artigo-imasters

```

## Gerar o container

No diretório raiz executar

```
    docker build . -t eminetto/artigo-imasters

```

## Executar dentro do container

```
    docker run --rm -p 8082:8082 eminetto/artigo-imasters

```

## Enviar para o Docker Hub

```
    docker login
    docker push eminetto/artigo-imasters

```


## Executar em "produção"

```
    docker pull eminetto/artigo-imasters
    docker run --rm -p 8082:8082 eminetto/artigo-imasters

```
