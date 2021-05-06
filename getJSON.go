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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

var client = http.Client{Transport: transport, Timeout: time.Minute * 5}

func getJSON(u url.URL, data interface{}) {

	var req *http.Request

	proxyconnect := func () {
		configProxy()
		getJSON(u, data)
	}

	transport.Proxy = http.ProxyURL(&url.URL{Host: config.Proxy.Host+":"+config.Proxy.Port})

	if r, e := http.NewRequest("GET", u.String(), nil); e == nil {
		req = r
	} else {
		log.Fatal(e)
	}

	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Language", "en-gb")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Accept-Encoding", "deflate, br")
	req.Header.Add("Host", "www.britishhorseracing.com")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Referer", "https://www.britishhorseracing.com/")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15")

	if r, e := client.Do(req); e == nil {

		if r.StatusCode == 200 {
			fmt.Print("\033[1;" + colorM+"m")
			log.Println("get status:\033[0;" + colorM+"m", r.StatusCode, "\033[000m")
		} else {
			fmt.Print("\033[1;" + colorE+"m")
			log.Println("get status:\033[0;" + colorE+"m", r.StatusCode, "\033[000m")
			proxyconnect()
		}
	
		if b, e := ioutil.ReadAll(r.Body); e == nil {
	
			if e := json.Unmarshal(b, data); e != nil {
				log.Fatal(e)
			}
		}

	} else {

		se := e.Error()

		fmt.Print("\033[1;" + colorE+"m")
		log.Println("\033[0;" + colorE+"m"+se+"\033[000m")
		
		if strings.Contains(se, "proxyconnect") || strings.Contains(se, "connection")|| strings.Contains(se, "exceeded") {
			proxyconnect()
		}
	}
}

func getJSONRacecourses() []Racecourse {

	r := make([]Racecourse, 0)

	d := struct {
		Data []struct{
			CourseId int
			Name string
			Type string
			TrackHandedness string
			Region string
			Postcode string
			Latitude string
			Longitude string
			FirstRace string
			NextFixtureDate string
		} `json:"data"`
	} {}

	getJSON(genURLRacecourses(), &d)


	for _, d := range d.Data {

		var nextFixture time.Time

		latitude, err := strconv.ParseFloat(strings.TrimSpace(d.Latitude), 64)
		if err != nil {
			log.Fatal(err)
		}

		longitude, err := strconv.ParseFloat(strings.TrimSpace(d.Longitude), 64)
		if err != nil {
			log.Fatal(err)
		}

		if t, e := time.Parse("2006-01-02", strings.TrimSpace(d.NextFixtureDate));  e == nil {
			nextFixture = t
		}

		r = append(r, Racecourse{
			ID : d.CourseId,
			Name: d.Name,
			Type : strings.ToUpper(d.Type),
			Handedness : strings.ToUpper(d.TrackHandedness),
			Region : strings.ToUpper(d.Region),
			Post : strings.ToUpper(d.Postcode),
			Coordinate: GEO{Latitude : latitude, Longitude : longitude},
			FirstRace : datatimeParse(d.FirstRace),
			NextFixture : nextFixture,
		})
	}

	return r
}

func getJSONFixture(year, fixtureID int) Fixture {

	var r Fixture

	d := struct {
		Data []struct{
			FixtureID int 
			FixtureYear int 
			FixtureDate string 
			MetingID int 
			CourseID int 
			CourseName string 
			TicketsLink string 
			AlertLevel int 
			AbandonedReasonCode int 
			FixtureType string 
			FixtureSession string 
			RacingTrackType string 
			RacePlanningCode string 
			StewardsReport string  
			ResultsAvailable int 
			WeatherText string 
			WeatherUpdatedAt string 
			StallsText string 
			StallsUpdatedAt string 
			GoingText string 
			GoingUpdatedAt string 
			InspectionsText string 
			InspectionsUpdatedAt string
			RailText string 
			RailUpdatedAt string 
			OtherText string 
			OtherUpdatedAt string
			WateringText string
			WateringUpdatedAt string
			LastUpdated string
		}
	} {}

	getJSON(genURLFixture(year, fixtureID), &d)

	if len(d.Data) > 0 {

		r = Fixture {
			ID: d.Data[0].FixtureID,
			Date: dataParse(d.Data[0].FixtureDate),
			MetingID : d.Data[0].MetingID,
			RacecourseID: d.Data[0].CourseID,
			Racecourse: d.Data[0].CourseName,
			Abandoned : d.Data[0].AbandonedReasonCode > 0,
			Type: d.Data[0].FixtureType,
			Session: d.Data[0].FixtureSession,
			Surface: d.Data[0].RacingTrackType,
			Planning: d.Data[0].RacePlanningCode,
			Weather: TU{T: d.Data[0].WateringText, U: datatimeParse(d.Data[0].WateringUpdatedAt)},
			Stalls: TU{T: d.Data[0].StallsText, U: datatimeParse(d.Data[0].StallsUpdatedAt)},
			Going: TU{T: d.Data[0].GoingText, U: datatimeParse(d.Data[0].GoingUpdatedAt)},
			Inspection: TU{T: d.Data[0].InspectionsText, U: datatimeParse(d.Data[0].InspectionsText)},
			Rail: TU{T: d.Data[0].RailText, U: datatimeParse(d.Data[0].RailUpdatedAt)},
			Watering: TU{T: d.Data[0].WateringText, U: datatimeParse(d.Data[0].WateringUpdatedAt)},
			Other: TU{T: d.Data[0].OtherText, U: datatimeParse(d.Data[0].OtherUpdatedAt)},
			Updated: datatimeParse(d.Data[0].LastUpdated),
		}
	}

	return r
}

