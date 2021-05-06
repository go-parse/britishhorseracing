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
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var config = struct {
	DB struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"mdatabase"`

	Proxy struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"proxy"`
}{}

var flags = struct {
	db    *bool
	proxy *bool
} {}

var DB *sql.DB
var Proxy = ""

const colorM = "32" // Message
const colorE = "31" // Error

func main() {

	listener()
	initialization()

	// for _, d := range  getJSONRacecourses() {
	// 	fmt.Println("ID",  d.ID)
	// 	fmt.Println("Name",  d.Name)
	// 	fmt.Println("Type",  d.Type)
	// 	fmt.Println("Handedness",  d.Handedness)
	// 	fmt.Println("Region",  d.Region)
	// 	fmt.Println("Post",  d.Post)
	// 	fmt.Println("Coordinate",  d.Coordinate)
	// 	fmt.Println("FirstRace",  d.FirstRace)
	// 	fmt.Println("NextFixture",  d.NextFixture)
	// 	fmt.Println("_______________")
	// }

	// fixture := getJSONFixture(2021, 12763)

	// fmt.Println("ID:", fixture.ID)
	// fmt.Println("Date:", fixture.Date)
	// fmt.Println("MetingID:", fixture.MetingID)
	// fmt.Println("RacecourseID:", fixture.RacecourseID)
	// fmt.Println("Racecourse:", fixture.Racecourse)
	// fmt.Println("Abandoned:", fixture.Abandoned)
	// fmt.Println("Type:", fixture.Type)
	// fmt.Println("Type:", fixture.Type)
	// fmt.Println("Session:", fixture.Session)
	// fmt.Println("Surface:", fixture.Surface)
	// fmt.Println("Planning:", fixture.Planning)
	// fmt.Println("Weather:", fixture.Weather)
	// fmt.Println("Stalls:", fixture.Stalls)
	// fmt.Println("Going:", fixture.Going)
	// fmt.Println("Inspection:", fixture.Inspection)
	// fmt.Println("Rail:", fixture.Rail)
	// fmt.Println("Watering:", fixture.Watering)
	// fmt.Println("Other:", fixture.Other)
	// fmt.Println("Updated:", fixture.Updated)


	// for _, d := range  getJSONRaces(2021, 12763) {
	// 	fmt.Println("ID:", d.ID)
	// 	fmt.Println("Year:", d.Year)
	// 	fmt.Println("Division:", d.Division)
	// 	fmt.Println("Datatime:", d.Datatime)
	// 	fmt.Println("Name:", d.Name)
	// 	fmt.Println("Age:", d.Age)
	// 	fmt.Println("Prize:", d.Prize)
	// 	fmt.Println("Currency:", d.Currency)
	// 	fmt.Println("Class:", d.Class)
	// 	fmt.Println("Band:", d.Band)
	// 	fmt.Println("Distance:", d.Distance)
	// 	fmt.Println("Change:", d.Change)
	// 	fmt.Println("Type:", d.Type)
	// 	fmt.Println("Abandoned:", d.Abandoned)
	// 	fmt.Println("Black:", d.Black)
	// 	fmt.Println("Plus10:", d.Plus10)
	// 	fmt.Println("_______________")
	// }
	// https://www.britishhorseracing.com/feeds/v3/fixtures/2021/715/going

	// going := getJSONGoing(2021, 715)

	// fmt.Println("CourseID:", going.CourseID)
	// fmt.Println("CourseID:", going.CourseID)
	// fmt.Println("Datatime:", going.Datatime)
	// fmt.Println("Type:", going.Type)
	// fmt.Println("Code:", going.Code)
	// fmt.Println("Ground:", going.Ground)
	// fmt.Println("Stick:", going.Stick)
	// fmt.Println("Rails:", going.Rails)
	// fmt.Println("Stalls:", going.Stalls)
	// fmt.Println("Weather:", going.Weather)
	// fmt.Println("Watering:", going.Watering)
	// fmt.Println("WateringStatus:", going.WateringStatus)

	// from := time.Now().UTC().AddDate(-2, 0, 0)

	// to := from.AddDate(0, 0, 10)

	// fields := make([]string, 0)
	// fields = append(fields, "courseId")
	// fields = append(fields, "fixtureId")
	// fields = append(fields, "meetingId")
	// fields = append(fields, "fixtureDate")
	// fields = append(fields, "firstRaceTime")
	// fields = append(fields, "fixtureName")
	// fields = append(fields, "fixtureSession")
	// fields = append(fields, "bcsEvent")
	// fields = append(fields, "fixtureType")
	// fields = append(fields, "highlightTitle")
	// fields = append(fields, "firstRace")
	// fields = append(fields, "majorEvent")
	// fields = append(fields, "distance")
	// fields = append(fields, "courseName")
	// fields = append(fields, "fixtureYear")
	// fields = append(fields, "abandonedReasonCode")

	for _, d := range getJSONOfficials(2021, 12763) {

		fmt.Println("Category:", d.Category)

		for _, o := range d.Officials {
			fmt.Println(o)
		}
		
		fmt.Println("_______________")
	}

	// racecourses := genURLRacecourses()
	// fixtures := genURLFixture(2021, 478)
	// races := genURLRaces(2021, 478)
	// going := genURLGoing(2021, 478)
	// officials := genURLOfficials(2021, 478)
	// race := genURLRace(2021, 45563)
	// entries := genURLEntries(2021, 45563)
	// nonrunners := genURLNonrunners(2021, 45563)
	// fixturesFromTo := genURLFixturesFromTo(1, 3, from, to, true, fields)
	// fixturesForMonth := genURLFixturesForMonth(1, 3, 2021, 5, true, fields)

	// fmt.Println(racecourses.String())
	// fmt.Println(fixtures.String())
	// fmt.Println(races.String())
	// fmt.Println(going.String())
	// fmt.Println(officials.String())
	// fmt.Println(race.String())
	// fmt.Println(entries.String())
	// fmt.Println(nonrunners.String())
	// fmt.Println(fixturesFromTo.String())
	// fmt.Println(fixturesForMonth.String())

}

func dataParse(datatime string) time.Time {

	var r time.Time

	if t, e := time.Parse("2006-01-02", strings.TrimSpace(datatime)); e == nil {
		r = t.UTC()
	}
	
	return r
}

func datatimeParse(datatime string) time.Time {

	var r time.Time

	if t, e := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(datatime)); e == nil {
		r = t.UTC()
	}
	
	return r
}
