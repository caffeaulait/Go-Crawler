package fetcher

import (
	"bufio"
	"fmt"
	"go_crawler/crawler/config"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getEncoding(r *bufio.Reader) encoding.Encoding {
	//不能直接读,否则1024字节不能再读
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fail to determine encoding")
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	//<- rateLimiter
	log.Printf("fetching url %s", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", res.StatusCode)
	}
	reader := bufio.NewReader(res.Body)
	e := getEncoding(reader)
	ut8Reader := transform.NewReader(reader, e.NewDecoder())
	all, err := ioutil.ReadAll(ut8Reader)
	if err != nil {
		return nil, err
	} else {
		return all, nil
	}
}