func getJSONRaces(year, fixtureID int) []Race {

	r := make([]Race, 0)

	d := struct {
		Data []struct{
			RaceId int
			YearOfRace int
			DivisionSequence int
			RaceDate string
			RaceTime string
			RaceName string
			AgeLimit string
			PrizeAmount int
			PrizeCurrency string
			RaceClass int
			RatingBand string
			RawDistanceText string
			DistanceValue int
			DistanceText string
			DistanceChange int
			RaceCriteriaRaceType string
			AbandonedReasonCode int
			BlackTypeRace int
			DistanceChangeText string
			Plus10 bool
			WinnersDetails []struct {
				Position int
				JockeyName string
				Trainername string
				SilkImage string
				RacehorseName string
			}

		}
	} {}

	getJSON(genURLRaces(year, fixtureID), &d)

	for _, d := range d.Data {

		r = append(r, Race{
			ID: d.RaceId,
			Division: d.DivisionSequence, 
			Datatime: datatimeParse(d.RaceDate+" "+d.RaceTime),
			Name: d.RaceName,
			Age: d.AgeLimit,
			Prize: d.PrizeAmount, 
			Currency: d.PrizeCurrency, 
			Class: d.RaceClass,
			Band: d.RatingBand,
			Distance: d.DistanceValue,
			Change: d.DistanceChange,
			Type: d.RaceCriteriaRaceType,
			Abandoned:  d.AbandonedReasonCode > 0,
			Black: d.BlackTypeRace > 0,
			Plus10 : d.Plus10,
		})
	}

	return r
}

func getJSONGoing(year, fixtureID int) Going {

	d := struct {
		Data struct{
			FixtureId int
			CourseId int
			FixtureYear int
			FixtureDate string
			FixtureType string
			Conditions struct {
				Ground int
				GroundText string
				GoingStick string
				GoingStickAvailable int
				GoingStickUpdatedAt string
				Rails string
				Stalls string
				WeatherComment string
				Other string
				Watering string
				WateringStatus string
			}

			ConditionsHistory []struct {
                FixtureId int
                CourseId int
                FixtureYear int
                FixtureDate string
                FixtureType string
                TrackType string
                FixtureSession string
				Conditions struct {
					Ground int
					GoingStick string
					GoingStickAvailable int
					GoingStickUpdatedAt string
					Rails string
					Stalls string
					WeatherComment string
					BookingComment string
					Other string
					Watering string
					WateringStatus string
					groundText struct {
                        Code int
                        Description string
					}
					CreationTimestamp string
				}

				Tracks [] struct {
					TrackId int
					RaceType string
				}
			}
		}
	} {}

	getJSON(genURLGoing(year, fixtureID), &d)

	stick := -1.0

	if f, e := strconv.ParseFloat(d.Data.Conditions.GoingStick, 64); e == nil {
		stick = f
	}

	return Going {
		FixtureID: d.Data.FixtureId,
		CourseID: d.Data.CourseId,
		Datatime: dataParse(d.Data.FixtureDate),
		Type: d.Data.FixtureType,
		Code: d.Data.Conditions.Ground,
		Ground: d.Data.Conditions.GroundText,
		Stick: FU{F: stick, U: datatimeParse(d.Data.Conditions.GoingStickUpdatedAt)},
		Rails: d.Data.Conditions.Rails,
		Stalls: d.Data.Conditions.Stalls,
		Weather: d.Data.Conditions.WeatherComment,
		Watering: d.Data.Conditions.Watering,
		WateringStatus: d.Data.Conditions.WateringStatus,
	}
}

func getJSONOfficials(year, fixtureID int) [] Official {

	r := make([]Official, 0)

	d := struct {
		Data []struct{
            Category string
            Officials []string
		}
	} {}

	getJSON(genURLOfficials(year, fixtureID), &d)

	for _, d := range d.Data {

		r = append(r, Official{
			Category: d.Category,
			Officials: d.Officials,
		})
	}

	return r
}

