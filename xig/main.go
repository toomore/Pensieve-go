package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	sizeR   = regexp.MustCompile(`/[a-z][0-9]+x[0-9]+`)
	filterV = regexp.MustCompile(`<script type="text/javascript">window._sharedData = (.+);</script>`)
	user    = flag.String("name", "", "IG username")
	getAll  = flag.Bool("all", false, "Get all data")
)

func fetch(user string) *http.Response {
	log.Printf("Fetch data from `%s`\n", user)
	resp, err := http.Get(fmt.Sprintf(`https://www.instagram.com/%s/?hl=zh-tw`, user))
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
	Owner struct {
		ID string `json:"id"`
	} `json:"owner"`
}

type media struct {
	Count    int    `json:"count"`
	Nodes    []node `json:"nodes"`
	PageInfo struct {
		EndCursor       string `json:"end_cursor"`
		HasNextPage     bool   `json:"has_next_page"`
		HasPreviousPage bool   `json:"has_previous_page"`
		StartCursor     string `json:"start_cursor"`
	} `json:"page_info"`
}

type profilepage struct {
	User struct {
		Biography          string `json:"biography"`
		FullName           string `json:"full_name"`
		HasRequestedViewer bool   `json:"has_requested_viewer"`
		ID                 string `json:"id"`
		IsPrivate          bool   `json:"is_private"`
		Media              media  `json:"media"`
		ProfilePicURL      string `json:"profile_pic_url"`
		ProfilePicURLHd    string `json:"profile_pic_url_hd"`
		Username           string `json:"username"`
		FollowedBy         struct {
			Count int `json:"count"`
		} `json:"followed_by"`
		Follows struct {
			Count int `json:"count"`
		} `json:"follows"`
	} `json:"user"`
}

// IGData struct
type IGData struct {
	Code      string `json:"country_code"`
	EntryData struct {
		ProfilePage []profilepage `json:"ProfilePage"`
	} `json:"entry_data"`
}

type queryData struct {
	Status string `json:"status"`
	Media  media  `json:"media"`
}

func parseIndexJSON(data []byte) *IGData {
	var r = &IGData{}
	err := json.Unmarshal(data, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func downloadNodeImage(node node, user string, wg *sync.WaitGroup) {
	defer wg.Done()
	path := sizeR.ReplaceAllString(node.DisplaySrc, "")
	url, err := url.Parse(path)
	if err != nil {
		log.Fatal(err)
	}
	data, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(fmt.Sprintf("./%s/img/%s%s", user, node.Code, strings.Replace(url.Path, "/", "_", -1)), body, 0644); err != nil {
		log.Fatal(err)
	} else {
		log.Println(fmt.Sprintf("Saved `%s`, `%s`", node.Code, node.DisplaySrc))
	}
}

func downloadAvatar(user string, path string) {
	path = sizeR.ReplaceAllString(path, "")
	url, err := url.Parse(path)
	if err != nil {
		log.Fatal(err)
	}
	data, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(fmt.Sprintf("./%s/avatar/%s", user, strings.Replace(url.Path, "/", "_", -1)), body, 0644); err != nil {
		log.Fatal(err)
	} else {
		log.Println(fmt.Sprintf("Saved `%s`, `%s`", user, path))
	}
}

func fetchAll(id string, username string, endCursor string, count int, cookies *http.Cookie) {
	v := url.Values{}
	v.Set("q", fmt.Sprintf(`ig_user(%s) { media.after(%s, %d) {
  count,
  nodes {
    caption,
    code,
    comments {
      count
    },
    comments_disabled,
    date,
    dimensions {
      height,
      width
    },
    display_src,
    id,
    is_video,
    likes {
      count
    },
    owner {
      id
    },
    thumbnail_src,
    video_views
  },
  page_info
}
 }`, id, endCursor, count))
	v.Set("ref", "users::show")

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.instagram.com/query/", strings.NewReader(v.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", fmt.Sprintf("https://www.instagram.com/%s/", username))
	req.Header.Set("x-csrftoken", cookies.Value)
	req.AddCookie(cookies)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data = &queryData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	var wg = &sync.WaitGroup{}
	wg.Add(len(data.Media.Nodes) * 2)
	for _, node := range data.Media.Nodes {
		go downloadNodeImage(node, username, wg)
		go saveNodeContent(node, username, wg)
	}
	wg.Wait()
}

func saveNodeContent(node node, user string, wg *sync.WaitGroup) {
	jsonStr, err := json.Marshal(node)
	if err != nil {
		log.Fatal(err)
	}
	basePath := fmt.Sprintf("./%s/content/%d_%s_%s.%%s", user, node.Date, node.Code, node.ID)
	if err := ioutil.WriteFile(fmt.Sprintf(basePath, "json"), jsonStr, 0644); err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(fmt.Sprintf(basePath, "txt"),
		[]byte(
			fmt.Sprintf("Code: %s\nCaption: %s\nDate: %s\nDisplaySrc: %s\nID: %s",
				node.Code, node.Caption, time.Unix(int64(node.Date), 0).Format(time.RFC3339), node.DisplaySrc, node.ID)),
		0644)
	log.Printf("Save content `%s`\n", node.Code)
	wg.Done()
}

func dosomebad(user string) {
	prepareBox(user)

	// Get nodes
	fetchData := fetch(user)
	defer fetchData.Body.Close()
	data := parseIndexJSON(filter1(fetchData.Body))
	var wg = &sync.WaitGroup{}
	UserData := data.EntryData.ProfilePage[0].User
	wg.Add(len(UserData.Media.Nodes) * 2)
	for _, node := range UserData.Media.Nodes {
		go downloadNodeImage(node, user, wg)
		go saveNodeContent(node, user, wg)
	}

	// Get avatar
	downloadAvatar(user, UserData.ProfilePicURLHd)

	wg.Wait()

	if *getAll {
		log.Println("Get all data!!!!")
		fetchAll(UserData.ID, UserData.Username, UserData.Media.PageInfo.EndCursor, UserData.Media.Count, fetchData.Cookies()[0])
	}

	fmt.Println("Username: ", UserData.Username)
	fmt.Println("Count: ", UserData.Media.Count)
}

func prepareBox(user string) {
	if err := os.Mkdir(fmt.Sprintf("./%s", user), 0777); err != nil {
		log.Println(err)
	}
	if err := os.Mkdir(fmt.Sprintf("./%s/img", user), 0777); err != nil {
		log.Println(err)
	}
	if err := os.Mkdir(fmt.Sprintf("./%s/avatar", user), 0777); err != nil {
		log.Println(err)
	}
	if err := os.Mkdir(fmt.Sprintf("./%s/content", user), 0777); err != nil {
		log.Println(err)
	}
}

func main() {
	flag.Parse()
	if len(*user) > 0 {
		dosomebad(*user)
	} else {
		flag.PrintDefaults()
	}
}
