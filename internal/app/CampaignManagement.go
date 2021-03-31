package app
import "github.com/TheTh1rt33nth/Nexus-octo-engine/pkg/model"

func CreateCampaign (name string, descr string, public bool, startingPL int){
  //if /*check if campaign with such name already exists in DB*/{
    //fmt.Println("Campaign name is already occupied. Time to get more creative!")
 /* }
  else{
  campaign=model.Campaign{
    Name: name,
    Descr: descr,
    IsPublic: public,
    StartingPL: startingPL,
  }
  ownerID := 0 //Get current userID from session
  //askDB for a new campaignPlayer Entity, where userID is associated with
  //upforementioned campaign and has a U_role of "Admin"
  */
}
func AddPlayedBattle (name string, campaignID int, player1 int, player2 int, victor int, lore string)
{
  battlereport=model.CampaignEvent{
    Name: name,
    CampaignID: campaignID,
    Player1: player1,
    Player2: player2,
    Victor: victor,
  }
  if lore!=nil{
    battlereport.Lore=lore
  }
  //send to DB
}
func AddTally (battleID int, tallytype string, playerID int, unitID int, tallyvalue int) {
  tally=models.Tally{
    PlayerID: playerID,
    Tally: tallytype,
    Value: tallyvalue,
  }
  //send to DB, get its id
  tallyID:=0
  rosterID:=0 //get roster ID by checking user's army in this campaign by getting it through battleID
  tallyToEvent=models.EventTallies{
  EventID: battleID,
  UnitID: unitID,
  ArmyID: rosterID,
  TallyID: tallyID,
  }
  //send to DB, we're done with that
  //do recalculations for that unit
  //increase rosterID's roster RP by 1
}
