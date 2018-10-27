package main

import (
    "github.com/ssor/science_scraper/scraper"
    "github.com/ssor/zlog"
    "os"
)

var (
    normalOutputDir = "data"
)

func main() {
    err := os.MkdirAll(normalOutputDir, os.ModePerm)
    if err != nil {
        zlog.Warn("cannot create dir: ", err)
        return
    }

    s := scraper.NewScraper("", "data")
    ati, err := s.Analyze()
    if err != nil {
        zlog.Warn("scrape failed: ", err)
        return
    }
    zlog.Info("mp3 address ", ati)
    zlog.Info("script: ", ati.Script[:])
}
