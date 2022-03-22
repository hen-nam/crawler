package fetcher

import (
	"bufio"
	"crawler/crawler_distributed/config"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(time.Second / config.Qps)

// Fetch 获取内容
func Fetch(url string) ([]byte, error) {
	<-rateLimiter

	log.Printf("Fetching url %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("wrong status code: %d", resp.StatusCode)
		return nil, err
	}

	reader := bufio.NewReader(resp.Body)
	e := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// determineEncoding 获取字符编码
func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	contents, err := reader.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(contents, "")
	return e
}
