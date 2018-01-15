package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
)

var appDir = path.Join(os.Getenv("HOME"), "Library", "Application Support", "##ServiceName##")

func main() {
	url := "##ServiceURL##"
	done := make(chan bool)
	go serve(done)
	askAuth(url)

	ok := <-done
	if ok {
		fmt.Println("OAuth token saved")
	}
}

func askAuth(url string) {
	exec.Command("open", url).Start()
}

func serve(done chan bool) {
	http.HandleFunc("/catch", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Thank you, GoSnatch can now access your spotify account.\nYou may close this window.\n")

		code := r.URL.Query()

		exchangeCode(code, done)
	})
	http.ListenAndServe(":3456", nil)
}

func exchangeCode(code string, done chan bool) {
	// proces will vary by servie

	// if token is got and saved
	done <- true
}
