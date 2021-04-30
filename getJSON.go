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

var getJSONData interface{}

func getJSON(u *url.URL) {

	proxyconnect := func () {
		configProxy()
		getJSON(u)
	}

	var req *http.Request
	var res *http.Response

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
		res = r
	} else {

		se := e.Error()

		fmt.Print("\033[1;" + colorE+"m")
		
		if strings.Contains(se, "proxyconnect") {
			log.Println("\033[0;" + colorE+"m"+se+"\033[000m")
			proxyconnect()
		} else {
			log.Fatal("\033[0;" + colorE+"m"+se+"\033[000m")
		}
	}

	if res.StatusCode == 200 {
		fmt.Print("\033[1;" + colorM+"m")
		log.Println("get status:\033[0;" + colorM+"m", res.StatusCode, "\033[000m")
	} else {
		fmt.Print("\033[1;" + colorE+"m")
		log.Println("get status:\033[0;" + colorE+"m", res.StatusCode, "\033[000m")
		proxyconnect()
	}

	if b, e := ioutil.ReadAll(res.Body); e == nil {

		if e := json.Unmarshal(b, &getJSONData); e != nil {
			log.Fatal(e)
		}
	}
}

func getJSONRacecourses() []Racecourse {

	r := make([]Racecourse, 0)

	d := struct {
		Data []struct{
			ID int `json:"courseId"`
			Name string `json:"name"`
			Type string `json:"type"`
			Handedness string `json:"trackHandedness"`
			Region string `json:"region"`
			Post string `json:"postcode"`
			Latitude string `json:"latitude"`
			Longitude string `json:"longitude"`
			FirstRace string `json:"firstRace"`
			NextFixture string `json:"nextFixtureDate"`
		} `json:"data"`
	} {}

	u := genURLRacecourses()

	getJSONData = &d

	getJSON(&u)


	for _, d := range d.Data {

		var firstRace time.Time
		var nextFixture time.Time

		latitude, err := strconv.ParseFloat(strings.TrimSpace(d.Latitude), 64)
		if err != nil {
			log.Fatal(err)
		}

		longitude, err := strconv.ParseFloat(strings.TrimSpace(d.Longitude), 64)
		if err != nil {
			log.Fatal(err)
		}

		if t, e := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(d.FirstRace));  e == nil {
			firstRace = t
		}

		if t, e := time.Parse("2006-01-02", strings.TrimSpace(d.NextFixture));  e == nil {
			nextFixture = t
		}

		r = append(r, Racecourse{
			ID : d.ID,
			Name: d.Name,
			Type : strings.ToUpper(d.Type),
			Handedness : strings.ToUpper(d.Handedness),
			Region : strings.ToUpper(d.Region),
			Post : strings.ToUpper(d.Post),
			Latitude : latitude,
			Longitude : longitude,
			FirstRace : firstRace,
			NextFixture : nextFixture,
		})
	}

	return r

}
