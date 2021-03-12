package app

import "github.com/TheTh1rt33nth/Nexus-octo-engine/pkg/model"

func CreateRoster(user int, campaign int, name string) {
	dummyCampaign := model.Campaign{
		Name:       "Battle for Rum bottle",
		StartingPL: 25,
	}

	// if /*check DBcontext if campaign already have a roster associated with tis name*/ {
	/* fmt.Println("Go fuck yourself, you are already in this shit")
	}
	else
	{ */
	roster := model.Army{
		Name:        name,
		UserID:      user,
		CampaignID:  campaign,
		SupplyLimit: dummyCampaign.StartingPL, // ask context for ACTUAL campaign by id
		SupplyUsed:  0,                        // tomodel
		// send dis shit to DBcontext, we're done here
	}
}

func CountCP(roster model.Army) int {
	cp := 0
	pl := 0
	dummydetaches := []model.Detachment{}
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

func CPbyPL(pl int) int {
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

func DetachCost(dType string) int {
	switch {
	case dType == "Patrol":
		return 2
	case dType == "Batallion":
		return 3
	case dType == "Brigade":
		return 4
	case dType == "Vanguard":
		return 3
	case dType == "Spearhead":
		return 3
	case dType == "Outrider":
		return 3
	case dType == "Supreme Command":
		return 0
	case dType == "Super-heavy":
		return 3
	case dType == "Super-heavy auxillary":
		return 3
	case dType == "Fortification network":
		return 1
	case dType == "Auxillary support":
		return 2
	}

	return 0
}

func AddDetachment(roster int, fraction string, detachType string) {
	dummyroster := model.Army{
		Name:        "Klinika Valhalla",
		UserID:      1,
		CampaignID:  1,
		SupplyLimit: 25,
	}
	var rosterObj model.Army
	rosterObj = dummyRoster // ask context for roster by id
	RosterCP = CountCp(rosterObj)
	RosterCP -= DetachCost(detachType)
	Detachhment := model.Detachment{
		ArmyID:         roster,
		Fraction:       fraction,
		CpCost:         DetachCost(detachType),
		DetachmentType: detachType,
	}
	rosterObj.AvailableCP -= DetachCost(detachType) // and update this stuff in DB
	// send this to DBcontext, seems okayish  }
}

// validateRoster
