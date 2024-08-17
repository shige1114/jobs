package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shige1114/03_recruit/internal/app/interfaces"
	"github.com/shige1114/03_recruit/internal/infra/message"
	"github.com/shige1114/03_recruit/internal/interface/api/dto/mapper"
	"github.com/shige1114/03_recruit/internal/interface/api/dto/request"
)

type Controller struct {
	service interfaces.ResultService
}

func NewRecruitController(g *gin.Engine, service interfaces.ResultService) {
	controller := &Controller{
		service: service,
	}

	g.POST("/api/v1/recruits", controller.CreateController)
	g.GET("/api/v1/recruits", controller.GetController)
}

func (c *Controller) CreateController(g *gin.Context) {
	var createRecruit request.CreateRecruit

	// fmt.Println(">")
	if err := g.BindJSON(&createRecruit); err != nil {
		g.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse"})
		return
	}
	// fmt.Printf("Parsed CreateRecruit: %+v\n", createRecruit)
	// fmt.Println(">>")
	recruitCommand, err := request.ToCreateCommand(&createRecruit)
	if err != nil {
		g.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parce command"})
		return
	}
	// fmt.Printf("Converted RecruitCommand: %+v\n", recruitCommand)
	steps := []message.SagaStep{
		&message.Step1{UserID: createRecruit.UserID},
		&message.Step2{CompnayId: createRecruit.CompanyID},
	}

	if err := message.RunSaga(steps); err != nil {
		g.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to saga"})
		return
	}

	// fmt.Println(">>>")
	if err := c.service.Create(recruitCommand); err != nil {
		g.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create command"})
		return
	}

	// fmt.Println(">>")
	g.JSON(http.StatusCreated, map[string]string{"message": "success"})
}

func (c *Controller) GetController(g *gin.Context) {
	userId, err := uuid.Parse(g.Query("userId"))
	if err != nil {
		g.JSON(http.StatusBadRequest, map[string]string{"errors": fmt.Sprintf("Failed to parse %v", err)})
		return
	}
	recruits, err := c.service.FindByUserId(userId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, map[string]string{"errors": "Failed to get command"})
		return
	}

	// fmt.Printf("Converted RecruitCommand: %+v\n", recruits)
	response := mapper.ToListResponse(recruits.Result)

	g.JSON(http.StatusOK, response)
}
