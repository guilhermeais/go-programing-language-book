package fetch

import (
	"fmt"
	"io"
	"net/http"
	url "net/url"
	"os"
	"strings"
	"time"

	uuidgen "github.com/nu7hatch/gouuid"
)

func Fetch(url string, out io.Writer) error {
	if !strings.HasPrefix(url, "http") {
		url = fmt.Sprintf("http://%s", url)
	}
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("fetch: error on http get %v", err)
	}
	fmt.Fprintln(out, resp.Status)
	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("fetch: error when reading the body %v", err)
	}
	return nil
}

func FetchWithChannel(rawUrl string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(rawUrl, "http") {
		rawUrl = fmt.Sprintf("http://%s", rawUrl)
	}
	resp, err := http.Get(rawUrl)
	if err != nil {
		ch <- fmt.Sprintf("fetch: error on http get %v", err)
	}
	defer resp.Body.Close()

	workDir, err := os.Getwd()
	if err != nil {
		ch <- fmt.Sprintf("fetch: error when geting current word dir %v", err)
		return
	}
	parseUrl, err := url.Parse(rawUrl)
	if err != nil {
		ch <- fmt.Sprintf("fetch: error when parsing url %v", err)
		return
	}
	uuid, err := uuidgen.NewV4()
	if err != nil {
		ch <- fmt.Sprintf("fetch: error when generating uuid %v", err)
		return
	}
	filepath := fmt.Sprintf("%s/websites/%s.txt", workDir, fmt.Sprintf("%s-%s-%d", parseUrl.Hostname(), uuid.String(), time.Now().Unix()))
	fmt.Println(filepath)

	file, err := os.Create(filepath)
	if err != nil {
		ch <- fmt.Sprintf("fetch: error when creating file %v", err)
		return
	}
	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("fetch: error when reading the body %v", err)
		return
	}

	resp.Body.Close()
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, rawUrl)
}
