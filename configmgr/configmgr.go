package configmgr

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// braindead config storage
var config = make(map[string]string)

// initialize default values
func init() {
	config["server.listen"] = ":10666"
	// more default options to be added there
}

// load config from file
// being totally lazy there. but meh not expecting big config files
func LoadConfig(fname string) error {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	lines := bytes.Split(buf, []byte("\n"))
	var currentsection = ""
	for i := range lines {
		if len(lines[i]) > 0 && lines[i][len(lines[i])-1] == '\r' {
			// CRLF seq
			lines[i] = lines[i][:len(lines[i])-1]
		}
		// skip whitespace
		var start int
		for start = 0; start < len(lines[i]); start++ {
			if lines[i][start] != ' ' && lines[i][start] != '\t' {
				break
			}
		}
		if start > 0 {
			lines[i] = lines[i][start:]
		}
		if len(lines[i]) == 0 {
			// empty
			continue
		}
		if lines[i][0] == '#' || lines[i][0] == ';' || (len(lines[i]) >= 2 && (string(lines[i][:2]) == "//" || string(lines[i][:2]) == "--")) {
			// comment
			continue
		}
		if lines[i][0] == '[' {
			// section
			var end int
			var found bool = false
			for end = len(lines[i]) - 1; end > 0; end-- {
				if lines[i][end] == ']' {
					found = true
					break
				}
			}
			if !found {
				fmt.Fprintf(os.Stderr, "cfg: missing section ] ending: %s\n", lines[i])
				continue
			}
			end--
			for start = 1; start <= end; start++ {
				if lines[i][start] != ' ' && lines[i][start] != '\t' {
					break
				}
			}
			for ; end >= start; end-- {
				if lines[i][end] != ' ' && lines[i][end] != '\t' {
					break
				}
			}
			if start <= end {
				currentsection = string(lines[i][start : end+1])
			} else {
				currentsection = ""
			}
		} else {
			// valname
			for start = 0; start < len(lines[i]); start++ {
				if lines[i][start] == ' ' || lines[i][start] == '\t' || lines[i][start] == '=' || lines[i][start] == ':' {
					break
				}
			}
			var valname = string(lines[i][:start])
			// skip ws
			for ; start < len(lines[i]); start++ {
				if lines[i][start] != ' ' && lines[i][start] != '\t' {
					break
				}
			}
			if start < len(lines[i]) && (lines[i][start] == '=' || lines[i][start] == ':') {
				config[currentsection+"."+valname] = string(lines[i][start+1:])
			} else {
				if start >= len(lines[i]) {
					config[currentsection+"."+valname] = ""
				} else {
					opt := strings.ToLower(string(lines[i][start:]))
					switch opt {
					case "yes":
					case "on":
					case "true":
					case "1":
						config[currentsection+"."+valname] = ""
					case "no":
					case "off":
					case "false":
					case "0":
						delete(config, currentsection+"."+valname)
					default:
						fmt.Fprintf(os.Stderr, "cfg: unrecognised special setting: %s\n", lines[i])
					}
				}
			}
		}
	}
	return nil
}

func GetOption(name string) (string, bool) {
	s, k := config[name]
	return s, k
}

func GetListenHost() string {
	s, _ := GetOption("server.host")
	return s
}
