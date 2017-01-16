all:
	go build importer.go
	go build main.go
clean:
	rm -rf *.db
	rm -rf data
	rm -rf *.csv
	rm -rf a.out
	rm -rf *.sql
build:
	curl https://glide.sh/get | sh
	glide install
	go build -o ibpt main.go
