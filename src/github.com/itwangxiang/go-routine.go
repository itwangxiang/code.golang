package itwangxiang

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

func main() {
	doc, err := goquery.NewDocument("http://t66y.com/thread0806.php?fid=16")
	if err != nil{
		log.Fatal(err)
	}

	log.Print(doc)

	doc.Find(".content").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		fmt.Printf("Review %d: %s \n", i, s.Text())
	})
}
