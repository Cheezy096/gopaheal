package gopaheal

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func constructUrl(page int, tags []string) string {
	tagsJoined := strings.Join(tags, "%20")
	return fmt.Sprintf("https://rule34.paheal.net/post/list/%s/%d", tagsJoined, page)
}

func getLastPage(url string) (int, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return 0, err
	}

	var x []string
	var countString string
	var countInt int

	doc.Find(".blockbody:contains(Last)").Find("a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "Last" {
			band, ok := s.Attr("href")
			if ok {
				countString = band
			} else {
				return
			}
		}
	})

	if len(countString) < 1 {
		return 0, fmt.Errorf("could not find last page link")
	}

	x = strings.Split(countString, "/")
	countInt, err = strconv.Atoi(x[len(x)-1])
	if err != nil {
		return 0, err
	}

	return countInt, nil
}

func getPosts(url string) ([]string, error) {
	var failed int
	var posts []string

	for failed < 100 {
		doc, err := goquery.NewDocument(url)
		if err != nil {
			return nil, err
		}

		doc.Find(".blockbody").Find("div").Find("div:contains(File)").Find("a").Each(func(i int, s *goquery.Selection) {
			if s.Text() == "File Only" {
				band, ok := s.Attr("href")
				if ok {
					posts = append(posts, band)
				} else {
					return
				}
			}
		})

		// Sometimes the body returned is blank, I found that waiting
		// for a few seconds helps this connection issue.
		// Probably a rate limit from the server.
		if len(posts) < 1 {
			failed += 1
			time.Sleep(2 * time.Second)
		} else {
			failed = 5
			return posts, nil
		}
	}

	return nil, fmt.Errorf("no posts found on this page")
}
