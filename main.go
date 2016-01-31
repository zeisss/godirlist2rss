package main

import (
	"io/ioutil"
	"log"
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
)

type Walker struct {
	Path string
}

func (w Walker) Walk() <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)

		filepath.Walk(w.Path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			return nil
		})
	}()
	return result
}

func main() {
	pflag.Parse()

	now := time.Now()
	feed := feeds.Feed{
		Title: "My Files TODO",
		Link: &feeds.Link{
			Href: "http://moinz.de",
		},
		Description: "Description TODO",
		Author: &feeds.Author{
			"Stephan Zeissler", "stephan.zeissler@moinz.de",
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
			Id:    "uri:file:" + path,
			Title: info.Name(),
			Link: &feeds.Link{
				Href: *feedFilesBaseUrl + relPath,
			},
			Description: path,
			// Author:  &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
			Created: info.ModTime(),
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
