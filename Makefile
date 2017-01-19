all: dependencies build

clean:
	rm -rf *.db
	rm -rf data
	rm -rf *.csv
	rm -rf a.out
	rm -rf *.sql
	rm -rf src/compufour/vendor
	rm -rf bin
	rm -rf pkg

dependencies:
	if [ ! -d "bin" ]; then mkdir bin; fi
	curl https://glide.sh/get | sh
	cd ./src/compufour && glide install
	cd ./src/compufour/cmd/importer && go install -v

build: 
	go build -o bin/ibpt src/compufour/cmd/service/main.go
	go build -o bin/importer src/compufour/cmd/importer/main.go

database:
	sqlite3 ibpt.db 'CREATE TABLE `produto` ( `codigo` VARCHAR(10) not null, `Uf` VARCHAR(100) null, `Ex` int null, `Descricao` VARCHAR(100) null, `Nacional` real null, `Estadual` real null, `Importado` real null, `Municipal` real null, `Tipo` VARCHAR(100) null, `VigenciaInicio` VARCHAR(100) null, `VigenciaFim` VARCHAR(100) null, `Chave` VARCHAR(100) null, `Versao` VARCHAR(100) null, `Fonte` VARCHAR(100) null); '

