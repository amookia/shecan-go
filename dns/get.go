package dns

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

//Crawl dns from site
func GetDns() []string {
	m := make([]string,0)
	resp,err := http.Get("https://shecan.ir")
	if err != nil {
		log.Fatal(err)
	}
	//respbody,_ := ioutil.ReadAll(resp.Body)
	doc,_ := goquery.NewDocumentFromReader(resp.Body)
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		attr,_ := s.Attr("class")
		if attr == "shecan-dns-ips" {
			m = append(m, s.Text())
		}
	})
	return m
}
