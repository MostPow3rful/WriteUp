package structure

import (
	"encoding/xml"
)

type Rss struct {
	XMLName xml.Name  `xml:"rss"`
	Channel []Channel `xml:"channel"`
}
