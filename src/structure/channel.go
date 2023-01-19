package structure

import (
	"encoding/xml"
)

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []Items  `xml:"item"`
}
