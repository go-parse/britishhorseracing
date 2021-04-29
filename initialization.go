package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
	
	initializationDB()
}

func configProxy() {

 	fmt.Print("\033[1;" + colorM + "mproxy host: \033[0;" + colorM + "m")
	fmt.Scanln(&config.Proxy.Host)

	fmt.Print("\033[1;" + colorM + "mproxy port: \033[0;" + colorM + "m")
	fmt.Scanln(&config.Proxy.Port)

	fmt.Print("\033[000m")

	configSave()
}

func initializationDB() {

	if db, e := sql.Open("mysql", config.DB.User+":"+config.DB.Pass+"@tcp("+config.DB.Host+":"+config.DB.Port+")/"+config.DB.Name+"?net_write_timeout=8640000"); e == nil {
	
		if e := db.Ping(); e == nil {

			DB = db
			
		} else {
			fmt.Print("\033[1;" + colorE+"m")
			log.Println("\033[0;" + colorE+"m"+e.Error()+"\033[000m")

			configDB()
		}
	
	} else {
		log.Fatal(e)
	}
}

func initialization() {

	defaults := func ()  {

		if len(config.DB.Host) == 0 {
			config.DB.Host = "127.0.0.1"
		}

		if len(config.DB.Port) == 0 {
			config.DB.Port = "3306"
		}

		if len(config.DB.Name) == 0 {
			config.DB.Name = "britishhorseracingGG"
		}

		if len(config.DB.User) == 0 {
			config.DB.User = "britishhorseracing"
		}

		if len(config.DB.Pass) == 0 {
			config.DB.Pass = "password"
		}


		if len(config.Proxy.Host) == 0 {
			config.Proxy.Host = "host"
		}

		if len(config.Proxy.Port) == 0 {
			config.Proxy.Port = "port"
		}
	}

	if f, e := os.UserHomeDir(); e == nil {

		if f, err := ioutil.ReadFile(f+"/.britishhorseracing.yaml"); err == nil {

			if e := yaml.Unmarshal(f, &config); e != nil {
				log.Fatal(e)
			}
		}

	} else {
		log.Fatal(e)
	}

	defaults()

	if *flags.proxy {
		configProxy()
	}

	if *flags.db {
		configDB()
	} else {
		initializationDB()
	}
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