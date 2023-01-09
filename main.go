package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURLWithContext("http://rthk9.rthk.hk/rthk/news/rss/c_expressnews_clocal.xml", ctx)
	fmt.Println(feed.Title, feed.PublishedParsed)

	for _, item := range feed.Items {
		fmt.Println(">", item.Title, item.Description, "<")
	}
}
