package main

import (
	"flag"
	"fmt"
)

func main() {
	loginPtr := flag.String("login", "", "login azure devops")
	passwordPtr := flag.String("password", "", "PAT для azure devops")

	flag.Parse()

	fmt.Println("login:", *loginPtr)
	fmt.Println("password:", *passwordPtr)

	var ado = AzureDevOps{}
	ado.Init(*loginPtr, *passwordPtr)
	ado.GetWorkitems()
}
