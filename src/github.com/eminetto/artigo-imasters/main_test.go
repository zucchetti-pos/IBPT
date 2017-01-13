package main

import "testing"

func TestGetProduto(t *testing.T) {
    var expected = Produto{"70101000","SC",0,"Ampolas de vidro,p/transporte ou embalagem",13.45,17.00,18.02,0.00,"0","01/01/2017","30/06/2017","W7m9E1","17.1.A","IBPT"}
    var result = GetProduto("70101000", "SC", 0)
    if result !=  expected {
        t.Error(
                "expected", expected,
                "got", result,
            )
    }
}
