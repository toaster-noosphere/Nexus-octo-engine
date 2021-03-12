package model

type Army struct {
	id          int
	userID      int
	campaignID  int
	lore        string
	name        string
	supplyLimit int
	supplyUsed  int
	availableCP int
	validated   bool
}

type Advancement struct {
	id       int
	advType  string
	descr    string
	crusCost int
}

type CampaignAchievements struct {
	achievementID string
	campaignID    int
	userID        int
}

type CampaignEvent struct {
	id      int
	name    string
	number  int
	player1 int
	player2 int
	victor  int
	lore    string
}

type CampaignEvents struct {
	eventID    int
	campaignID int
}

type Campaign struct {
	id       int
	name     string
	descr    string
	ispublic bool
}

type CampaignPlayers struct {
	userID     int
	campaignID int
	u_role     string
}

type DetachmentModel struct {
	id             int
	armyID         int
	fraction       int
	cpCost         int
	detachmentType string
	powerLevel     int
}

type EventTallies struct {
	eventID int
	tallyID int
	unitID  int
}

type Tally struct {
	id    int
	tally int
	value int
}

type UnitGrade struct {
	unitID    int
	advanceID int
}

type Unit struct {
	id            int
	detachmentID  int
	name          string
	unitType      string
	role          string
	keywords      string // logically it is a string[]. For DB it is a string
	experience    int
	crusadePoints int
	powerLevel    int
}

type User struct {
	id       int
	name     string
	password string
}
