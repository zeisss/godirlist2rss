# godirlist2rss

Generates an RSS or ATOM feed with links to files. In case you ever want to be updated on new files in a folder that is also accessable by HTTP.

## Usage

My usage currently looks like this:

```bash
#!/bin/bash
godirlist2rss \
  --input-dir=/data/videos \
  --output-format=atom --output-file=listing.atom \
  --feed-title="$HOSTNAME Videos" --feed-author-name=Youtube-DL \
  --feed-public-url=https://$HOSTNAME/videos/feed.atom \
  --feed-filesbaseurl=https://$HOSTNAME/videos/
```

This is executed by a crontab entry regularly. This way the feel contains an entry for every file. If new files appear, new entries are generated. My RSS reader notifies me then. Each entry will link to the file (based on `--feed-filesbaseurl`).