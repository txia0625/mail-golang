package main
import (
	"log"
	"net/http"
	"fmt"
)
const webPort = "80"

type Config struct{}

func main(){
	app := Config{}
	log.Printf("Starting Broker Service on Port %s", webPort)
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
	
}