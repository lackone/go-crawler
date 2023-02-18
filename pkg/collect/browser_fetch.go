package collect

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/lackone/go-crawler/pkg/proxy"
	"github.com/lackone/go-crawler/pkg/utils"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"time"
)

type BrowserFetch struct {
	Timeout time.Duration
	Proxy   proxy.ProxyFunc
}

func (b *BrowserFetch) Get(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: b.Timeout,
	}
	if b.Proxy != nil {
		transport := http.DefaultTransport.(*http.Transport)
		transport.Proxy = b.Proxy
		client.Transport = transport
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.46")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("error status code %d", resp.StatusCode))
	}

	reader := bufio.NewReader(resp.Body)
	e := utils.GetEncode(reader)
	newReader := transform.NewReader(reader, e.NewDecoder())
	return io.ReadAll(newReader)
}
