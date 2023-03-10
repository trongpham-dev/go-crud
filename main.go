package main

import (
	"go-simple-api/component"
	"go-simple-api/middleware"
	"go-simple-api/modules/user/usertransport/ginuser"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DB_CONN_STR")
	secretKey := os.Getenv("SECRET_KEY")

	db, err := sqlx.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	appCtx := component.NewAppContext(db, secretKey)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))

	r.Run()

}
