package scraper

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/davecgh/go-spew/spew"
    "net/http"
)

type AudioTextInfo struct {
    AudioAddress string
    Script       string
}

type Scraper struct {
    Url string
}

func NewScraper(url, outputDir string) *Scraper {
    return &Scraper{
        Url: url,
    }
}

func (s *Scraper) Analyze() (*AudioTextInfo, error) {
    res, err := http.Get(s.Url)
    if err != nil {
        logger.Errorf("get %s failed: %s", s.Url, err)
        return nil, err
    }
    logger.Pass("get url OK")

    defer res.Body.Close()
    if res.StatusCode != 200 {
        logger.Warnf("status code error, url: %s status: %s", s.Url, res.Status)
        return nil, err
    }

    // Load the HTML document
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        logger.Error("load body failed: ", err)
        return nil, err
    }
    logger.Pass("parse html OK")

    result := doc.Find("#jplayer_tc_yinpin")
    spew.Dump(result.Attr("href"))

    ati := AudioTextInfo{}
    ati.AudioAddress, _ = result.Attr("href")
    return &ati, fmt.Errorf("")
}
