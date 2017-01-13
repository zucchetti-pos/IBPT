package main

import (
    "io"
    "net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Produto struct {
    Codigo string
    Uf string
    Ex int
    Descricao string
    Nacional float64
    Estadual float64
    Importado float64
    Municipal float64
    Tipo string
    VigenciaInicio string
    VigenciaFim string
    Chave string
    Versao string
    Fonte string
}

func GetProduto(codigo string, uf string, ex int) Produto {
    var result = Produto{}
    db, err := sql.Open("sqlite3", "./artigo.db")
    defer db.Close()
    checkErr(err)
    rows, err := db.Query("SELECT * FROM produto where codigo = ? and uf = ? and ex = ?", codigo, uf, ex)
    checkErr(err)
    for rows.Next() {
        err = rows.Scan(&result.Codigo, &result.Uf, &result.Ex, &result.Descricao, &result.Nacional, &result.Estadual, &result.Importado, &result.Municipal, &result.Tipo, &result.VigenciaInicio, &result.VigenciaFim, &result.Chave, &result.Versao, &result.Fonte)
        checkErr(err)
    }

    rows.Close() //good habit to close

    return result
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    http.HandleFunc("/", HandleIndex)
    http.ListenAndServe(":8082", nil)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello World")
}

