package request

import (
	"github.com/JesusKian/WriteUp/src/config"
	"github.com/JesusKian/WriteUp/src/structure"

	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWriteUps(_url string) {
	var (
		err      error          = nil
		rss      *structure.Rss = &structure.Rss{}
		response *http.Response = &http.Response{}
		writeUps []byte         = []byte{}
	)

	response, err = http.Get(_url)
	if err != nil {
		config.SetLog("E", fmt.Sprintf("request.GetWriteUps() -> Couldn't Get Response From %s", _url))
		config.SetLog("D", err.Error())
	}
	defer response.Body.Close()

	writeUps, err = ioutil.ReadAll(response.Body)
	if err != nil {
		config.SetLog("E", fmt.Sprint("request.GetWriteUps() -> Couldn't Read Response Body"))
		config.SetLog("D", err.Error())
	}

	xml.Unmarshal(writeUps, rss)
	fmt.Println(rss.Channel[0].Items[0].Title)
}
