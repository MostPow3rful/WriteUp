package request

import (
	"github.com/JesusKian/WriteUp/src/config"
	"github.com/JesusKian/WriteUp/src/sql"
	"github.com/JesusKian/WriteUp/src/structure"

	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWriteUps(_url string) {
	var (
		rss      *structure.Rss = &structure.Rss{}
		response *http.Response = &http.Response{}
		writeUps []byte         = []byte{}
		err      error          = nil
	)

	response, err = http.Get(_url)
	if err != nil {
		config.SetLog("E", fmt.Sprintf("request.GetWriteUps() -> Couldn't Get Response From %s", _url))
		config.SetLog("D", err.Error())
	}
	defer response.Body.Close()

	writeUps, err = ioutil.ReadAll(response.Body)
	if err != nil {
		config.SetLog("E", "request.GetWriteUps() -> Couldn't Read Response Body")
		config.SetLog("D", err.Error())
	}

	xml.Unmarshal(writeUps, rss)

	for _, item := range rss.Channel[0].Items {
		sql.CheckWriteUp(
			item.Title,
			item.Link,
			item.PublishDate,
		)
	}
}
