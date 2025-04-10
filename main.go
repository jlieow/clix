/*
Copyright © 2025 NAME HERE mail@jeromelieow.com
*/
package main

import (
	"clix/cmd"
	_ "clix/cmd/config"
	_ "clix/cmd/gui"
	_ "clix/cmd/symlink"
)

func main() {
	cmd.Execute()
}
