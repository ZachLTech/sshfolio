package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	ZachLTechMD, ZachLTechMDerr := os.ReadFile("./ZachLTech.md")
	check(ZachLTechMDerr)

	fmt.Println(string(ZachLTechMD))
}
