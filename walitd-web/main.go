package main

import (
	"../configmgr"
	"fmt"
	"net/http"
	"os"
	"strings"
	//"../hostutil"
	"../forum_rdr"
	"../main_rdr"
)

var configFile = "config.ini"

type HandlerType struct{}

func (HandlerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "" {
		r.URL.Path = "/"
	}
	path := r.URL.Path[1:]
	i := strings.IndexByte(r.URL.Path, '/')
	if i >= 0 {
		switch path[:i] {
		case "f":
			forum_rdr.HandleRequest(w, r, i)
			return
		}
	}

	main_rdr.HandleRequest(w, r)
}

func main() {
	// option processing
	for i := 1; i < len(os.Args); i++ {
		if len(os.Args) > 0 && os.Args[i][0] != '-' {
			fmt.Fprintf(os.Stderr, "unprocessed non-option argument: %s\n", os.Args[i])
			return
		}
		if len(os.Args[i]) < 2 {
			continue
		}
		if len(os.Args[i]) > 2 {
			fmt.Fprintf(os.Stderr, "please split options: %s\n", os.Args[i])
			return
		}
		switch os.Args[i][1] {
		case 'c':
			i++
			if i < len(os.Args) {
				configFile = os.Args[i]
			} else {
				fmt.Fprintf(os.Stderr, "no config file specified for -c option\n")
				return
			}
		case '-':
			// stop processing rest of arguments
			break
		default:
			fmt.Fprintf(os.Stderr, "unrecognised option: %s\n", os.Args[i])
			return
		}
	}
	if configFile != "" {
		// load config
		err := configmgr.LoadConfig(configFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading config: %s\n", err)
		}
	}
	// k..
	http.ListenAndServe(configmgr.GetListenHost(), &HandlerType{})
}
