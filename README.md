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

fixture := getJSONFixture(2021, 12763)

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

going := getJSONGoing(2021, 715)
	
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

for _, d := range getJSONOfficials(2021, 12763) {

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

race := getJSONRace(2021, 21301)

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