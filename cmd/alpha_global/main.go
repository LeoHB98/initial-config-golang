package main

import (
	"log"

	"gitalpha.com/go/nome_projeto/config"
	"gitalpha.com/go/nome_projeto/utils"
)

func main() {

	cfg, err := config.New(true)
	if err != nil {
		panic(err)
	}

	utils := utils.Start(cfg)
	log.Println(utils)

}
