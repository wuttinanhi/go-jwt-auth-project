package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuttinanhi/go-jwt-auth-project/auth"
	jwtservice "github.com/wuttinanhi/go-jwt-auth-project/jwt-service"
	"github.com/wuttinanhi/go-jwt-auth-project/middleware"
)

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func DefaultRoute(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "world"})
}

func LoginRoute(c *gin.Context) {
	var dto LoginDTO
	c.BindJSON(&dto)

	success := auth.GetAuthService().Login(dto.Username, dto.Password)

	if !success {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwtservice.GetJWTService().GenerateToken(&jwtservice.AuthJWT{UserId: dto.Username})

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func ProtectedRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"route":  "protected",
		"userId": c.GetString("userId"),
	})
}

func main() {
	// create gin router
	router := gin.Default()

	// apply recovery middleware
	router.Use(gin.Recovery())

	// default route
	router.GET("/", DefaultRoute)

	// login route
	router.POST("/login", LoginRoute)

	// protected router group
	protected := router.Group("/protected")
	protected.Use(middleware.JwtAuth())
	{
		protected.GET("/", ProtectedRoute)
	}

	// run server
	router.Run(":3000")
}
