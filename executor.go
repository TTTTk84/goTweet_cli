package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)



func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err .env loading ", err)
	}
}

func executor(s string) {
    Env_load()
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"),os.Getenv("ACCESS_TOKEN_SECRET"))

    args := strings.Split(s, " ")
    if len(args) > 2 {
        log.Fatal("args orver error")
        return
    }

    switch(args[0]){
    case "-post":
        // t, err := api.PostTweet(args[1],nil)
        // if err != nil {
        // 	log.Fatal("err :",err)
        // }
        // fmt.Println("success :",t.Text)
        fmt.Println("ok")
        return
    case "-search":
        searchUser(api, args[1])
        return
    default:
        log.Fatal("args fatal error")
        return
    }

}

func searchUser(api *anaconda.TwitterApi, word string) {
    t, err := api.GetUserSearch(word, nil)
    if err != nil {
        log.Fatal("err :",err)
    }
    for _,j := range t {
        fmt.Printf("Name: %s  ,ScreenName: %s  ,URL: %s \n",j.Name,  j.ScreenName, j.URL )
    }

}
