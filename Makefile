clean:
	rm -rf *.db
	rm -rf data
	rm -rf *.csv
	rm -rf a.out
	rm -rf *.sql
	rm -rf src/compufour/vendor
	rm -rf bin
	rm -rf pkg
build:
	if [ ! -d "bin" ]; then mkdir bin; fi
	curl https://glide.sh/get | sh
	cd ./src/compufour && glide install
	cd ./src/compufour/cmd/importer && go install -v
	go build -o bin/ibpt src/compufour/cmd/service/main.go
	go build -o bin/importer src/compufour/cmd/importer/main.go
