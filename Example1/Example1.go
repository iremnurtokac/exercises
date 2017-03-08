package main

import (
	"fmt"

	"github.bus.zalan.do/laas/hecate-report/config"
	_ "github.com/lib/pq"
)

type Attachment struct {
	Filename string
	Data     []byte
}

var conf *config.Config

func main() {
	conf = config.Get("config/config.json")
	//assingning variables as strings
	var user string
	var emailAddr string
	var usedEnvInput string
	var usedEnv config.Environment

	fmt.Print("Plz enter db-username (ldap): ")
	fmt.Scan(&user) // point to user

}
