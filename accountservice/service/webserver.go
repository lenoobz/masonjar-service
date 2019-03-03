/**
 * Create by Le Trong on 17/Feb/2019
 */

package service

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/letrong/masonjar-service/accountservice/conf"
)

// StartWebServer : Start web server
func StartWebServer(config *conf.Config, logger *logrus.Entry) {
	r := NewRouter()
	http.Handle("/", r)

	port := fmt.Sprintf(":%d", config.Port)
	logger.Infof("Listen to port %s. Running ...", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		logger.Errorf("An error occured starting HTTP listener at port %d", config.Port)
		logger.Errorf("Error: " + err.Error())
	}
}
