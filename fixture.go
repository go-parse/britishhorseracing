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
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type fixture struct {
	ID          int
	Date        time.Time
	Course      string
	CourseID    int
	Country     string
	IsAvailable bool
	Session     string
}

func getFixtures(link string) *[]fixture {

	f := make([]fixture, 0)

	d := struct {
		Data []struct {
			ID          int    `json:"fixtureId"`
			Year        int    `json:"fixtureYear"`
			Date        string `json:"fixtureDate"`
			Course      string `json:"courseName"`
			CourseID    int    `json:"courseId"`
			Country     string `json:"region"`
			IsAvailable bool   `json:"resultsAvailable"`
			Session     string `json:"fixtureSession"`
		} `json:"data"`
	}{}

	get(link, &d)

	for _, d := range d.Data {

		if t, e := time.Parse("2006-01-02", strings.TrimSpace(d.Date)); e == nil {

			f = append(f, fixture{
				ID:          d.ID,
				Date:        t.UTC(),
				Course:      d.Course,
				CourseID:    d.CourseID,
				Country:     d.Country,
				IsAvailable: d.IsAvailable,
				Session:     d.Session,
			})

		} else {
			log.Fatal(e)
		}
	}

	return &f
}

func getFixturesByMonth(year, month int) *[]fixture {

	return getFixtures("https://www.britishhorseracing.com/feeds/v3/fixtures?fields=courseId,fixtureDate,fixtureSession,abandonedReasonCode,&resultsAvailable=true&order=desc&year=" + strconv.Itoa(year) + "&month=" + strconv.Itoa(month) + "=5&per_page=1000")
}

func getFixturesNextDays(days int) *[]fixture {

	from := time.Now().UTC()

	to := from.AddDate(0, 0, days)

	dfb := "https://www.britishhorseracing.com/feeds/v3/fixtures?fields=abandonedReasonCode,courseId,courseName,fixtureYear,fixtureId,fixtureDate,distance,firstRace,firstRaceTime,fixtureName,fixtureSession,fixtureType,highlightTitle,majorEvent,meetingId,&resultsAvailable=false,bcsEvent&fromdate=" + from.Format("20060102") + "&per_page=1000&todate=" + to.Format("20060102")

	return getFixtures(dfb)
}

func saveFixtures() {

	for _, d := range *getFixturesByMonth(2009, 1) {

		fmt.Println(d)

	}

}
