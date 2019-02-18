/**
 * Create by Le Trong on 17/Feb/2019
 */

package main

import (
	"fmt"

	"github.com/letrong/masonjar-service/accountservice/service"
)

var appName = "accountService"

func main() {
	fmt.Printf("Starting %s\n", appName)
	service.StartWebServer("6767")
}
