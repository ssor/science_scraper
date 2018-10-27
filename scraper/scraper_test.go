package scraper

import (
    "github.com/stretchr/testify/assert"
    "strings"
    "testing"
)

var (
    htmlFile     = "sixty_second_science.html"
    audioAddress = "http://online1.tingclass.net/voaspe/2018/20181024sa_science.mp3"
    keyWords     = []string{
        "Asocial Octopuses Become Cuddly on MDMA",
        "When humans take the drug MDMA",
        "a good test subject for the question at hand",
        "The researchers set up a simple test",
        "versus the chamber with the toy",
        "After MDMA, all of the animals spent significantly more time",
        "as well as give clues about possible treatments for serotonin-related human conditions",
    }
)

func TestScraper_Run(t *testing.T) {
    url := "http://localhost:3000/" + htmlFile
    s := NewScraper(url, "data")
    ati, err := s.Analyze()
    assert.Nil(t, err)
    assert.NotNil(t, ati)
    assert.Equal(t, ati.AudioAddress, audioAddress)
    for _, kw := range keyWords {
        assert.True(t, strings.Contains(ati.Script, kw))
    }
}
