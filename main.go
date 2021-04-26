package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type configuration struct {
	DBUser     string `json:"DB_USER"`
	DBName     string `json:"DB_NAME"`
	DBPassword string `json:"DB_PASSWORD"`
	DBHost     string `json:"DB_HOST"`
	DBPort     string `json:"DB_PORT"`
	Proxy      string `json:"PROXY"`
}

var DB *sql.DB
var Proxy = ""

func main() {

	configuration := configuration{}

	if home, e := os.UserHomeDir(); e == nil {

		if f, e := os.Open(home + "/.britishhorseracing.json"); e == nil {

			if b, e := ioutil.ReadAll(f); e == nil {

				json.Unmarshal(b, &configuration)

			} else {
				log.Fatal(e)
			}

		} else if f, e := json.MarshalIndent(configuration, "", " "); e == nil {
			if e := ioutil.WriteFile(home+"/.britishhorseracing.json", f, 0644); e != nil {
				log.Fatal(e)
			}
		}

	} else {
		log.Fatal(e)
	}

	fmt.Println(configuration)

	Proxy = configuration.Proxy

	if db, e := sql.Open("mysql", configuration.DBUser+":"+configuration.DBPassword+"@tcp("+configuration.DBHost+":"+configuration.DBPort+")/"+configuration.DBName+"?net_write_timeout=8640000"); e == nil {
		DB = db
	} else {
		log.Fatal(e)
	}

}
