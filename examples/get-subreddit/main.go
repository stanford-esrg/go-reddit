package main

import (
	"context"
	"fmt"
	"log"
)

var ctx = context.Background()

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	// sr, _, err := reddit.DefaultClient().Subreddit.Get(ctx, "golang")
	c, err := NewClient(Credentials{}, WithUserAgent("test"))
	c, err := NewClient
	if err != nil {
		return
	}

	fmt.Printf("%s was created on %s and has %d subscribers.\n", sr.NamePrefixed, sr.Created.Local(), sr.Subscribers)
	return
}
