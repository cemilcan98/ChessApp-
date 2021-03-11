package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)
func main(){

	instance := echo.New()
	instance.GET("/game/:game",postGame)
	instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	if err:=instance.Start(":5001");err!=nil{
		log.Fatalf("Failed to service start.")
	}
}

func postGame(c echo.Context) error {
	game := c.Param("game")
	fen := c.QueryParam("fen")
	pgn := c.QueryParam("pgn")
	fmt.Printf("Game Number: %s, Fen:%s\n",game,fen)
	fmt.Printf("%s\n",pgn)

	return c.JSON(http.StatusOK,"ok")
}