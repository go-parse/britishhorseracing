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

fmt.Println(getJSONFixture(year, fixtureID))
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