/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"clix/cmd"
	_ "clix/cmd/config"
  _ "clix/cmd/symlink"
)

func main() {
	cmd.Execute()
}
