package main

type room struct {
	// forwardは他のクライアントに’転送するためのメッセージを保持するチャネルです
	forward chan []byte
}
