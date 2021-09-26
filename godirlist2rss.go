package godirlist2rss

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/feeds"
)

type BuildFlags struct {
	FeedTitle               string
	FeedDescription         string
	PublishedURL            string // url where the feed will be availble
	AuthorName, AuthorEmail string

	BaseURL string

	InputPath string

	OutputFile   string
	OutputFormat string
	OutputMode   os.FileMode
}

func (flags BuildFlags) BuildFeed() error {
	now := time.Now()
	feed := feeds.Feed{
		Title: flags.FeedTitle,
		Link: &feeds.Link{
			Href: flags.PublishedURL,
		},
		Description: flags.FeedDescription,
		Author: &feeds.Author{
			Name:  flags.AuthorName,
			Email: flags.AuthorEmail,
		},
		Created: now,
	}

	err := filepath.Walk(flags.InputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(flags.InputPath, path)
		if err != nil {
			return err
		}

		feed.Items = append(feed.Items, &feeds.Item{
			Id:    "file:" + url.QueryEscape(path),
			Title: info.Name(),
			Link: &feeds.Link{
				Href: flags.BaseURL + relPath,
			},
			Description: path,
			Author: &feeds.Author{
				Name:  flags.AuthorName,
				Email: flags.AuthorEmail,
			},
			Created: info.ModTime(),
		})

		return nil
	})
	if err != nil {
		return fmt.Errorf("scanning files failed: %w", err)
	}

	var content string
	switch flags.OutputFormat {
	case "rss":
		content, err = feed.ToRss()
	case "atom":
		content, err = feed.ToAtom()
	default:
		return fmt.Errorf("unsupported output format: %s", flags.OutputFormat)
	}

	if err != nil {
		return fmt.Errorf("failed to generate feed: %w", err)
	}

	mode := flags.OutputMode
	if mode == 0 {
		mode = 0666
	}
	err = os.WriteFile(flags.OutputFile, []byte(content), mode)
	if err != nil {
		return fmt.Errorf("writing output file failed: %w", err)
	}
	return nil
}
