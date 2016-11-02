package main

import (
	"crypto/md5"
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

func downloadNodeImage(node node, user string, wg *sync.WaitGroup) {
	defer wg.Done()

	path := sizeR.ReplaceAllString(node.DisplaySrc, "")
	url, err := url.Parse(path)
	if err != nil {
		log.Fatal(err)
	}
	err = downloadAndSave(path,
		fmt.Sprintf("./%s/img/%s%s", user, node.Code, strings.Replace(url.Path, "/", "_", -1)))

	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Saved `%s`, `%s`", node.Code, node.DisplaySrc))
}

func downloadAvatar(user string, path string, wg *sync.WaitGroup) {
	defer wg.Done()

	path = sizeR.ReplaceAllString(path, "")
	url, err := url.Parse(path)
	if err != nil {
		log.Fatal(err)
	}
	err = downloadAndSave(path,
		fmt.Sprintf("./%s/avatar/%s", user, strings.Replace(url.Path, "/", "_", -1)))

	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Saved avatar `%s`, `%s`", user, path))
}

func downloadAndSave(url string, path string) error {
	data, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	return ioutil.WriteFile(path, body, 0644)
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
	defer wg.Done()

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
}

func saveBiography(data profile, wg *sync.WaitGroup) {
	defer wg.Done()

	text := fmt.Sprintf("Username: %s\nFullName: %s\nID: %s\nIsPrivate: %t\nProfilePicURLHd: %s\nFollows: %d\nFollowsBy: %d\nBio: %s",
		data.Username, data.FullName, data.ID, data.IsPrivate, data.ProfilePicURLHd,
		data.Follows.Count, data.FollowedBy.Count, data.Biography)

	hex := md5.New()
	io.WriteString(hex, text)
	ioutil.WriteFile(fmt.Sprintf("./%s/profile/%s_%x.txt", data.Username, data.Username, hex.Sum(nil)), []byte(text), 0644)
	log.Printf("Save profile `%s` `%x`\n", data.Username, hex.Sum(nil))
}

func dosomebad(user string) {
	prepareBox(user)

	// Get nodes
	fetchData := fetch(user)
	defer fetchData.Body.Close()

	var data = &IGData{}
	if err := json.Unmarshal(filter1(fetchData.Body), &data); err != nil {
		log.Fatal(err)
	}

	if !data.EntryData.ProfilePage[0].User.IsPrivate {
		UserData := data.EntryData.ProfilePage[0].User

		var wg = &sync.WaitGroup{}
		wg.Add(len(UserData.Media.Nodes)*2 + 2)

		// Get avatar
		go downloadAvatar(user, UserData.ProfilePicURLHd, wg)
		go saveBiography(UserData, wg)

		for _, node := range UserData.Media.Nodes {
			go downloadNodeImage(node, user, wg)
			go saveNodeContent(node, user, wg)
		}

		wg.Wait()

		if *getAll {
			log.Println("Get all data!!!!")
			fetchAll(UserData.ID, UserData.Username, UserData.Media.PageInfo.EndCursor, UserData.Media.Count, fetchData.Cookies()[0])
		}

		fmt.Println("Username:", UserData.Username)
		fmt.Println("Count:", UserData.Media.Count)
	}
}

func prepareBox(user string) {
	for _, path := range [5]string{"", "/img", "/avatar", "/content", "/profile"} {
		if err := os.Mkdir(fmt.Sprintf("./%s%s", user, path), 0755); err != nil {
			log.Println(err)
		}
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
