package main

import (
	"fmt"
	lissajous "lissajous/lib"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/gif")
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("error on parsing form: %v", err)
			w.WriteHeader(400)
			fmt.Fprintf(w, "Error on parsing the request")
			return
		}

		params := lissajous.LissajousParams{}

		if cycles, err := strconv.Atoi(r.FormValue("cycles")); err == nil {
			params.Cycles = cycles
		}

		if res, err := strconv.ParseFloat(r.FormValue("res"), 64); err == nil {
			params.Res = res
		}

		if size, err := strconv.Atoi(r.FormValue("size")); err == nil {
			params.Size = size
		}

		if nframes, err := strconv.Atoi(r.FormValue("nframes")); err == nil {
			params.Nframes = nframes
		}

		if delay, err := strconv.Atoi(r.FormValue("delay")); err == nil {
			params.Delay = delay
		}

		fmt.Print(params)

		lissajous.Lissajous(w, params)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
