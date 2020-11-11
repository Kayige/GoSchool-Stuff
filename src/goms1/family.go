package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// People struct to store json
type People struct {
	Name struct {
		First, Last string
	}
	DOB struct {
		Day, Month, Year int
	}
	Contact struct {
		Email  string
		Mobile int
	}
}

// People2 Struct
type People2 struct {
}

func main() {
	var members []People
	jsonString := `[{
		"Name": {
			"first": "Wei-Meng",
			"last": "Lee"
		},
		"DOB": {
			"Day": 10,
			"Month": 8,
			"Year": 1990
		},
		"Contact": {
			"Email": "weimenglee@learn2develop.net",
			"Mobile": 1234567
		}
	},
	{
		"Name": {
			"first": "Ah Kow",
			"last": "Lee"
		},
		"DOB": {
			"Day": 4,
			"Month": 12,
			"Year": 1980
		},
		"Contact": {
			"Email": "ahkowlee@gmail.com",
			"Mobile": 7654321
		}
	},
	{
		"Name": {
			"first": "Ah Ngeow",
			"last": "Lee"
		},
		"DOB": {
			"Day": 3,
			"Month": 12,
			"Year": 1980
		},
		"Contact": {
			"Email": "ahkowlee@gmail.com",
			"Mobile": 7654321
		}
	}]`

	json.Unmarshal([]byte(jsonString), &members)
	//fmt.Println(len(person))

	sort.SliceStable(members, func(i, j int) bool {
		if members[i].DOB.Year != members[j].DOB.Year {
			return members[i].DOB.Year < members[j].DOB.Year
		}
		if members[i].DOB.Month != members[j].DOB.Month {
			return members[i].DOB.Month < members[j].DOB.Month
		}
		return members[i].DOB.Day < members[j].DOB.Day
	})

	for _, member := range members {
		fmt.Println(member.Name.First)
		fmt.Println(member.Name.Last)
	}

}
