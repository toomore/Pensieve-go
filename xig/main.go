package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

//const sample = `<script type="text/javascript">window._sharedData = {"country_code": "TW"};</script>`
var (
	filterV = regexp.MustCompile(`<script type="text/javascript">window._sharedData = (.+);</script>`)
	user    = flag.String("name", "", "ig id")
)

const ig = `https://www.instagram.com/%s/?hl=zh-tw`

func fetch(user string) *http.Response {
	log.Printf("Fetch data from %s`\n", user)
	resp, err := http.Get(fmt.Sprintf(ig, user))
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func filter1(html io.Reader) []byte {
	log.Println("Find json data ...")
	data, err := ioutil.ReadAll(html)
	if err != nil {
		log.Fatal(err)
	}
	if filterV.Match(data) {
		log.Println("Finded!!")
		for _, result := range filterV.FindAllSubmatch(data, -1) {
			//log.Printf("%d => %s", i, result[1])
			return result[1]
		}
	}
	return nil
}

func parseJSON(data []byte) {
	type puser struct {
		User struct {
			Username string `json:"username"`
			Follows  struct {
				Count int `json:"count"`
			}
		} `json:"user"`
	}

	type result struct {
		Code      string `json:"country_code"`
		EntryData struct {
			ProfilePage []puser `json:"ProfilePage"`
		} `json:"entry_data"`
	}

	var r result
	err := json.Unmarshal(data, &r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", r)

}

func dosomebad(user string) {
	resp := fetch(user)
	strJSON := filter1(resp.Body)
	parseJSON(strJSON)
}

func main() {
	//var v = regexp.MustCompile(r)
	//fmt.Println(v.MatchString(sample))
	//fmt.Println(v.FindAllStringSubmatch(sample, -1))
	flag.Parse()
	dosomebad(*user)
}
