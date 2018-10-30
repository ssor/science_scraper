package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
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

	ati := AudioTextInfo{}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logger.Error("load body failed: ", err)
		return nil, err
	}
	logger.Pass("parse html OK")

	result := doc.Find("#jplayer_tc_yinpin")
	ati.AudioAddress, _ = result.Attr("href")

	var script string
	doc.Find("#tab_fanyi_con2 p").Each(func(i int, selection *goquery.Selection) {
		t := selection.Text()
		if len(t) > 0 {
			t = strings.TrimSpace(t)
			t = strings.Replace(t, "\n", " ", -1)
			for {
				if strings.Contains(t, "  ") {
					t = strings.Replace(t, "  ", " ", -1)
				} else {
					break
				}
			}
			script += "\r\n" + t
		}
	})
	ati.Script = script
	fmt.Println(ati.Script)

	return &ati, nil
}
