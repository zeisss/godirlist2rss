package main

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/feeds"
	"github.com/spf13/pflag"
)

var (
	inputDir     = pflag.String("input-dir", ".", "The directory to scan for files")
	outputFile   = pflag.String("output-file", "listing.rss", "Where to write the generated feed?")
	outputFormat = pflag.String("output-format", "rss", "What format to write? atom or rss")

	feedFilesBaseUrl = pflag.String("feed-filesbaseurl", "http://example.com/feed/", "Base url for the files")
	feedAuthorName   = pflag.String("feed-author-name", "Anonymous", "The author name for the feed")
	feedAuthorEmail  = pflag.String("feed-author-email", "anonymous@nowhere.local", "The author email for the feed")
	feedPublishUrl   = pflag.String("feed-public-url", "http://example.com/feed/", "The url where the feed can be downloaded")
	feedTitle        = pflag.String("feed-title", "Gowalker File Listing", "Title for the feed document")
	feedDescription  = pflag.String("feed-description", "Listing of files via gowalker", "Description for the feed document")
)

func main() {
	pflag.Parse()

	now := time.Now()
	feed := feeds.Feed{
		Title: *feedTitle,
		Link: &feeds.Link{
			Href: *feedPublishUrl,
		},
		Description: *feedDescription,
		Author: &feeds.Author{
			*feedAuthorName, *feedAuthorEmail,
		},
		Created: now,
	}

	filepath.Walk(*inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(*inputDir, path)
		if err != nil {
			return err
		}

		feed.Items = append(feed.Items, &feeds.Item{
			Id:    "file:" + url.QueryEscape(path),
			Title: info.Name(),
			Link: &feeds.Link{
				Href: *feedFilesBaseUrl + relPath,
			},
			Description: path,
			Author:      &feeds.Author{*feedAuthorName, *feedAuthorEmail},
			Created:     info.ModTime(),
		})

		return nil
	})

	var content string
	var err error
	switch *outputFormat {
	case "rss":
		content, err = feed.ToRss()
	case "atom":
		content, err = feed.ToAtom()
	default:
		log.Fatalf("Unknown format: %s", *outputFormat)
	}

	if err != nil {
		log.Fatalf("Failed to generate feed: %#v", err)
	}

	err = ioutil.WriteFile(*outputFile, []byte(content), 0777)
	if err != nil {
		log.Fatalf("Failed to store feed: %#v", err)
	}
}
