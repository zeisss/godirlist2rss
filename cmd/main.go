package main

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/zeisss/godirlist2rss"
)

var (
	inputDir     = pflag.String("input-dir", ".", "The directory to scan for files")
	outputFile   = pflag.String("output-file", "listing.rss", "Where to write the generated feed?")
	outputFormat = pflag.String("output-format", "rss", "What format to write? atom or rss")

	feedFilesBaseUrl = pflag.String("feed-filesbaseurl", "http://example.com/feed/", "Base url for the files")
	feedAuthorName   = pflag.String("feed-author-name", "Anonymous", "The author name for the feed")
	feedAuthorEmail  = pflag.String("feed-author-email", "anonymous@nowhere.local", "The author email for the feed")
	feedPublishUrl   = pflag.String("feed-public-url", "http://example.com/feed/", "The url where the feed can be downloaded")
	feedTitle        = pflag.String("feed-title", "godirlist2rss File Listing", "Title for the feed document")
	feedDescription  = pflag.String("feed-description", "Listing of files via godirlist2rss", "Description for the feed document")
)

func main() {
	pflag.Parse()

	flags := godirlist2rss.BuildFlags{
		FeedTitle:       *feedTitle,
		FeedDescription: *feedDescription,
		PublishedURL:    *feedPublishUrl,
		AuthorName:      *feedAuthorName,
		AuthorEmail:     *feedAuthorEmail,

		BaseURL: *feedFilesBaseUrl,

		InputPath: *inputDir,

		OutputFile:   *outputFile,
		OutputFormat: *outputFormat,
	}

	if err := flags.BuildFeed(); err != nil {
		log.Fatalf("Failed to build: %v", err)
	}
}
