package fetch_test

import (
	"bytes"
	fetch "fetch/lib"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	t.Run("correct url", func(t *testing.T) {
		mockedResponse := "Hello, world!!"
		want := fmt.Sprintf("%s\n%s", "200 OK", mockedResponse)
		fakeServer := makeServer(mockedResponse, 200)
		defer fakeServer.Close()

		buffer := &bytes.Buffer{}
		err := fetch.Fetch(fakeServer.URL, buffer)
		res := buffer.String()
		if err != nil {
			t.Fatalf("expected %s, got an error %v", mockedResponse, err)
		}

		if res != want {
			t.Fatalf("expected %s, got %s", want, res)
		}
	})

	t.Run("whitout http prefix", func(t *testing.T) {
		mockedResponse := "Hello, world!!"
		want := fmt.Sprintf("%s\n%s", "200 OK", mockedResponse)
		fakeServer := makeServer(mockedResponse, 200)
		defer fakeServer.Close()

		buffer := &bytes.Buffer{}
		url := strings.ReplaceAll(fakeServer.URL, "http://", "")
		err := fetch.Fetch(url, buffer)
		res := buffer.String()
		if err != nil {
			t.Fatalf("expected %s, got an error %v", want, err)
		}

		if res != want {
			t.Fatalf("expected %s, got %s", want, res)
		}
	})
}

func makeServer(res string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(res))
	}))
}
