/*
Copyright © 2025 Abinand P
*/
package main

import (
	"log"

	"github.com/Abiji-2020/codetool/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
