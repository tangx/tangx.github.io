package qcdn

import (
	"encoding/xml"
	"os"
)

type SiteMap struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
	LastMod string   `xml:"lastmod,omitempty"`
}

func mustURLs(flag *Flag) []string {

	data, err := os.ReadFile(flag.SiteMap)
	if err != nil {
		panic(err)
	}

	sitemap := &SiteMap{}
	err = xml.Unmarshal(data, sitemap)
	if err != nil {
		panic(err)
	}

	urls := make([]string, len(sitemap.URLs))
	for i, u := range sitemap.URLs {
		urls[i] = u.Loc
	}

	return urls
}
