package main

import (
	"gateway/auth"
	"gateway/util"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := util.LoadConfig("/home/asishst/Desktop/Work/social/services/project/config/env")
	if err != nil {
		log.Fatalln("Failed to read config file: ", err)
	}
	log.Println("Loaded config")

	r := gin.Default()

	auth.RegisterRoutes(r, &config)

	r.Run(config.Port)

}
