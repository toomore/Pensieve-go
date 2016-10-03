package main

import (
	"log"
	"net/url"
	"regexp"
)

var mailPath = regexp.MustCompile("^/mail(/.)?")
var pagePath = regexp.MustCompile("^/page/(?P<id>[a-zA-Z]+)/(?P<page>[0-9]+)")
var indexPath = regexp.MustCompile("^(/)?$")

func parse(web string) {
	log.Println("Start Parse url:", web)
	if p, err := url.Parse(web); err == nil {
		log.Printf("Path: [%s], RawPath: [%s], Host: [%s]\n", p.Path, p.RawPath, p.Host)
		switch {
		case mailPath.MatchString(p.Path):
			log.Println("In mail", p.Path)
		case indexPath.MatchString(p.Path):
			log.Println("In index", p.Path)
		case pagePath.MatchString(p.Path):
			match := pagePath.FindStringSubmatch(p.Path)
			log.Println("In page", p.Path)
			log.Println(match[1], match[2])
		default:
			log.Println("No case", p.Path)
		}
	} else {
		log.Println(err)
	}
}

func replace() {
	str := "/page/pppqqqq/5"
	strSlide := []rune(str)
	index := pagePath.FindStringSubmatchIndex(str)
	for i := range index {
		if i%2 == 0 {
			log.Printf("%s\n", string(strSlide[index[i]:index[i+1]]))
		}
	}
}

func main() {
	parse("https://www.google.com:8080")
	parse("https://www.google.com:8080/")
	parse("https://www.google.com:8080/mail")
	parse("https://www.google.com:8080/mail/")
	parse("https://www.google.com:8080/mail/one")
	parse("https://www.google.com:8080/mail%2fone")
	parse("https://www.google.com:8080/page/pppqqqq/5")
	parse("https://www.google.com:8080/nonono")
	replace()
}
