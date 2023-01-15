package controllers

import (
	"encoding/json"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"projectA/db"
	"projectA/middlewares"
	"projectA/models"
)

type GatewayController struct{}

type Event struct {
	Type string      `json:"t"`
	Data interface{} `json:"d"`
}

type Authorization struct {
	Token string `json:"token"`
}

type Authenticated struct {
	IsAuthenticated bool `json:"isAuthenticated"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func sendError(ws *websocket.Conn) {
	err := ws.WriteJSON(Authenticated{
		IsAuthenticated: false,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (controller GatewayController) Handler(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
	}

	var user models.User
	authMiddleware := middlewares.GetJWT()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			return
		}

		var auth Authorization

		if err := json.Unmarshal(message, &auth); err != nil {
			sendError(ws)
			continue
		}

		token, err := authMiddleware.ParseTokenString(auth.Token)
		if err != nil {
			sendError(ws)
			continue
		}

		userId := jwt.ExtractClaimsFromToken(token)

		db.GetDB().First(&user, userId["id"])

		err = ws.WriteJSON(Authenticated{
			IsAuthenticated: true,
		})
		if err != nil {
			sendError(ws)
			continue
		}

		break
	}

	pubsub := db.GetRedis().Subscribe(db.GetContext(), "ch1")
	channel := pubsub.Channel()

	go func() {
		defer ws.Close()
		defer pubsub.Close()
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				return
			}

			if string(message) == "ping" {
				err := ws.WriteMessage(1, []byte("pong"))
				if err != nil {
					return
				}
			}
		}
	}()

	go func() {
		for msg := range channel {
			var raw Event

			if err := json.Unmarshal([]byte(msg.Payload), &raw); err != nil {
				fmt.Println(err)
				continue
			}

			err = ws.WriteJSON(raw)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}()
}
