package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"projectA/utils"
)

var ctx = context.Background()
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Router(r *gin.Engine) {
	r.GET("/gateway", gateway)
}

func gateway(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
	}

	pubsub := utils.Redis.Subscribe(ctx, "ch1")
	defer pubsub.Close()
	defer ws.Close()

	channel := pubsub.Channel()

	for msg := range channel {
		var raw map[string]interface{}

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
}
