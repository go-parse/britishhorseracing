package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Fixture struct {
	ID          int
	Date        time.Time
	Course      string
	CourseID    int
	Country     string
	IsAvailable bool
	Session     string
}

func getFixtures(link string) *[]Fixture {

	f := make([]Fixture, 0)

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

			f = append(f, Fixture{
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

func getFixturesByMonth(year, month int) *[]Fixture {

	return getFixtures("https://www.britishhorseracing.com/feeds/v3/fixtures?fields=courseId,fixtureDate,fixtureSession,abandonedReasonCode,&resultsAvailable=true&order=desc&year=" + strconv.Itoa(year) + "&month=" + strconv.Itoa(month) + "=5&per_page=1000")
}

func getFixturesNextDays(days int) *[]Fixture {

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
