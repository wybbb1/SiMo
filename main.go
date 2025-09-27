package main

import (
	"github.com/wybbb1/SiMo/internal/log"
)


func init() {
	log.SetupLogger()
}

func main() {
	defer log.Sync()
}