package main

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/zeisss/godirlist2rss"
)

var flags godirlist2rss.Flags

func init() {
	pflag.StringVar(&flags.InputPath, "input-dir", ".", "The directory to scan for files")
	pflag.StringVar(&flags.OutputFile, "output-file", "listing.rss", "Where to write the generated feed?")
	pflag.StringVar(&flags.OutputFormat, "output-format", "rss", "What format to write? atom or rss")
	pflag.StringVar(&flags.BaseURL, "feed-filesbaseurl", "http://example.com/feed/", "Base url for the files")
	pflag.StringVar(&flags.AuthorName, "feed-author-name", "Anonymous", "The author name for the feed")
	pflag.StringVar(&flags.AuthorEmail, "feed-author-email", "anonymous@nowhere.local", "The author email for the feed")
	pflag.StringVar(&flags.PublishedURL, "feed-public-url", "http://example.com/feed/", "The url where the feed can be downloaded")
	pflag.StringVar(&flags.FeedTitle, "feed-title", "godirlist2rss File Listing", "Title for the feed document")
	pflag.StringVar(&flags.FeedDescription, "feed-description", "Listing of files via godirlist2rss", "Description for the feed document")
}

func main() {
	pflag.Parse()

	if err := flags.BuildFeed(); err != nil {
		log.Fatalf("Failed to build: %v", err)
	}
}
