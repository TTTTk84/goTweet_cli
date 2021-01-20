package goTweet_cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)





func (c *Completer) Executor(s string) {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"),os.Getenv("ACCESS_TOKEN_SECRET"))

    args := strings.Split(s, " ")
    if len(args) < 2 {
        fmt.Fprintf(os.Stdout, "err: not enough args\n")
        return
    }

    switch(args[0]){
    case "-tweet":
        postTweet(api, args)
        return
    case "-search":
        getUserSearch(api, args)
        return
    default:
        fmt.Fprintf(os.Stdout, "err: args[0], not parameter \n")
        return
    }

}