func getJSONRace(year, fixtureID int) Race {

	var r Race

	d := struct {
		Data []struct{
			RaceId int
            FixtureId int
            RaceNumber string
            YearOfRace int
            DivisionSequence int
            RaceDate string
            RaceTime string
			RaceName string
            AgeLimit string
            SexLimit string
            PrizeAmount int
            PrizeCurrency string
            DistanceValue int
            DistanceChange int
            RatingBand string
            RaceCriteriaRaceType string
            AbandonedReasonCode int
            BlackTypeRace int
            Plus10 bool
            RacingUK int
            RiderType string
            AnimalType string
            WinTime string
            Runners int
            MaxRunners int
            ResultsAvailable int
            RacecardAvailable int
            RaceCriteriaMinimumWeight int
            RaceCriteriaWeightsRaised int
            Challenger bool
		}
	} {}

	getJSON(genURLRace(year, fixtureID), &d)

	if len(d.Data) > 0 {
		
		number := -1
		if i, e := strconv.ParseInt(d.Data[0].RaceNumber, 10, 64); e == nil {
			number = int(i)
		}

		r = Race{
			ID: d.Data[0].RaceId,
			FixtureID: d.Data[0].FixtureId,
			Number: number,
			Division: d.Data[0].DivisionSequence,
			Name: d.Data[0].RaceName,
			Age: d.Data[0].AgeLimit,
			Sex: d.Data[0].SexLimit,
			Prize: d.Data[0].PrizeAmount,
			Currency: d.Data[0].PrizeCurrency,
			Band: d.Data[0].RatingBand,
			Datatime: datatimeParse(d.Data[0].RaceDate+" "+d.Data[0].RaceTime),
			Distance: d.Data[0].DistanceValue,
			Change: d.Data[0].DistanceChange,
			Type: d.Data[0].RaceCriteriaRaceType,
			Abandoned: d.Data[0].AbandonedReasonCode > 0,
			Black: d.Data[0].BlackTypeRace > 0,
			Plus10: d.Data[0].Plus10,
			RacingUK: d.Data[0].RacingUK > 0,
			Challenger: d.Data[0].Challenger,
			Rider: d.Data[0].RiderType,
			Animal: d.Data[0].AnimalType,
			WinTime: d.Data[0].WinTime,
			Runners: d.Data[0].Runners,
			MaxRunners: d.Data[0].MaxRunners,
			MinimumWeight: d.Data[0].RaceCriteriaMinimumWeight,
			WeightsRaised: d.Data[0].RaceCriteriaWeightsRaised,
		}
	}

	return r
}

func getJSONEntries(year, fixtureID int) []Entry {

	r := make([]Entry, 0)

	d := struct {
		Data []struct{
			RaceId int
            AnimalId int
            YearOfRace int
            DivisionSequence int
            RacehorseName string
            AgeYears int
            SexType string
            ClothNumber int
            DrawnStall int
            BhaRating int
            WeightValue string
			WeightText string
            PenaltyValue int
            NonRunnerDeclaredReason string
            nonRunnerDeclaredDate string
            nonRunnerDeclaredTime string
            Status string
            JockeyId int
            JockeyName string
			TrainerId int
            TrainerName string
			OwnerId int
            OwnerName string
            WeightsJockeyClaiming int
            HeadGearAbbreviation string
            WindSurgeryFirstRun int
            WbSilkCode string
            WbSilkDescription string
		}
	} {}

	getJSON(genURLEntries(year, fixtureID), &d)

	for _, d := range d.Data {

		silkCode := -1
		if i, e := strconv.ParseInt(d.WbSilkCode, 10, 64); e == nil {
			silkCode = int(i)
		}

		r = append(r, Entry {
			RaceID: d.RaceId,
			Horse: Participant{Name: d.RacehorseName, ID: d.AnimalId},
			Jockey: Participant{Name: d.JockeyName, ID: d.JockeyId},
			Trainer: Participant{Name: d.TrainerName, ID: d.TrainerId},
			Owner: Participant{Name: d.OwnerName, ID: d.OwnerId},
			Division: d.DivisionSequence,
			Age: d.AgeYears,
			Sex: d.SexType,
			Number: d.ClothNumber,
			Drawn: d.DrawnStall,
			Rating: d.BhaRating,
			Weight: d.WeightValue,
			Penalty: d.PenaltyValue,
			Nonrunner: Nonrunner{
				Horse: d.RacehorseName,
				Reason: d.NonRunnerDeclaredReason,
				Datatime: datatimeParse(d.nonRunnerDeclaredDate+" "+d.nonRunnerDeclaredTime),
			},
			Status: d.Status,
			JockeyClaim: d.WeightsJockeyClaiming,
			HeadGear: d.HeadGearAbbreviation,
			WindSurgeryFirstRun: d.WindSurgeryFirstRun,
			SilkCode: silkCode,
			SilkDescription: d.WbSilkDescription,
		})
	}

	return r
}
