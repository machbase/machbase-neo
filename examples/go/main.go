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
		"grpc_exec_cretable":  grpc_exec_cretable,
		"grpc_exec_droptable": grpc_exec_droptable,
		"grpc_exec_insert":    grpc_exec_insert,
		"grpc_queryrow":       grpc_queryrow,
		"grpc_query":          grpc_query,
		"grpc_append":         grpc_append,
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
