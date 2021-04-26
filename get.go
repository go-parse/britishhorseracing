package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

var client = http.Client{Transport: transport, Timeout: time.Minute * 5}

func get(link string, data interface{}) {

	var req *http.Request
	var res *http.Response

	transport.Proxy = http.ProxyURL(&url.URL{Host: Proxy})

	if r, e := http.NewRequest("GET", link, nil); e == nil {
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
		log.Fatal(e)
	}

	log.Println("get:", res.StatusCode, link)

	if b, e := ioutil.ReadAll(res.Body); e == nil {
		if e := json.Unmarshal(b, &data); e != nil {
			log.Fatal(e)
		}
	} else {
		log.Fatal(e)
	}
}
