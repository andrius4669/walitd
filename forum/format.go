package forum

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
)

const (
	tagGreentext = iota
)

var tagMap = map[uint]struct{ start, end []byte }{
	tagGreentext: {[]byte("<span class=\"greentext\">"), []byte("</span>")},
}

// also sets up backlinks
func localValidatePost(t *threadPage, p *postContent, post uint32, thread *uint32) bool {
	var rpi int
	var ok bool
	rpi, ok = t.refMap[post]
	var rp *postContent
	if ok {
		if rpi == 0 {
			rp = &t.OP
		} else {
			rp = &t.Replies[rpi-1]
		}
		rp.References = append(rp.References, p.PostID)
		*thread = t.ID
		return true
	}
	return false
}

// check existence of cross-linking, ex: >>>/b/ >>>/pol/13548
func checkCrossPattern(b []byte, src int, end *int, board *string, post *uint32) bool {
	// shortest crosslink: >>>/a/ - 6 chars
	if src+6 > len(b) || b[src+1] != '>' || b[src+2] != '>' || b[src+3] != '/' {
		return false
	}

	src += 4
	idx := src
	for ; ; idx++ {
		if idx >= len(b) {
			return false
		}
		if b[idx] == '/' {
			if idx > src {
				break
			} else {
				return false
			}
		}
		if (b[idx] < 'a' || b[idx] > 'z') && (b[idx] < 'A' || b[idx] > 'Z') && (b[idx] < '0' || b[idx] > '9') {
			return false
		}
	}
	// can only break out with syntaxically correct board name
	*board = string(b[src:idx])
	idx++
	src = idx
	for ; ; idx++ {
		if idx >= len(b) || b[idx] < '0' || b[idx] > '9' {
			break
		}
	}
	*post = 0
	if idx > src {
		v, e := strconv.ParseUint(string(b[src:idx]), 10, 32)
		if e == nil {
			*post = uint32(v)
		}
	}
	*end = idx
	return true
}

func checkLinkPattern(b []byte, src int, end *int, post *uint32) bool {
	// shortest link: >>1 - 3 chars
	if src+3 > len(b) || b[src+1] != '>' {
		return false
	}

	src += 2
	idx := src
	for ; ; idx++ {
		if idx >= len(b) || b[idx] < '0' || b[idx] > '9' {
			break
		}
	}
	if idx > src {
		v, e := strconv.ParseUint(string(b[src:idx]), 10, 32)
		if e == nil {
			*post = uint32(v)
			*end = idx
			return true
		}
	}
	return false
}

// escapes and formats message
var (
	htmlQuot = []byte("&#34;") // shorter than "&quot;"
	htmlApos = []byte("&#39;") // shorter than "&apos;" and apos was not in HTML until HTML5
	htmlAmp  = []byte("&amp;")
	htmlLt   = []byte("&lt;")
	htmlGt   = []byte("&gt;")
	htmlBr   = []byte("<br />")
)

func formatPost(t *threadPage, p *postContent, db *sql.DB) {
	b := []byte(p.Message)
	var w bytes.Buffer
	src, last := 0, 0

	var tagList []uint

	for src < len(b) {
		c := b[src]
		var inc int
		var esc []byte
		switch c {
		case '"':
			esc = htmlQuot
			inc = 1
		case '\'':
			esc = htmlApos
			inc = 1
		case '&':
			esc = htmlAmp
			inc = 1
		case '<':
			esc = htmlLt
			inc = 1
		case '>':
			var board string
			var post uint32
			var end int
			if checkCrossPattern(b, src, &end, &board, &post) {
				if post != 0 {
					var pthread uint32
					if sqlValidatePost(db, board, post, &pthread) {
						// lookup successful
						esc = append(esc, []byte(fmt.Sprintf("<a class=\"crosslink\" href=\"/forum/%s/thread/%d#%d\">", board, pthread, post))...)
						esc = append(esc, htmlGt...)
						esc = append(esc, htmlGt...)
						esc = append(esc, htmlGt...)
						esc = append(esc, []byte(fmt.Sprintf("/%s/%d</a>", board, post))...)
					} else {
						// lookup fail, either board or post doesn't exist
						esc = append(esc, []byte("<span class=\"deadcrosslink\">")...)
						esc = append(esc, htmlGt...)
						esc = append(esc, htmlGt...)
						esc = append(esc, htmlGt...)
						esc = append(esc, []byte(fmt.Sprintf("/%s/%d</span>", board, post))...)
					}
				} else {
					if sqlValidateBoard(db, board) {
						// board ok
						esc = append(esc, []byte(fmt.Sprintf("<a class=\"crossboard\" href=\"/%s/\">", board))...)
						esc = append(esc, htmlGt...)
						esc = append(esc, htmlGt...)
						esc = append(esc, htmlGt...)
						esc = append(esc, []byte(fmt.Sprintf("/%s/</a>", board))...)
					} else {
						// board not ok
						esc = append(esc, []byte("<span class=\"deadcrossboard\">")...)
						esc = append(esc, htmlGt...)
						esc = append(esc, htmlGt...)
						esc = append(esc, htmlGt...)
						esc = append(esc, []byte(fmt.Sprintf("/%s/</span>", board))...)
					}
				}
				inc = end - src
			} else if checkLinkPattern(b, src, &end, &post) {
				pboard := t.Board
				var pthread uint32
				// CAUTION: local replies may be limited for board view
				// but we can skip this info in template if we find missing information
				// to be worse than no information at all
				var badlink = false
				if !localValidatePost(t, p, post, &pthread) {
					if !sqlValidatePost(db, pboard, post, &pthread) {
						esc = append(esc, []byte("<span class=\"deadlink\">")...)
						esc = append(esc, htmlGt...)
						esc = append(esc, htmlGt...)
						esc = append(esc, []byte(fmt.Sprintf("%d</span>", post))...)
						badlink = true
					}
				}
				if !badlink {
					esc = append(esc, []byte(fmt.Sprintf("<a class=\"postlink\" href=\"/forum/%s/thread/%d#%d\">", pboard, pthread, post))...)
					esc = append(esc, htmlGt...)
					esc = append(esc, htmlGt...)
					esc = append(esc, []byte(fmt.Sprintf("%d</a>", post))...)
				}
				inc = end - src
			} else if len(tagList) == 0 && (src == 0 || b[src-1] == '\n') {
				esc = append(tagMap[tagGreentext].start, htmlGt...)
				tagList = append(tagList, tagGreentext)
				inc = 1
			} else {
				esc = htmlGt
				inc = 1
			}
		case '\n':
			// bit fucked up way for doing this. TODO: do it in diferent way
			for i := int(len(tagList) - 1); i >= 0; i-- {
				if tagList[i] == tagGreentext {
					for j := int(len(tagList) - 1); j >= i; j-- {
						esc = append(esc, tagMap[tagList[j]].end...)
					}
					tagList = tagList[:i]
				}
			}
			esc = append(esc, htmlBr...)
			inc = 1
		case '\r':
			inc = 1 // just skip it
		default:
			src++
			continue
		}
		w.Write(b[last:src])
		w.Write(esc)
		src += inc
		last = src
	}
	w.Write(b[last:])
	for i := int(len(tagList) - 1); i >= 0; i-- {
		w.Write(tagMap[tagList[i]].end)
	}
	p.FMessage = w.String()
}
