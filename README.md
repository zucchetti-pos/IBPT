IBPT
====

```
git clone https://github.com/Compufour/ibpt.git
cd ibpt
curl https://glide.sh/get | sh
glide install
go build main -o ibpt
```

## Executar

```
    ./ibpt

```

## Create the database


```
sqlite3 ibpt.db 'CREATE TABLE `produto` ( `codigo` VARCHAR(10) not null, `Uf` VARCHAR(100) null, `Ex` int null, `Descricao` VARCHAR(100) null, `Nacional` real null, `Estadual` real null, `Importado` real null, `Municipal` real null, `Tipo` VARCHAR(100) null, `VigenciaInicio` VARCHAR(100) null, `VigenciaFim` VARCHAR(100) null, `Chave` VARCHAR(100) null, `Versao` VARCHAR(100) null, `Fonte` VARCHAR(100) null); '

```

Populate the databases
----------------------

```
go build importer
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
'http://iws.ibpt.org.br/api/deolhonoimposto/Produtos?token=vgN7SceD-SV2wYa4E7zDarTnofKAOXDH5EJpfJmtYrKWHnw35Jhy_XxC4qc5WlBm&cnpj=00445335000113&codigo=70101000&uf=sc&ex=0'

```


