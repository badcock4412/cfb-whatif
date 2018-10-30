package main

import (
	"fmt"
)

type Team struct {
	Name       string
	ShortName  string
	Conference string
	Division   string
}

type Game struct {
	AwayTeam *Team
	HomeTeam *Team
}

type TeamList []Team

func (tl *TeamList) Get(name string) (*Team, error) {
	for _, d := range *tl {
		if d.Name == name {
			return &d, nil
		}
	}
	return nil, fmt.Errorf("Team Name %s does not exist",name)
}

func main() {
	teams := TeamList{
		Team{
			Name:       "Georgia Tech",
			ShortName:  "GT",
			Conference: "ACC",
			Division:   "Coastal",
		},
		Team{
			Name:       "Virginia Tech",
			ShortName:  "VT",
			Conference: "ACC",
			Division:   "Coastal",
		},
		Team{
			Name:       "Virginia",
			ShortName:  "UVA",
			Conference: "ACC",
			Division:   "Coastal",
		},
		Team{
			Name:       "Miami",
			ShortName:  "UMia",
			Conference: "ACC",
			Division:   "Coastal",
		},
		Team{
			Name:       "North Carolina",
			ShortName:  "UNC",
			Conference: "ACC",
			Division:   "Coastal",
		},
		Team{
			Name:       "Duke",
			ShortName:  "Duke",
			Conference: "ACC",
			Division:   "Coastal",
		},
		Team{
			Name:       "Pittsburgh",
			ShortName:  "Pitt",
			Conference: "ACC",
			Division:   "Coastal",
		},
	}

	fmt.Println(teams.Get("Georga Tech"))
}
