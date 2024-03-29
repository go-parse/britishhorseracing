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

type Participant struct {
	Name string `json:"NAME"`
	ID int `json:"ID"` 
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
	Bcs bool `json:"BCS"`
	Abandoned bool `json:"ABANDONED"`
	Region string `json:"REGION"`
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
	RacecardAvailable bool `json:"RACECARD_AVAILABLE"`
	EntriesAvailable bool `json:"ENTRIES_AVAILABLE"`
	BlackTypeRaces bool `json:"BLACK_TYPERACES"`
	ResultsAvailable bool `json:"RESULTS_AVAILABLE"`
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

type Nonrunner struct {
	Horse string `json:"HORSE"`
	Reason string `json:"REASON"`
	Datatime time.Time `json:"DATATIME"`
}

type Nonrunners struct {
	Title string `json:"Title"`
	Datatime time.Time `json:"DATATIME"`
	NR []Nonrunner `json:"NR"`
}

type Entry struct {
	RaceID int `json:"RACE_ID"`
	Horse Participant `json:"HORSE"`
	Jockey Participant `json:"JOCKEY"`
	Trainer Participant `json:"TRAINER"`
	Owner Participant `json:"OWNER"`
	Division int `json:"DIVISION"`
	Age int `json:"AGE"`
	Sex string `json:"SEX"`
	Number int `json:"Number"`
	Drawn int `json:"Drawn"`
	Rating int `json:"RATING"`
	Weight string `json:"Weight"`
	Penalty int `json:"Penalty"`
	Nonrunner Nonrunner `json:"NON_RUNNER"`
	Status string `json:"STATUS"`
	JockeyClaim int `json:"JOCKEY_CLAIM"`
	HeadGear string `json:"HEAD_GEAR"`
	WindSurgeryFirstRun int `json:"WIND_SURGERY_FIRST_RUN"`
	SilkCode int `json:"SILK_CODE"`
	SilkDescription string `json:"SILK_DESCRIPTION"`
}