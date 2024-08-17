package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shige1114/03_recruit/internal/app/service"
	"github.com/shige1114/03_recruit/internal/infra/datasource"
	"github.com/shige1114/03_recruit/internal/interface/api"
)

func main() {
	gormDB, err := datasource.Open()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	log.Println(gormDB)
	err = gormDB.AutoMigrate(&datasource.Recruit{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	recruitRepo := datasource.NewRepository(gormDB)

	recruitService := service.NewRecruitService(recruitRepo)

	g := gin.New()

	api.NewRecruitController(g, recruitService)

	if err := g.Run(":1020"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
