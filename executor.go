package main

import (
	"fmt"
	"log"
	"net/url"
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

    switch(args[0]){
    case "-post":
        t, err := api.PostTweet(args[1],nil)
        if err != nil {
        	log.Fatal("err :",err)
        }
        fmt.Println("success tweet :",t.Text)
        return
    case "-search":
        searchUser(api, args)
        return
    default:
        log.Fatal("args fatal error")
        return
    }

}

func remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}


func searchUser(api *anaconda.TwitterApi, args []string) {
    searchTerm := args[len(args) - 1]
    args = remove(args, len(args)-1)
    args = remove(args, 0)

    v := url.Values{}
    if len(args) % 2 == 0 {
        for i:=0; i<len(args); i+=2{
            k := strings.Split(args[i], "")
            _k := strings.Join(k[1:], "")
            _v := args[len(args) - 1]
            v.Set(_k, _v)
        }
    } else {
        log.Fatal("err : args fatal error")
    }

    t, err := api.GetUserSearch(searchTerm , v)
    if err != nil {
        log.Fatal("err :",err)
    }
    for _, j := range t {
        fmt.Printf("Name: %s  ,ScreenName: %s  ,URL: %s \n",j.Name,  j.ScreenName, j.URL )
    }

}
