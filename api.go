package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vartanbeno/go-reddit/reddit"
)

func GetGameScores() ([]interface{}, error) {

	resp, err := http.Get("http://gotto.io/cfb")
	if err != nil {
		return nil, err
	}

	var games []interface{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&games)

	if err != nil {
		return nil, err
	}

	return games, nil
}

func GetRedditFeeds() (map[string][]Subreddit, error) {

	var ctx = context.Background()

	credentials := reddit.Credentials{
		ID:       os.Getenv("REDDIT_ID"),
		Secret:   os.Getenv("REDDIT_KEY"),
		Username: os.Getenv("REDDIT_USER"),
		Password: os.Getenv("REDDIT_PASSWORD"),
	}

	httpClient := &http.Client{Timeout: time.Second * 30}
	client, err := reddit.NewClient(httpClient, &credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to create Reddit client: %v", err)
	}

	subscribedSubreddits, _, err := client.Subreddit.Subscribed(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscribed subreddits: %v", err)
	}

	subredditNames := []string{}
	for _, sub := range subscribedSubreddits.Subreddits {
		subredditNames = append(subredditNames, sub.Name)
	}
	opts := reddit.ListOptions{
		Limit: 5,
	}
	result := make(map[string][]Subreddit)
	for _, subredditName := range subredditNames {
		hotPosts, _, err := client.Subreddit.HotPosts(ctx, subredditName, &opts)
		if err != nil {
			log.Printf("Error retrieving hot posts for subreddit %s: %v", subredditName, err)
			continue
		}

		var subredditPosts []Subreddit
		for _, post := range hotPosts.Posts {
			subredditPosts = append(subredditPosts, Subreddit{
				Title:    post.Title,
				Score:    post.Score,
				ImageUrl: post.URL,
				PostUrl:  "https://reddit.com" + post.Permalink,
			})
		}
		result[subredditName] = subredditPosts
	}

	return result, nil
}

type Subreddit struct {
	Title    string
	Score    int
	ImageUrl string
	PostUrl  string
}
