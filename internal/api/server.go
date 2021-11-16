package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheTh1rt33nth/Nexus-octo-engine/internal/app"
	"github.com/gin-gonic/gin"
)

type Server struct {
	http   *http.Server
	engine *app.Engine
}

func NewServer(addr string, engine *app.Engine) *Server {
	s := &Server{
		http: &http.Server{
			Addr: addr,
		},
		engine: engine,
	}
	s.http.Handler = s.setupRouter()

	return s
}

func (s *Server) Run() {
	errs := make(chan error, 1)

	defer func() {
		if err := s.http.Close(); err != nil {
			log.Fatal(fmt.Errorf("server stopped with error: %w", err))
		}
	}()

	go func() {
		log.Printf("server started")
		errs <- s.http.ListenAndServe()
	}()

	err := <-errs
	if err != nil {
		log.Fatal("server exited with error: %w", err)
	}
}

func (s *Server) setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/helloworld", s.helloworld)
	//	r.POST("/login", s.login)
	//	r.POST("/register", s.register)

	//	r.GET("/api/roster/{rosterID}", s.getRoster)
	//	r.PUT("/api/roster/{rosterID}", s.updateRoster)
	//	r.DELETE("/api/roster/{rosterID}", s.deleteRoster)

	//	r.GET("/api/campaigns", s.getActiveCampaigns)
	//	r.GET("/api/campaigns/archived", s.getArchivedCampaigns)
	//	r.POST("/api/campaign", s.startCampaign)
	//	r.GET("/api/campaign/{campaignID}", s.getCampaign)
	//	r.PUT("/api/campaign/{campaignID}", s.updateCampaign)
	//	r.DELETE("/api/campaign/{campaignID}", s.deleteCampaign)
	//	r.GET("/api/campaign/{campaignID}/rosters", s.getRostersInCampaign)
	//	r.POST("/api/campaign/{campaignID}/roster", s.addRoster)
	//	r.GET("/api/campaign/{campaignID}/battles", s.getBattlesInCampaign)
	//	r.POSTT("/api/campaign/{campaignID}/battle", s.addBattle)
	//	r.GET("/api/campaign/{campaignID}/players", s.getPlayersInCampaign)
	//	r.POST("/api/campaign/{campaignID}/players", s.addPlayer)
	//	r.DELETE("/api/campaign/{campaignID}/players/{userID}", s.deletePlayer)

	//	r.PUT("/api/battle/{batleID}", s.updateBattle)
	//	r.DELETE("/api/battle/{batleID}", s.deleteBattle)
	//	r.GET("/api/battle/{batleID}/story", s.getBattleLore)

	//	r.GET("/api/player/{userID}", s.getPlayer)
	//	r.PUT("/api/player/{userID}", s.updatePlayer)

	//	r.GET("/api/player/{userID}/participate", s.getCampaignsForPlayer)

	return r
}

func (s *Server) helloworld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
}
