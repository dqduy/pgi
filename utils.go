//Model
package pgi

import (
	"fmt"
	"strconv"
)

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

func (dt *MatchDetails) CaculatePoint() {
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
var ListOfTeam []Team = make([]Team, 0)
var ListOfPointDistribution []PointDistribution = make([]PointDistribution, 0)
var ListOfMatchesTpp []Match = make([]Match, 0)
var ListOfMatchesFpp []Match = make([]Match, 0)

func MakeTeam(id int, name string, regional string) Team {
	return Team{id, name, regional}
}

func MakePointDistribution(placement int, point int) PointDistribution {
	return PointDistribution{placement, point}
}

func SearchTeam(id int) Team {
	for _, item := range ListOfTeam {
		if id == item.Id {
			return item
		}
	}

	return Team{0, "", ""}
}

func SearchPointDistribution(placement int) PointDistribution {
	for _, item := range ListOfPointDistribution {
		if placement == item.Placement {
			return item
		}
	}

	return PointDistribution{0, 15}
}

/////////////////////////////////////////////////////////////

//Testing
func DisplayTeams() {
	if ListOfTeam != nil {
		fmt.Println("Total team: ", len(ListOfTeam))

		for _, item := range ListOfTeam {
			fmt.Println(item.Id, " - ", item.Name, " - ", item.Regional)
		}

		fmt.Print("\n")
	}
}

func DisplayPointDistribution() {
	if ListOfPointDistribution != nil {
		fmt.Println("Total point distribution: ", len(ListOfPointDistribution))

		for _, item := range ListOfPointDistribution {
			fmt.Println(item.Placement, " - ", item.Point)
		}
		fmt.Print("\n")
	}
}

func DisplayMatch(match Match) string {
	var result = ""

	result = "ID:         " + strconv.Itoa(match.Id) + "\n" +
		"Match Name: " + match.MatchName + "\n" +
		"Map:        " + match.MapName

	result += "\n\tDetails: \n"

	for _, item := range match.ListOfMatchDetails {
		result += item.Team.Name + ", " +
			item.Team.Regional + ", " +
			strconv.Itoa(item.Kill) + ", " +
			strconv.Itoa(item.TotalPoint) + "\n"
	}

	return result
}

func DisplayMatches() {
	if ListOfMatchesTpp != nil {
		fmt.Println("Total matches: ", len(ListOfMatchesTpp))

		for _, item := range ListOfMatchesTpp {
			fmt.Print(DisplayMatch(item))
			fmt.Print("\n")
		}
	}
}
