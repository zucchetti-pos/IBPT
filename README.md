IBPT
====

```
git clone https://github.com/Compufour/ibpt.git
cd ibpt
export GOPATH=`pwd`
make build
```

## Executar

```
    ./bin/ibpt

```

## Create the database


```
make database

```

Populate the databases
----------------------

```
./import-tables ~/Downloads

```

Usage example
--------------

```
curl http://127.0.0.1:8082 -d "uf=GO&ex=0&codigo=2801" 

```

Alternatively one can use the real api
--------------------------------------

```
curl -X GET --header 'Accept: application/json'
'http://iws.ibpt.org.br/api/deolhonoimposto/Produtos?token=$DE_OLHO_NO_IMPOSTO_CNPJ_TOKEN&cnpj=00445335000113&codigo=70101000&uf=sc&ex=0'

```


