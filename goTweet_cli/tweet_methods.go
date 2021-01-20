package goTweet_cli

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)


func remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}

func getUrlquery(args []string) (string, url.Values, bool){
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
				return searchTerm,v,false
		}

		return searchTerm,v,true
}

func getUserSearch(api *anaconda.TwitterApi, args []string) {
		searchTerm, v, found := getUrlquery(args)
		if !found {
			fmt.Fprintf(os.Stdout, "err: missing query\n")
			return
		}

    t, err := api.GetUserSearch(searchTerm , v)
    if err != nil {
        log.Fatal("err: ",err)
    }
    for _, j := range t {
        fmt.Printf("Name: %s  ,ScreenName: %s  ,URL: %s \n",j.Name,  j.ScreenName, j.URL )
    }

}

func postTweet(api *anaconda.TwitterApi, args []string) {
		searchTerm, v, found := getUrlquery(args)
		if !found {
			fmt.Fprintf(os.Stdout, "err: missing query\n")
			return
		}

		t, err := api.PostTweet(searchTerm, v)
		if err != nil {
			log.Fatal("err: ",err)
		}
		fmt.Println("success tweet: ",t.Text)
}
