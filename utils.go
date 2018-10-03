//Model
package pgi

//Data model for PGI
type Team struct {
	Id       int
	Name     string
	Regional string
}

type PointDistribution struct {
	Placement int
	Point     int
}

type Match struct {
	Id                 int
	MatchName          string
	MapName            string
	ListOfMatchDetails []MatchDetails
}

/////////////////////////////////////////////////////////////
type MatchDetails struct {
	Team       Team
	Detail     PointDistribution
	Kill       int
	TotalPoint int
}

func (dt *MatchDetails) caculatePoint() {
	dt.TotalPoint = dt.Detail.Point + dt.Kill*15
}

/////////////////////////////////////////////////////////////

type ResultItem struct {
	Team               Team
	ListOfMatchDetails []MatchDetails
	TotalPoint         int
}

type ByPoint []ResultItem

func (this ByPoint) Len() int {
	return len(this)
}

func (this ByPoint) Less(i, j int) bool {
	return this[i].TotalPoint > this[j].TotalPoint
}

func (this ByPoint) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

/////////////////////////////////////////////////////////////

//Global data
var listOfTeam []Team = make([]Team, 0)
var listOfPointDistribution []PointDistribution = make([]PointDistribution, 0)
var listOfMatchesTpp []Match = make([]Match, 0)
var listOfMatchesFpp []Match = make([]Match, 0)

func MakeTeam(id int, name string, regional string) Team {
	return Team{id, name, regional}
}

func MakePointDistribution(placement int, point int) PointDistribution {
	return PointDistribution{placement, point}
}

func SearchTeam(id int) Team {
	for _, item := range listOfTeam {
		if id == item.Id {
			return item
		}
	}

	return Team{0, "", ""}
}

func SearchPointDistribution(placement int) PointDistribution {
	for _, item := range listOfPointDistribution {
		if placement == item.Placement {
			return item
		}
	}

	return PointDistribution{0, 15}
}

/////////////////////////////////////////////////////////////

//Testing
func DisplayTeams() {
	if listOfTeam != nil {
		fmt.Println("Total team: ", len(listOfTeam))

		for _, item := range listOfTeam {
			fmt.Println(item.Id, " - ", item.Name, " - ", item.Regional)
		}

		fmt.Print("\n")
	}
}

func DisplayPointDistribution() {
	if listOfPointDistribution != nil {
		fmt.Println("Total point distribution: ", len(listOfPointDistribution))

		for _, item := range listOfPointDistribution {
			fmt.Println(item.Placement, " - ", item.Point)
		}
		fmt.Print("\n")
	}
}

func DisplayMatches() {
	if listOfMatchesTpp != nil {
		fmt.Println("Total matches: ", len(listOfMatchesTpp))

		for _, item := range listOfMatchesTpp {
			fmt.Print(DisplayMatch(item))
			fmt.Print("\n")
		}
	}
}
