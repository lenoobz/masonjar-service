/**
 * Create by Le Trong on 17/Feb/2019
 */

package main

import (
	"log"

	"github.com/letrong/masonjar-service/accountservice/cmd"
)

var appName = "accountService"

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
