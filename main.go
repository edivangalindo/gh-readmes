package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	verbose := flag.Bool("v", true, "show verbose output")

	token := os.Getenv("GH_AUTH_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No repos detected. Hint: cat repos.txt | gh-readmes")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		owner := Before(line, "/")
		repo := After(line, "/")

		readme, resp, err := client.Repositories.GetReadme(ctx, owner, repo, nil)

		if err != nil {
			fmt.Println(err)

			if resp != nil {
				if resp.StatusCode == 404 {
					continue
				}
	
				if resp.Remaining == 0 {
					fmt.Printf("Rate limit exceeded, waitting for %+v minutes\n", resp.Rate.Reset.Sub(time.Now()).Minutes())
					time.Sleep(resp.Rate.Reset.Sub(time.Now()))
					continue
				}
			} else {
				fmt.Println("It looks like we are dealing with a network error... let's wait for 1 minute and follow the flow with other repositories...")
				time.Sleep(60 * time.Second)
				continue
			}
		}

		if resp.StatusCode != 200 {
			continue
		}

		fmt.Println("Remaining requests:", resp.Remaining, " - ", time.Now().Format("2006-01-02 15:04:05"), " - Github responded with status code ", resp.StatusCode)

		// Check if the readme is empty
		r, err := readme.GetContent()

		if err != nil {
			continue
		}
		if r == "" {
			continue
		}

		data, err := base64.StdEncoding.DecodeString(*readme.Content)

		if err != nil {
			fmt.Println(err)
			continue
		}

		SaveFile("readmes", owner+"-"+repo+".md", data)

		if *verbose {
			fmt.Println(string(data))
		}
	}
}
