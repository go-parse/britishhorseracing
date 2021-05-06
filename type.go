package main

import "time"

type TU struct {
	T string `json:"T"` // Text
	U time.Time `json:"U"` // Updated 
}
type FU struct {
	F float64 `json:"T"` // Float
	U time.Time `json:"U"` // Updated 
}
type GEO  struct {
	Latitude float64 `json:"LATITUDE"` 
	Longitude float64 `json:"LONGITUDE"`
}

type Racecourse struct {
	ID int `json:"ID"`
	Name string `json:"NAME"`
	Type string `json:"TYPE"`
	Handedness string `json:"HANDEDNESS"`
	Region string `json:"REGION"`
	Post string `json:"POST"`
	Coordinate GEO `json:"COORDINATE"`
	FirstRace time.Time `json:"FIRST_RACE"`
	NextFixture time.Time  `json:"NEXT_FIXTURE"`
}

type Fixture struct {
	ID int `json:"ID"`
	Date time.Time `json:"DATE"`
	MetingID int `json:"METING_ID"`
	RacecourseID int `json:"RACECOURSE_ID"`
	Racecourse string `json:"RACECOURSE"`
	Abandoned bool `json:"ABANDONED"`
	Type string `json:"TYPE"`
	Session string `json:"SESSION"`
	Surface string `json:"SURFACE"`
	Planning string `json:"PLANNING"`
	Weather TU `json:"WEATHER"`
	Stalls TU `json:"STALLS"`
	Going TU `json:"GOING"`
	Inspection TU `json:"INSPECTION"`
	Rail TU `json:"RAIL"`
	Watering TU `json:"WATERING"`
	Other TU `json:"OTHER"`
	Updated time.Time `json:"UPDATED"`
}

type Race struct {
	ID int `json:"ID"`
	FixtureID int `json:"FIXTURE_ID"`
	Number int `json:"NUMBER"`
	Division int `json:"DIVISION"`
	Datatime time.Time `json:"DATATIME"`
	Name string `json:"NAME"`
	Age string `json:"AGE"`
	Sex string `json:"SEX"`
	Prize int `json:"PRIZE"`
	Currency string `json:"CURRENCY"`
	Class int `json:"CLASS"`
	Band string `json:"BAND"`
	Distance int `json:"DISTANCE"`
	Change int `json:"CHANGE"`
	Type string `json:"TYPE"`
	Abandoned bool `json:"ABANDONED"`
	Black bool `json:"BLACK"`
	Plus10 bool `json:"PLUS10"`
	RacingUK bool `json:"RACING_UK"`
	Challenger bool `json:"CHALLENGER"`
	Rider string `json:"RIDER"`
	Animal string `json:"ANIMAL"`
	WinTime string `json:"WIN_TIME"`
	Runners int `json:"RUNNERS"`
	MaxRunners int `json:"MAX_RUNNERS"`
	MinimumWeight int `json:"MINIMUM_WEIGHT"`
	WeightsRaised int `json:"WEIGHTS_RAISED"`
}
type Going struct {
	FixtureID int `json:"FIXTURE_ID"`
	CourseID int `json:"COURSE_ID"`
	Datatime time.Time `json:"DATATIME"`
	Type string `json:"TYPE"`
	Code int `json:"CODE"`
	Ground string `json:"GROUND"`
	Stick FU `json:"STICK"`
	Rails string `json:"RAILS"`
	Stalls string `json:"STALLS"`
	Weather string `json:"WEATHER"`
	Watering string `json:"WATERING"`
	WateringStatus string `json:"WATERING_STATUS"`
}

type Official struct {
	Category string `json:"CATEGORY"`
	Officials []string `json:"OFFICIALS"`
}