package collect

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/lackone/go-crawler/pkg/utils"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

type BaseFetch struct {
}

func (b *BaseFetch) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
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


