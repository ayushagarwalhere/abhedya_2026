package main

import (
	"abhedya_2026/internal/configs"
	"abhedya_2026/internal/server"
)

func init() {
	configs.LoadEnv()
}

func main() {
	server.Run()
}
