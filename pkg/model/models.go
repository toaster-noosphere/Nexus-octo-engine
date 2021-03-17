package model

type Detach string

const (
	Patrol         Detach = "patrol"
	Batallion      Detach = "batallion"
	Brigade        Detach = "Brigade"
	Vanguard       Detach = "Vanguard"
	Spearhead      Detach = "Spearhead"
	Outrider       Detach = "Outrider"
	SupremeCommand Detach = "Supreme Command"
	SuperHeavy     Detach = "Super-heavy"
	SuperHeavyAux  Detach = "Super-heavy auxillary"
	Fortification  Detach = "Fortification network"
	Auxillary      Detach = "Auxillary support"
)

type Army struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	CampaignID  int    `json:"campaign_id"`
	Lore        string `json:"lore"`
	Name        string `json:"name"`
	SupplyLimit int    `json:"supply_limit"`
	SupplyUsed  int    `json:"supply_used"`
	AvailableCP int    `json:"available_cp"`
	Validated   bool   `json:"validated"`
}

type Advancement struct {
	ID       int    `json:"id"`
	AdvType  string `json:"adv_type"`
	Descr    string `json:"descr"`
	CrusCost int    `json:"crus_cost"`
}

type CampaignAchievements struct {
	AchievementID string `json:"achievement_id"`
	CampaignID    int    `json:"campaign_id"`
	UserID        int    `json:"user_id"`
}

type CampaignEvent struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Number  int    `json:"number"`
	Player1 int    `json:"player_one"`
	Player2 int    `json:"player_two"`
	Victor  int    `json:"victor"`
	Lore    string `json:"lore"`
}

type CampaignEvents struct {
	EventID    int `json:"event_id"`
	CampaignID int `json:"campaign_id"`
}

type Campaign struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Descr      string `json:"descr"`
	IsPublic   bool   `json:"is_public"`
	StartingPL int    `json:"starting_pl"`
}

type CampaignPlayers struct {
	UserID     int    `json:"user_id"`
	CampaignID int    `json:"campaign_id"`
	U_role     string `json:"u_role"`
}

type Detachment struct {
	ID             int    `json:"id"`
	ArmyID         int    `json:"army_id"`
	Fraction       string `json:"fraction"`
	CpCost         int    `json:"cp_cost"`
	DetachmentType Detach `json:"detach_type"`
	PowerLevel     int    `json:"power_level"`
}

type EventTallies struct {
	EventID int `json:"event_id"`
	TallyID int `json:"tally_id"`
	UnitID  int `json:"unit_id"`
}

type Tally struct {
	Id    int `json:"id"`
	Tally int `json:"tally"`
	Value int `json:"value"`
}

type UnitGrade struct {
	UnitID    int `json:"unit_id"`
	AdvanceID int `json:"advance_id"`
}

type Unit struct {
	ID            int      `json:"id"`
	DetachmentID  int      `json:"detach_id"`
	Name          string   `json:"name"`
	UnitType      string   `json:"unit_type"`
	Role          string   `json:"role"`
	Keywords      []string `json:"keywords"` // logically it is a string[]. For DB it is a string
	Experience    int      `json:"experience"`
	CrusadePoints int      `json:"crusade_points"`
	PowerLevel    int      `json:"power_level"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
