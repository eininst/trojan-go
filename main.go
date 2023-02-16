package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/p4gefau1t/trojan-go/common"
	"github.com/p4gefau1t/trojan-go/log"

	_ "github.com/p4gefau1t/trojan-go/build"
)

func main() {
	flag.Parse()

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "网站建设中")
	})
	go func() {
		http.ListenAndServe(":80", nil)
	}()

	for {
		h, err := common.PopOptionHandler()
		if err != nil {
			log.Fatal("invalid options")
		}
		err = h.Handle()
		if err == nil {
			break
		}
	}
}
