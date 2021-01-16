package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err loading ", err)
	}
}

func main() {
	Env_load()
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"),os.Getenv("ACCESS_TOKEN_SECRET"))

	text := "hello world"
	t, err := api.PostTweet(text,nil)
	if err != nil {
		log.Fatal("err :",err)
	}
	fmt.Println(t.Text)
}
