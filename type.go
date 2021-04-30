package main

import "time"

type Racecourse struct {
	ID int `json:"ID"`
	Name string `json:"NAME"`
	Type string `json:"TYPE"`
	Handedness string `json:"HANDEDNESS"`
	Region string `json:"REGION"`
	Post string `json:"POST"`
	Latitude float64 `json:"LATITUDE"`
	Longitude float64 `json:"LONGITUDE"`
	FirstRace time.Time `json:"FIRST_RACE"`
	NextFixture time.Time  `json:"NEXT_FIXTURE"`
}