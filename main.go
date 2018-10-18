package main

import (
	"flag"
	"time"
  "math/rand"
	"github.com/rs/zerolog/log"
  "gopkg.in/kyokomi/emoji.v1"
)

func main() {
	delay := flag.Duration("delay", 5*time.Second, "sleep this duration before responding")
	flag.Parse()
  messages := []string{
    "I am thirsty and require :beer:",
    "I am starving and require :pizza:",
  }
  rand.Seed(time.Now().Unix())
  for {
    log.Info().Msg((emoji.Sprint(messages[rand.Intn(len(messages))])))
    time.Sleep(*delay)
  }
}
