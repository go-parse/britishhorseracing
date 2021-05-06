# britishhorseracing
 
## Examples
#### **Get racecourses**
```Go
for _, d := range  getJSONRacecourses() {
	fmt.Println("ID",  d.ID)
	fmt.Println("Name",  d.Name)
	fmt.Println("Type",  d.Type)
	fmt.Println("Handedness",  d.Handedness)
	fmt.Println("Region",  d.Region)
	fmt.Println("Post",  d.Post)
	fmt.Println("Coordinate",  d.Coordinate)
	fmt.Println("FirstRace",  d.FirstRace)
	fmt.Println("NextFixture",  d.NextFixture)
}
```
#### **Get fixture**
```Go
year := 2021
fixtureID := 12763

fixture := getJSONFixture(year, fixtureID)

fmt.Println("ID:", fixture.ID)
fmt.Println("Date:", fixture.Date)
fmt.Println("MetingID:", fixture.MetingID)
fmt.Println("RacecourseID:", fixture.RacecourseID)
fmt.Println("Racecourse:", fixture.Racecourse)
fmt.Println("Abandoned:", fixture.Abandoned)
fmt.Println("Type:", fixture.Type)
fmt.Println("Type:", fixture.Type)
fmt.Println("Session:", fixture.Session)
fmt.Println("Surface:", fixture.Surface)
fmt.Println("Planning:", fixture.Planning)
fmt.Println("Weather:", fixture.Weather)
fmt.Println("Stalls:", fixture.Stalls)
fmt.Println("Going:", fixture.Going)
fmt.Println("Inspection:", fixture.Inspection)
fmt.Println("Rail:", fixture.Rail)
fmt.Println("Watering:", fixture.Watering)
fmt.Println("Other:", fixture.Other)
fmt.Println("Updated:", fixture.Updated)
```
#### **Get races**
```Go
year := 2021
fixtureID := 12763

for _, d := range  getJSONRaces(year, fixtureID) {
	fmt.Println("ID:", d.ID)
	fmt.Println("Year:", d.Year)
	fmt.Println("Division:", d.Division)
	fmt.Println("Datatime:", d.Datatime)
	fmt.Println("Name:", d.Name)
	fmt.Println("Age:", d.Age)
	fmt.Println("Prize:", d.Prize)
	fmt.Println("Currency:", d.Currency)
	fmt.Println("Class:", d.Class)
	fmt.Println("Band:", d.Band)
	fmt.Println("Distance:", d.Distance)
	fmt.Println("Change:", d.Change)
	fmt.Println("Type:", d.Type)
	fmt.Println("Abandoned:", d.Abandoned)
	fmt.Println("Black:", d.Black)
	fmt.Println("Plus10:", d.Plus10)
}
```

#### **Get going**
```Go
year := 2021
fixtureID := 12763

going := getJSONGoing(year, fixtureID)
	
fmt.Println("CourseID:", going.CourseID)
fmt.Println("CourseID:", going.CourseID)
fmt.Println("Datatime:", going.Datatime)
fmt.Println("Type:", going.Type)
fmt.Println("Code:", going.Code)
fmt.Println("Ground:", going.Ground)
fmt.Println("Stick:", going.Stick)
fmt.Println("Rails:", going.Rails)
fmt.Println("Stalls:", going.Stalls)
fmt.Println("Weather:", going.Weather)
fmt.Println("Watering:", going.Watering)
fmt.Println("WateringStatus:", going.WateringStatus)
```

