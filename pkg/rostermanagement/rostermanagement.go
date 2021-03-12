package rostermanagement

import (
	"fmt"
	"github.com/TheTh1rt33nth/Nexus-octo-engine/pkg/model/Models.go"
)

func CreateRoster(user int, campaign int, name string) {
	dummyCampaign := CampaignModel{
		name:       "Battle for Rum bottle",
		startingPL: 25,
	}

	// if /*check DBcontext if campaign already have a roster associated with tis name*/ {
	/* fmt.Println("Go fuck yourself, you are already in this shit")
	}
	else
	{ */
	roster := ArmyModel{
		name:        name,
		userID:      user,
		campaignID:  campaign,
		supplyLimit: dummyCampaign.startingPL, // ask context for ACTUAL campaign by id
		supplyUsed:  0,                        // tomodel
		// }
		// send dis shit to DBcontext, we're done here
	}
}

func CountCP(roster ArmyModel) int {
	cp = 0
	pl = 0
	dummydetaches := []DetachmentModel{}
	if len(dummydetaches) {
		return 0
	}
	for _, d := range dummydetaches {
		cp -= d.cpCost
		pl += d.powerLevel
	}
	cp += CPbyPL(pl)
	return cp
}

func CPbyPL(pl int) int {
	switch pl {
	case pl <= 50:
		return 3
	case pl <= 100:
		return 6
	case pl <= 200:
		return 12
	case pl > 200:
		return 18
	}
}

func DetachCost(dType string) int {
	switch dType {
	case "Patrol":
		return 2
	case "Batallion":
		return 3
	case "Brigade":
		return 4
	case "Vanguard":
		return 3
	case "Spearhead":
		return 3
	case "Outrider":
		return 3
	case "Supreme Command":
		return 0
	case "Super-heavy":
		return 3
	case "Super-heavy auxillary":
		return 3
	case "Fortification network":
		return 1
	case "Auxillary support":
		return 2
	}
}

func AddDetachment(roster int, fraction string, detachType string) {
	dummyroster := ArmyModel{
		name:        "Klinika Valhalla",
		userId:      1,
		campaignID:  1,
		supplyLimit: 25,
	}
	var rosterObj ArmyModel
	rosterObj = dummyRoster // ask context for roster by id
	rosterCP = CountCp(rosterObj)
	rosterCP -= DetachCost(detachType)
	detachhment := DetachmentModel{
		armyID:         roster,
		fraction:       fraction,
		cpCost:         DetachCost(detachType),
		detachmentType: detachType,
	}
	rosterObj.availableCP -= DetachCost(detachType) // and update this stuff in DB
	// send this to DBcontext, seems okayish  }
}

// validateRoster
