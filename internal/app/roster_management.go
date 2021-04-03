package app

import (
	"fmt"

	"github.com/TheTh1rt33nth/Nexus-octo-engine/pkg/model"
)

var DetachCost = map[model.Detach]int{
	model.Patrol:         2,
	model.Batallion:      3,
	model.Brigade:        4,
	model.Vanguard:       3,
	model.Spearhead:      3,
	model.Outrider:       3,
	model.SupremeCommand: 0,
	model.SuperHeavy:     3,
	model.SuperHeavyAux:  3,
	model.Fortification:  1,
	model.Auxillary:      2,
}

func CreateRoster(user int, campaign int, name string) {
	/*dummyCampaign := model.Campaign{
			Name:       "Battle for Rum bottle",
	    Descr: "Klinika Valhalla looking for something to drink",
			StartingPL: 25,
		}*/

	// if /*check DBcontext if campaign already have a roster associated with tis name*/ {
	/* fmt.Println("Go fuck yourself, you are already in this shit")
	}
	else
	{ */
	/*roster := model.Army{
			Name:        name,
			UserID:      user,
			CampaignID:  campaign,
			SupplyLimit: dummyCampaign.StartingPL, // ask context for ACTUAL campaign by id
			SupplyUsed:  0,
	    RequisitionPoints: dummyCampaign.StartingRP,
			// send dis shit to DBcontext, we're done here
		}*/
}

func CountCP(roster *model.Army) int {
	cp := 0
	pl := 0
	dummydetaches := []model.Detachment{} //get actual detaches from DB
	if len(dummydetaches) == 0 {
		return 0
	}
	for _, d := range dummydetaches {
		cp -= d.CpCost
		pl += d.PowerLevel
	}
	cp += CPbyPL(pl)
	return cp
}

//spend 1 requisition point for additional 5 supply limit
func SpendRPonSL(armyID int) {
	//get actual roster by id
	dummyRoster := model.Army{
		Name:              "Klinika Valhalla",
		UserID:            1,
		CampaignID:        1,
		SupplyLimit:       25,
		SupplyUsed:        0,
		RequisitionPoints: 5,
	}
	rosterObj := dummyRoster // but actually ask DB for real roster by armyID
	if rosterObj.RequisitionPoints > 0 {
		rosterObj.RequisitionPoints -= 1
		rosterObj.SupplyLimit += 1
	} else {
		fmt.Println("Not enough RP. Fight harder, commander!")
	}
}

func PLtoCP(pl int) int {
	switch {
	case pl <= 50:
		return 3
	case pl <= 100:
		return 6
	case pl <= 200:
		return 12
	case pl > 200:
		return 18
	}

	return 0
}

func AddDetachment(roster int, fraction string, detachType model.Detach) {
	dummyRoster := model.Army{
		Name:              "Klinika Valhalla",
		UserID:            1,
		CampaignID:        1,
		SupplyLimit:       25,
		SupplyUsed:        0,
		RequisitionPoints: 5,
	}
	rosterObj := dummyRoster // ask context for roster by id
	//RosterCP := CountCP(&rosterObj)
	//RosterCP -= DetachCost[detachType]
	/*Detachhment := model.Detachment{
		ArmyID:         roster,
		Fraction:       fraction,
		CpCost:         DetachCost[detachType],
		DetachmentType: detachType,
	}*/
	rosterObj.AvailableCP -= DetachCost[detachType] // and update this stuff in DB
	// send this to DBcontext, seems okayish  }
}

// validateRoster
