package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"os"
	"strconv"
)

func getHelp() {
	fmt.Println(`Importer for IBPT
	Example:
		` + os.Args[0] + ` file1 ESTATE
	`)
	os.Exit(0)
}

func main() {

	defaultUF := "AC"
	fileName := flag.String("file", "fileName", "a csv to import")
	flag.Parse()

	db, err := sql.Open("sqlite3", "./ibpt.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ';'
	fmt.Println("Reading: " + *fileName)

	var first bool = true
	for {
		record, err := r.Read()

		if first {
			first = false
			continue
		}

		if err == io.EOF {
			break
		}

		ex, _ := strconv.Atoi(record[1])
		db.Exec(`INSERT INTO produto (
				Uf,
				codigo,
				Ex,
				Tipo,
				Descricao,
				Nacional,
				Importado,
				Estadual,
				Municipal,
				VigenciaInicio,
				VigenciaFim,
				Chave,
				Versao,
				Fonte
			) values (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
			)`,
			defaultUF,
			record[0],
			ex,
			record[2],
			record[3],
			record[4],
			record[5],
			record[6],
			record[7],
			record[8],
			record[9],
			record[10],
			record[11],
			record[12],
		)
	}
}