#### **Get officials**
```Go
year := 2021
fixtureID := 12763

for _, d := range getJSONOfficials(year, fixtureID) {

	fmt.Println("Category:", d.Category)

	for _, o := range d.Officials {
		fmt.Println(o)
	}
		
	fmt.Println("_______________")
}
```
#### **Get race**
```Go
year := 2021
fixtureID := 12763

race := getJSONRace(year, fixtureID)

fmt.Println("ID:", race.ID)
fmt.Println("FixtureID:", race.FixtureID)
fmt.Println("Number:", race.Number)
fmt.Println("Division:", race.Division)
fmt.Println("Name:", race.Name)
fmt.Println("Age:", race.Age)
fmt.Println("Sex:", race.Sex)
fmt.Println("Prize:", race.Prize)
fmt.Println("Currency:", race.Currency)
fmt.Println("Band:", race.Band)
fmt.Println("Datatime:", race.Datatime)
fmt.Println("Distance:", race.Distance)
fmt.Println("Change:", race.Change)
fmt.Println("Type:", race.Type)
fmt.Println("Abandoned:", race.Abandoned)
fmt.Println("Black:", race.Black)
fmt.Println("Plus10:", race.Plus10)
fmt.Println("RacingUK:", race.RacingUK)
fmt.Println("Challenger:", race.Challenger)
fmt.Println("Rider:", race.Rider)
fmt.Println("Animal:", race.Animal)
fmt.Println("WinTime:", race.WinTime)
fmt.Println("Runners:", race.Runners)
fmt.Println("MaxRunners:", race.MaxRunners)
fmt.Println("MinimumWeight:", race.MinimumWeight)
fmt.Println("WeightsRaised:", race.WeightsRaised)
```
#### **Get entries**
```Go
year := 2021
fixtureID := 38656

for _, d := range getJSONEntries(year, fixtureID) {
	fmt.Println("RaceID:",d.RaceID)
	fmt.Println("Horse ID:",d.Horse.ID)
	fmt.Println("Horse Name:",d.Horse.Name)
	fmt.Println("Jockey ID:",d.Jockey.ID)
	fmt.Println("Jockey Name:",d.Jockey.Name)
	fmt.Println("Trainer ID:",d.Trainer.ID)
	fmt.Println("Trainer Name:",d.Trainer.Name)
	fmt.Println("Owner ID:",d.Owner.ID)
	fmt.Println("Owner Name:",d.Owner.Name)
	fmt.Println("Division:",d.Division)
	fmt.Println("Age:",d.Age)
	fmt.Println("Sex:",d.Sex)
	fmt.Println("Number:",d.Number)
	fmt.Println("Drawn:",d.Drawn)
	fmt.Println("Rating:",d.Rating)
	fmt.Println("Weight:",d.Weight)
	fmt.Println("Penalty:",d.Penalty)
	fmt.Println("Nonrunner Horse:",d.Nonrunner.Horse)
	fmt.Println("Nonrunner Reason:",d.Nonrunner.Reason)
	fmt.Println("Nonrunner Datatime:",d.Nonrunner.Datatime)
	fmt.Println("Status:",d.Status)
	fmt.Println("JockeyClaim:",d.JockeyClaim)
	fmt.Println("HeadGear:",d.HeadGear)
	fmt.Println("WindSurgeryFirstRun:",d.WindSurgeryFirstRun)
	fmt.Println("SilkCode:",d.SilkCode)
	fmt.Println("SilkDescription:",d.SilkDescription)
	fmt.Println("_______________")
}
```
#### **Get nonrunners**
```Go
year := 2021
fixtureID := 38656

nonrunners := getJSONNonrunners(year, fixtureID)

fmt.Println("Title:", nonrunners.Title)
fmt.Println("Datatime:", nonrunners.Datatime)

for _, d := range nonrunners.NR {
	fmt.Println("Horse:", d.Horse)
	fmt.Println("Reason:", d.Reason)
	fmt.Println("Datatime:", d.Datatime)
	fmt.Println("_______________")
}
```
#### **Get fixtures**
```Go
fields := make([]string, 0)
fields = append(fields, "courseId")
fields = append(fields, "fixtureId")
fields = append(fields, "meetingId")
fields = append(fields, "fixtureDate")
fields = append(fields, "firstRaceTime")
fields = append(fields, "fixtureSession")
fields = append(fields, "bcsEvent")
fields = append(fields, "fixtureType")
fields = append(fields, "highlightTitle")
fields = append(fields, "firstRace")
fields = append(fields, "courseName")
fields = append(fields, "abandonedReasonCode")

fixturesForMonth := genURLFixturesForMonth(1, 3, 2021, 5, true, fields)

for _, d := range getJSONFixtures(fixturesForMonth){
	fmt.Println("ID:", d.ID)
	fmt.Println("MetingID:", d.MetingID)
	fmt.Println("RacecourseID:", d.RacecourseID)
	fmt.Println("Racecourse:", d.Racecourse)
	fmt.Println("Date:", d.Date)
	fmt.Println("Bcs:", d.Bcs)
	fmt.Println("Abandoned:", d.Abandoned)
	fmt.Println("Region:", d.Region)
	fmt.Println("Type:", d.Type)
	fmt.Println("RacecardAvailable:", d.RacecardAvailable)
	fmt.Println("EntriesAvailable:", d.EntriesAvailable)
	fmt.Println("BlackTypeRaces:", d.BlackTypeRaces)
	fmt.Println("ResultsAvailable:", d.ResultsAvailable)
	fmt.Println("_______________")
}
```