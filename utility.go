package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

func configDB() {

	fmt.Print("\033[1;" + colorM + "mdatabase host: \033[0;" + colorM + "m")
	fmt.Scanln(&config.DB.Host)

	fmt.Print("\033[1;" + colorM + "mdatabase port: \033[0;" + colorM + "m")
	fmt.Scanln(&config.DB.Port)

	fmt.Print("\033[1;" + colorM + "mdatabase name: \033[0;" + colorM + "m")
	fmt.Scanln(&config.DB.Name)

	fmt.Print("\033[1;" + colorM + "mdatabase user: \033[0;" + colorM + "m")
	fmt.Scanln(&config.DB.User)

	fmt.Print("\033[1;" + colorM + "mdatabase pass: \033[0;" + colorM + "m")
	fmt.Scanln(&config.DB.Pass)

	fmt.Print("\033[000m")

	configSave()

	dbOpen()
}

func configProxy() {

	fmt.Print("\033[1;" + colorM + "mproxy host: \033[0;" + colorM + "m")
	fmt.Scanln(&config.Proxy.Host)

	fmt.Print("\033[1;" + colorM + "mproxy port: \033[0;" + colorM + "m")
	fmt.Scanln(&config.Proxy.Port)

	fmt.Print("\033[000m")

	configSave()
}

func configSave() {

	if f, e := os.UserHomeDir(); e == nil {
		if b, e := yaml.Marshal(&config); e == nil {

			if e := ioutil.WriteFile(f+"/.britishhorseracing.yaml", b, 0755); e != nil {
				log.Fatal(e)
			}

		} else {
			log.Fatal(e)
		}
	}
}

func dbOpen() {

	if db, e := sql.Open("mysql", config.DB.User+":"+config.DB.Pass+"@tcp("+config.DB.Host+":"+config.DB.Port+")/"+config.DB.Name+"?net_write_timeout=8640000"); e == nil {

		if e := db.Ping(); e == nil {

			DB = db

		} else {
			fmt.Print("\033[1;" + colorE + "m")
			log.Println("\033[0;" + colorE + "m" + e.Error() + "\033[000m")

			configDB()
		}

	} else {
		log.Fatal(e)
	}
}

func dateParse(datetime string) time.Time {

	var r time.Time

	if t, e := time.Parse("2006-01-02", strings.TrimSpace(datetime)); e == nil {
		r = t.UTC()
	}

	return r
}

func datetimeParse(datetime string) time.Time {

	var r time.Time

	if t, e := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(datetime)); e == nil {
		r = t.UTC()
	}

	return r
}
