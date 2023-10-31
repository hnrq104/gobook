/*Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request to
https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd
that, using this index, prints the URL and transcript of each comic that matches a search term
provided on the command line. */

package xqcd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type XQCDComic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	day        string
	URL        string
}

type XQCDIndex struct {
	TotalCount int
	Items      []*XQCDComic
}

func SearchXQCD(stripNumber int) (*XQCDComic, error) {
	path := fmt.Sprintf("https://xkcd.com/%d/info.0.json", stripNumber)

	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("server status not OK: %s", resp.Status)
	}

	var result XQCDComic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	result.URL = fmt.Sprintf("https://xkcd.com/%d", stripNumber)

	resp.Body.Close()
	return &result, nil
}

func GetIndex(number int) (*XQCDIndex, error) {
	comics := make([]*XQCDComic, 0)
	for i := 1; i <= number; i++ {
		if i == http.StatusNotFound {
			continue
		}

		result, err := SearchXQCD(i)
		if err != nil {
			break
		}
		comics = append(comics, result)
	}

	var ind XQCDIndex
	ind.TotalCount = len(comics)
	ind.Items = comics
	return &ind, nil
}

func GetTotal() (int, error) {
	const (
		LastComic = "https://xkcd.com/info.0.json"
	)
	resp, err := http.Get(LastComic)
	if err != nil {
		resp.Body.Close()
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return 0, fmt.Errorf("server status not OK: %s", resp.Status)
	}
	var number struct{ Num int }
	if err := json.NewDecoder(resp.Body).Decode(&number); err != nil {
		resp.Body.Close()
		return 0, err
	}
	resp.Body.Close()
	return number.Num, nil

}

func WriteIndex(out io.Writer, index *XQCDIndex) error {
	data, err := json.MarshalIndent(index, "", "\t")
	if err != nil {
		return fmt.Errorf("json.Marshal: %v", err)
	}
	_, err = out.Write(data)
	return err
}

func SearchInIndex(s string, index *XQCDIndex) []*XQCDComic {
	m := make(map[*XQCDComic]bool, 0)
	// I always will look into lowecasestuff
	lower := strings.ToLower(s)
	for _, comic := range index.Items {
		if strings.Contains(strings.ToLower(comic.Title), lower) ||
			strings.Contains(strings.ToLower(comic.Transcript), lower) ||
			strings.Contains(strings.ToLower(comic.Alt), lower) {
			m[comic] = true
		}
	}

	result := make([]*XQCDComic, 0, len(m))
	for comic := range m {
		result = append(result, comic)
	}

	return result
}
