/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/programzheng/black-key-proxy/cmd"
)

func main() {
	cmd.Execute()
}
