package model

import "github.com/gorilla/websocket"

type Message struct {
	UserID  int    `json:"user_id"`
	DisId   int    `json:"dis_id"`
	Content string `json:"content"`
}

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}
