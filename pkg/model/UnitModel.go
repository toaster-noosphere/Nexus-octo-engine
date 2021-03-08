package model

type Unit struct {
	id            int
	armyID        int
	name          string
	unitType      string
	role          int    //move to enum once i learn how to enum
	keywords      string //logically it is a string[]. For DB it is a string
	experience    int
	crusadePoints int
	powerLevel    int
}
