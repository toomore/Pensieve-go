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
	log.Printf("Fetch data from `%s`\n", user)
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

type node struct {
	Caption          string `json:"caption"`
	Code             string `json:"code"`
	CommentsDisabled bool   `json:"comments_disabled"`
	Date             int    `json:"date"`
	DisplaySrc       string `json:"display_src"`
	ID               string `json:"id"`
	IsVideo          bool   `json:"is_video"`
	ThumbnailSrc     string `json:"thumbnail_src"`
	Comments         struct {
		Count int `json:"Count"`
	} `json:"comments"`
	Dimensions struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"dimensions"`
	Likes struct {
		Count int `json:"Count"`
	} `json:"likes"`
}

type puser struct {
	User struct {
		Biography          string `json:"biography"`
		FullName           string `json:"full_name"`
		HasRequestedViewer bool   `json:"has_requested_viewer"`
		ID                 string `json:"id"`
		IsPrivate          bool   `json:"is_private"`
		ProfilePicURL      string `json:"profile_pic_url"`
		ProfilePicURLHd    string `json:"profile_pic_url_hd"`
		Username           string `json:"username"`
		FollowedBy         struct {
			Count int `json:"count"`
		} `json:"followed_by"`
		Follows struct {
			Count int `json:"count"`
		} `json:"follows"`
		Media struct {
			Count    int `json:"count"`
			PageInfo struct {
				EndCursor       string `json:"end_cursor"`
				HasNextPage     bool   `json:"has_next_page"`
				HasPreviousPage bool   `json:"has_previous_page"`
				StartCursor     string `json:"start_cursor"`
			} `json:"page_info"`
			Nodes []node `json:"nodes"`
		} `json:"media"`
	} `json:"user"`
}

type result struct {
	Code      string `json:"country_code"`
	EntryData struct {
		ProfilePage []puser `json:"ProfilePage"`
	} `json:"entry_data"`
}

func parseJSON(data []byte) {

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
