/**
 * Create by Le Trong on 17/Feb/2019
 */

package service

import (
	"log"
	"net/http"
)

// StartWebServer : Start web server
func StartWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)

	log.Println("Start HTTP service at port " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
