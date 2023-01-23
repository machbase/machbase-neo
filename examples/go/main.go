package main

import (
	"fmt"
	"os"
)

/*
ex)
go run *.go grpc_exec_cretable
go run *.go grpc_exec_droptable
*/
func main() {
	ex := ""
	if len(os.Args) > 1 {
		ex = os.Args[1]
	}
	examples := map[string]func(){
		"grpc_cretable":  grpc_cretable,
		"grpc_droptable": grpc_droptable,
		"grpc_insert":    grpc_insert,
		"grpc_queryrow":  grpc_queryrow,
		"grpc_query":     grpc_query,
		"grpc_append":    grpc_append,
		"grpc_wave":      grpc_wave,
	}

	if fn, ok := examples[ex]; ok {
		fn()
	} else {
		fmt.Println("Usage")
		for k := range examples {
			fmt.Println("  go run *.go", k)
		}
	}
}
