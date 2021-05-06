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
	"net/url"
	"path"
	"strconv"
	"time"
)

func genURL() url.URL {

	u := url.URL{
		Scheme: "https",
		Host:   "www.britishhorseracing.com",
	}

	u.Path = path.Join("feeds", "v3")

	return u
}

func genURLRacecourses() url.URL {

	u := genURL()

	u.Path = path.Join("feeds/v1/racecourses")

	return u
}

func genURLFixture(year, fixtureID int) url.URL {

	u := genURL()

	u.Path = path.Join(u.Path, "fixtures")
	u.Path = path.Join(u.Path, strconv.Itoa(year))
	u.Path = path.Join(u.Path, strconv.Itoa(fixtureID))

	q := u.Query()

	u.RawQuery = q.Encode()

	return u
}

func genURLRaces(year, fixtureID int) url.URL {

	u := genURLFixture(year, fixtureID)

	u.Path = path.Join(u.Path, "races")

	q := u.Query()

	u.RawQuery = q.Encode()

	return u
}

func genURLGoing(year, fixtureID int) url.URL {

	u := genURLFixture(year, fixtureID)

	u.Path = path.Join(u.Path, "going")

	q := u.Query()

	u.RawQuery = q.Encode()

	return u
}

func genURLOfficials(year, fixtureID int) url.URL {

	u := genURLFixture(year, fixtureID)

	u.Path = path.Join(u.Path, "officials")

	q := u.Query()

	u.RawQuery = q.Encode()

	return u
}

func genURLRace(year, raceID int) url.URL {

	u := genURL()

	u.Path = path.Join(u.Path, "races")

	u.Path = path.Join(u.Path, strconv.Itoa(year))
	u.Path = path.Join(u.Path, strconv.Itoa(raceID))
	u.Path = path.Join(u.Path, "0")

	q := u.Query()

	u.RawQuery = q.Encode()

	return u
}

func genURLEntries(year, raceID int) url.URL {

	u := genURLRace(year, raceID)

	u.Path = path.Join(u.Path, "entries")

	q := u.Query()

	u.RawQuery = q.Encode()

	return u
}

func genURLNonrunners(year, raceID int) url.URL {

	u := genURLRace(year, raceID)

	u.Path = path.Join(u.Path, "nonrunners")

	q := u.Query()

	u.RawQuery = q.Encode()

	return u
}

func genURLFixturesFromTo(page, per int, from, to time.Time, isAvailable bool) url.URL {

	u := genURL()

	u.Path = path.Join(u.Path, "fixtures")

	q := u.Query()

	q.Set("page", strconv.Itoa(page))
	q.Set("per_page", strconv.Itoa(per))
	q.Set("resultsAvailable", strconv.FormatBool(isAvailable))
	q.Set("fromdate", from.Format("20060102"))
	q.Set("todate", to.Format("20060102"))

	u.RawQuery = q.Encode()

	return u
}

func genURLFixturesForMonth(page, per, year, month int, isAvailable bool) url.URL {

	u := genURL()

	u.Path = path.Join(u.Path, "fixtures")

	q := u.Query()

	q.Set("page", strconv.Itoa(page))
	q.Set("per_page", strconv.Itoa(per))
	q.Set("resultsAvailable", strconv.FormatBool(isAvailable))
	q.Set("year", strconv.Itoa(year))
	q.Set("month", strconv.Itoa(month))

	u.RawQuery = q.Encode()

	return u
}