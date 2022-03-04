package server

import (
	"log"
	"webxam/config"
)

func Init() {
	c := config.GetConfig()
	r := NewRouter()

	log.Println("starting application ...")
	r.Run(c.GetString("server.port"))

}
