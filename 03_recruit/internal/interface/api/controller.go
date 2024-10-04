package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shige1114/03_recruit/internal/app/interfaces"
	"github.com/shige1114/03_recruit/internal/infra/message"
	"github.com/shige1114/03_recruit/internal/interface/api/dto/mapper"
	"github.com/shige1114/03_recruit/internal/interface/api/dto/request"
	"github.com/shige1114/03_recruit/internal/interface/api/dto/response"
)

type Controller struct {
	service interfaces.ResultService
}

func NewRecruitController(g *gin.Engine, service interfaces.ResultService) {
	controller := &Controller{
		service: service,
	}

	// Configure and apply CORS middleware globally
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // フロントエンドからのアクセスを許可
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Apply the token verification middleware only to specific routes
	api := g.Group("/api/v1")
	api.Use(controller.VeryfyToken()) // Apply token verification middleware

	api.POST("/recruits", controller.CreateController)
	api.GET("/recruits", controller.GetController)
	api.PUT("/recruits", controller.UpdateController)
}
func (c *Controller) VeryfyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Request received")

		// Authorizationヘッダーからトークンを取得
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			ctx.Abort()
			return
		}

		// "Bearer "プレフィックスを削除してトークンを取得
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// if tokenString == authHeader {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		// 	ctx.Abort()
		// 	return
		// }

		// トークンをJSON形式で検証サーバーに送信
		verificationURL := "http://172.22.0.3:1000/verify_token" // 検証サーバーのURLに置き換えてください
		jsonData := map[string]string{"token": tokenString}
		jsonValue, err := json.Marshal(jsonData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JSON"})
			ctx.Abort()
			return
		}

		req, err := http.NewRequest("POST", verificationURL, bytes.NewBuffer(jsonValue))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			ctx.Abort()
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
			ctx.Abort()
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			ctx.Abort()
			return
		}

		// bodyの内容を出力

		var responseData response.AuthResponse
		if err := json.Unmarshal(body, &responseData); err != nil {
			fmt.Println("Unmarshal error:")             // エラーメッセージを表示
			fmt.Println("body (string):", string(body)) // bodyの内容を再度出力
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}

		// userIDをコンテキストに保存
		if userID := responseData.UserID; userID != "" {
			ctx.Set("UserID", userID)
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in response"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func (c *Controller) CreateController(g *gin.Context) {
	var createRecruit request.CreateRecruit
	fmt.Println(g)
	if err := g.BindJSON(&createRecruit); err != nil {
		fmt.Println(err)
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, exists := g.Get("UserID")
	if !exists {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "UserID not found"})
		return
	}
	createRecruit.UserID = userID.(string)
	fmt.Printf("Parsed CreateRecruit: %+v\n", createRecruit)
	// fmt.Println(">>")

	recruitCommand, err := request.ToCreateCommand(&createRecruit)
	if err != nil {
		log.Println(err)
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
	userID, exists := g.Get("UserID")
	if !exists {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "UserID not found"})
		return
	}
	// 型アサーションを行う
	userIDStr, ok := userID.(string)
	if !ok {
		g.JSON(http.StatusBadRequest, gin.H{"error": "UserID is not a string"})
		return
	}
	userId, err := uuid.Parse(userIDStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, map[string]string{"errors": fmt.Sprintf("Failed to parse %v", err)})
		return
	}
	recruits, err := c.service.FindByUserId(userId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, map[string]string{"errors": "Failed to get command"})
		return
	}
	for _, item := range recruits.Result {
		fmt.Printf("%+v\n", *item) // 各ポインタの内容を出力
	}

	log.Printf("Converted RecruitCommand: %+v\n", recruits)
	response := mapper.ToListResponse(recruits.Result)

	g.JSON(http.StatusOK, response)
}

func (c *Controller) UpdateController(g *gin.Context) {
	var request request.UpdateRequest
	if err := g.BindJSON(&request); err != nil { // ポインタを渡す
		g.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request"})
		return
	}

	updateCommand, err := request.ToUpdateRecruitCommand()
	if err != nil {
		log.Println(err)
		g.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse command"})
		return
	}

	commandResult, err := c.service.Update(updateCommand)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update recruit"})
		return
	}
	response := mapper.ToResponse(commandResult.Result)

	g.JSON(http.StatusOK, response)
}
