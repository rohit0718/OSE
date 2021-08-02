package main

import (
	"log"
	"net/http"

	"github.com/comail/colog"
	"github.com/julienschmidt/httprouter"
)

func main() {
	colog.SetDefaultLevel(colog.LInfo)
	colog.SetMinLevel(colog.LInfo)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
	colog.SetMinLevel(colog.LTrace)

	router := httprouter.New()
	router.POST("/filter", Filter)
	router.POST("/prioritize", Prioritize)

	log.Print("info: server starting on the port :8888")
	if err := http.ListenAndServe(":8888", router); err != nil {
		log.Fatal(err)
	}
}
