package structure

import (
	"encoding/xml"
)

type Items struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"guid"`
	PublishDate string   `xml:"pubDate"`
}
