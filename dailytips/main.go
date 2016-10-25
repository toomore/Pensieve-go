package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		shareFilms = [4]string{
			"20150920 大阪 京都 東京 [films]",
			"20160131 北海道 小樽 千葉 求名 銚子 東京",
			"20160521 橫濱 東京 二日往返",
			"20160924 廣島 福岡",
		}
		now = time.Now()
	)
	fmt.Printf("今天 %d %s, 是 %d 年的第 %d 天\n", now.Day(), now.Month(), now.Year(), now.YearDay())
	fmt.Printf("建議分享的是 \"%s\" 的主題\n", shareFilms[now.YearDay()%len(shareFilms)])
}
