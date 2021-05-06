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

	// race := getJSONRace(2021, 21301)

	// fmt.Println("ID:", race.ID)
	// fmt.Println("FixtureID:", race.FixtureID)
	// fmt.Println("Number:", race.Number)
	// fmt.Println("Division:", race.Division)
	// fmt.Println("Name:", race.Name)
	// fmt.Println("Age:", race.Age)
	// fmt.Println("Sex:", race.Sex)
	// fmt.Println("Prize:", race.Prize)
	// fmt.Println("Currency:", race.Currency)
	// fmt.Println("Band:", race.Band)
	// fmt.Println("Datatime:", race.Datatime)
	// fmt.Println("Distance:", race.Distance)
	// fmt.Println("Change:", race.Change)
	// fmt.Println("Type:", race.Type)
	// fmt.Println("Abandoned:", race.Abandoned)
	// fmt.Println("Black:", race.Black)
	// fmt.Println("Plus10:", race.Plus10)
	// fmt.Println("RacingUK:", race.RacingUK)
	// fmt.Println("Challenger:", race.Challenger)
	// fmt.Println("Rider:", race.Rider)
	// fmt.Println("Animal:", race.Animal)
	// fmt.Println("WinTime:", race.WinTime)
	// fmt.Println("Runners:", race.Runners)
	// fmt.Println("MaxRunners:", race.MaxRunners)
	// fmt.Println("MinimumWeight:", race.MinimumWeight)
	// fmt.Println("WeightsRaised:", race.WeightsRaised)


	// for _, d := range getJSONEntries(2021, 38656) {
	// 	fmt.Println("RaceID:",d.RaceID)
	// 	fmt.Println("Horse ID:",d.Horse.ID)
	// 	fmt.Println("Horse Name:",d.Horse.Name)
	// 	fmt.Println("Jockey ID:",d.Jockey.ID)
	// 	fmt.Println("Jockey Name:",d.Jockey.Name)
	// 	fmt.Println("Trainer ID:",d.Trainer.ID)
	// 	fmt.Println("Trainer Name:",d.Trainer.Name)
	// 	fmt.Println("Owner ID:",d.Owner.ID)
	// 	fmt.Println("Owner Name:",d.Owner.Name)
	// 	fmt.Println("Division:",d.Division)
	// 	fmt.Println("Age:",d.Age)
	// 	fmt.Println("Sex:",d.Sex)
	// 	fmt.Println("Number:",d.Number)
	// 	fmt.Println("Drawn:",d.Drawn)
	// 	fmt.Println("Rating:",d.Rating)
	// 	fmt.Println("Weight:",d.Weight)
	// 	fmt.Println("Penalty:",d.Penalty)
	// 	fmt.Println("Nonrunner Horse:",d.Nonrunner.Horse)
	// 	fmt.Println("Nonrunner Reason:",d.Nonrunner.Reason)
	// 	fmt.Println("Nonrunner Datatime:",d.Nonrunner.Datatime)
	// 	fmt.Println("Status:",d.Status)
	// 	fmt.Println("JockeyClaim:",d.JockeyClaim)
	// 	fmt.Println("HeadGear:",d.HeadGear)
	// 	fmt.Println("WindSurgeryFirstRun:",d.WindSurgeryFirstRun)
	// 	fmt.Println("SilkCode:",d.SilkCode)
	// 	fmt.Println("SilkDescription:",d.SilkDescription)
	// 	fmt.Println("_______________")
	// }

	// nonrunners := getJSONNonrunners(2021, 38656)

	// fmt.Println("Title:", nonrunners.Title)
	// fmt.Println("Datatime:", nonrunners.Datatime)

	// for _, d := range nonrunners.NR {
	// 	fmt.Println("Horse:", d.Horse)
	// 	fmt.Println("Reason:", d.Reason)
	// 	fmt.Println("Datatime:", d.Datatime)
	// 	fmt.Println("_______________")
	// }

	// from := time.Now().UTC().AddDate(-2, 0, 0)

	// to := from.AddDate(0, 0, 10)

	// fields := make([]string, 0)
	// fields = append(fields, "courseId")
	// fields = append(fields, "fixtureId")
	// fields = append(fields, "meetingId")
	// fields = append(fields, "fixtureDate")
	// fields = append(fields, "firstRaceTime")
	// fields = append(fields, "fixtureSession")
	// fields = append(fields, "bcsEvent")
	// fields = append(fields, "fixtureType")
	// fields = append(fields, "highlightTitle")
	// fields = append(fields, "firstRace")
	// fields = append(fields, "courseName")
	// fields = append(fields, "abandonedReasonCode")

	// for _, d := range getJSONOfficials(2021, 12763) {

	// 	fmt.Println("Category:", d.Category)

	// 	for _, o := range d.Officials {
	// 		fmt.Println(o)
	// 	}
		
	// 	fmt.Println("_______________")
	// }

	// getJSONRace(2021, 12763)

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

	
	// for _, d := range getJSONFixtures(fixturesForMonth){
	// 	fmt.Println("ID:", d.ID)
	// 	fmt.Println("MetingID:", d.MetingID)
	// 	fmt.Println("RacecourseID:", d.RacecourseID)
	// 	fmt.Println("Racecourse:", d.Racecourse)
	// 	fmt.Println("Date:", d.Date)
	// 	fmt.Println("Bcs:", d.Bcs)
	// 	fmt.Println("Abandoned:", d.Abandoned)
	// 	fmt.Println("Region:", d.Region)
	// 	fmt.Println("Type:", d.Type)
	// 	fmt.Println("RacecardAvailable:", d.RacecardAvailable)
	// 	fmt.Println("EntriesAvailable:", d.EntriesAvailable)
	// 	fmt.Println("BlackTypeRaces:", d.BlackTypeRaces)
	// 	fmt.Println("ResultsAvailable:", d.ResultsAvailable)
	// 	fmt.Println("_______________")
	// }

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
