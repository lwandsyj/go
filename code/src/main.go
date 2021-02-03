package main

import (
	"fmt"
	"os"
)

func main(){
	a:=os.Args
	fmt.Println(a)
	args:=a[1:]
	fmt.Println(args)
}