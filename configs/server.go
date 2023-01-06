package configs

import "os"

type ServerCfg struct {
	Address string
}

func NewServerCgf() *ServerCfg {
	return &ServerCfg{
		Address: os.Getenv("SERVER_ADDR"),
	}
}
