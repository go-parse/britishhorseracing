package main

/*
 * Copyright 2021 Vasiliy Vdovin
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

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

	Proxy = configuration.Proxy

	if db, e := sql.Open("mysql", configuration.DBUser+":"+configuration.DBPassword+"@tcp("+configuration.DBHost+":"+configuration.DBPort+")/"+configuration.DBName+"?net_write_timeout=8640000"); e == nil {
		DB = db
	} else {
		log.Fatal(e)
	}

	from := time.Now().UTC().AddDate(-2, 0, 0)

	to := from.AddDate(0, 0, 10)

	fields := make([]string, 0)
	fields = append(fields, "courseId")
	fields = append(fields, "fixtureId")
	fields = append(fields, "meetingId")
	fields = append(fields, "fixtureDate")
	fields = append(fields, "firstRaceTime")
	fields = append(fields, "fixtureName")
	fields = append(fields, "fixtureSession")
	fields = append(fields, "bcsEvent")
	fields = append(fields, "fixtureType")
	fields = append(fields, "highlightTitle")
	fields = append(fields, "firstRace")
	fields = append(fields, "majorEvent")
	fields = append(fields, "distance")
	fields = append(fields, "courseName")
	fields = append(fields, "fixtureYear")
	fields = append(fields, "abandonedReasonCode")

	fixtures := genURLFixture(2021, 478)
	races := genURLRaces(2021, 478)
	going := genURLGoing(2021, 478)
	race := genURLRace(2021, 45563)
	entries := genURLEntries(2021, 45563)
	nonrunners := genURLNonrunners(2021, 45563)
	fixturesFromTo := genURLFixturesFromTo(1, 3, from, to, true, fields)
	fixturesForMonth := genURLFixturesForMonth(1, 3, 2021, 5, true, fields)

	fmt.Println(fixtures.String())
	fmt.Println(races.String())
	fmt.Println(going.String())
	fmt.Println(race.String())
	fmt.Println(entries.String())
	fmt.Println(nonrunners.String())
	fmt.Println(fixturesFromTo.String())
	fmt.Println(fixturesForMonth.String())

}
