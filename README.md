# godirlist2rss
Puts a list of your files into an rss file

## Download

URL | Platform  | Arch
--- | --------- | ----
https://apps.moinz.de/fs/fs.php/artifacts/godirlist2rss/latest/godirlist2rss-darwin      | Darwin | Amd64
https://apps.moinz.de/fs/fs.php/artifacts/godirlist2rss/latest/godirlist2rss-linux-amd64 | Linux  | Amd64
https://apps.moinz.de/fs/fs.php/artifacts/godirlist2rss/latest/godirlist2rss-linux-arm6  | Linux  | Arm 6
https://apps.moinz.de/fs/fs.php/artifacts/godirlist2rss/latest/godirlist2rss-linux-arm7  | Linux  | Arm 7

The builds with a continous build pipeline provided by [wercker.com](http://wercker.com).

## Usage

My usage currently looks like this: 

```
#!/bin/bash
FSPATH=/feeds/$HOSTNAME-shows.atom

godirlist2rss \
  --input-dir=/data/videos \
  --output-format=atom --output-file=listing.atom \
  --feed-title="$HOSTNAME Videos" --feed-author-name=SABNzb \
  --feed-public-url=https://apps.moinz.de/fs/fs.php${FSPATH}

fs.bash push ${FSPATH} ./listing.atom public-read
```

This is executed by a crontab entry regularly. `fs.bash` pushes the generated Atom file to my webspace, where my reader picks it up. This way I get notifications when new files appear in my shows directory.