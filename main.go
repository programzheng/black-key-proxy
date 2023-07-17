/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/programzheng/black-key-proxy/cmd"
	_ "github.com/programzheng/black-key-proxy/internal/model"
)

func main() {
	cmd.Execute()
}
