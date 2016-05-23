package main

import (
	cfg "../configmgr"
	"fmt"
	"net/http"
	"os"
	str "strings"
	//"../hostutil"
	"../forum"
	"../main_rdr"
	"../news"
	"../poll"
	"../users"
	"../admin"
)

var configFile = "config.ini"

type handlerType struct{}

func (handlerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "" {
		r.URL.Path = "/"
	} else if r.URL.Path[0] != '/' {
		r.URL.Path = "/" + r.URL.Path
	}
	if i := str.IndexByte(r.URL.Path[1:], '/'); i >= 0 {
		i++
		switch r.URL.Path[1:i] {
		case "users": // users subsystem
			users.HandleRequest(w, r, i)
			return
		case "forum": // forum subsystem
			forum.HandleRequest(w, r, i)
			return
		case "news": // news subsystem
			news.HandleRequest(w, r, i)
			return
		case "poll": // poll subsystem
			poll.HandleRequest(w, r, i)
			return
		case "admin": // poll subsystem
			admin.HandleRequest(w, r, i)
			return
		}
	}
	// main page handler
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
			continue // literal '-'
		}
		if len(os.Args[i]) > 2 {
			// can't handle multiple options in one argument yet
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
		err := cfg.LoadConfig(configFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading config: %s\n", err)
		}
	}
	// load templates which modules may use
	forum.LoadTemplates()
	news.LoadTemplates()
	users.LoadTemplates()
	admin.LoadTemplates()
	// k..
	http.ListenAndServe(cfg.GetListenHost(), &handlerType{})
	// TODO(andrius) error handling
}
