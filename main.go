// Echo2 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// fmt.Println(os.Args[1:])
	for i, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println(i, os.Args[i])
	}
	// fmt.Println(s)
}
