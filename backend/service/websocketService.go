package service

import (
	"eSemaphore-backend/dto"
	"log"
	"github.com/gofiber/websocket/v2"
)

type ChatWebSocket struct {
	clients   map[string]map[*websocket.Conn]bool
	broadcast map[string]chan dto.ChatLine
}

func getChatWebSocket() *ChatWebSocket {
	return &ChatWebSocket{
		clients: make(map[string]map[*websocket.Conn]bool),
		broadcast: make(map[string]chan dto.ChatLine),
	}
}

func (c *ChatWebSocket) CreateChannel(chatId string){
	if _,exist := c.clients[chatId]; !exist{
		c.clients[chatId] = make(map[*websocket.Conn]bool)
	}

	if _,exist := c.broadcast[chatId]; !exist{
		c.broadcast[chatId] = make(chan dto.ChatLine)

		c.HandleMessages(chatId)
	}
}

func (c *ChatWebSocket) HandleMessages(chatId string) {
	for {
		msg,ok := <-c.broadcast[chatId]
		if !ok{
			break
		}

		log.Println(msg)
		for client := range c.clients[chatId] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("write error: %v", err)
				client.Close()
				delete(c.clients[chatId], client)
			}
		}
	}
}