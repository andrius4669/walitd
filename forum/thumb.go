package forum

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	thumbBgOp    = "red"
	thumbBgReply = "#D6DAF0"
)

const (
	thumbMaxW = 128
	thumbMaxH = 128
)

const (
	thumbIMagick = iota
	thumbConvert
	thumbGMConvert
)

// extensions/mime types/aliases mapped to converters/aliases
var thumbConvMap = map[string]string{
	">image":     "convert/jpg",
	"image/gif":  ">>image",
	"image/jpeg": ">>image",
	"image/png":  ">>image",
	"image/bmp":  ">>image",
	"":           "", // default
}

func findConverter(ext, mimetype string) (ret string) {
	var ok bool
	ret, ok = thumbConvMap["."+ext]
	if !ok {
		ret, ok = thumbConvMap[mimetype]
		if !ok && mimetype != "" {
			var msub string
			if i := strings.IndexByte(mimetype, '/'); i != -1 {
				mimetype, msub = mimetype[:i], mimetype[i+1:]
			}
			ret, ok = thumbConvMap[mimetype+"/*"]
			if !ok && msub != "" {
				ret, ok = thumbConvMap["*/"+msub]
			}
			if !ok {
				ret, ok = thumbConvMap[""]
			}
		}
	}
	for len(ret) > 0 && ret[0] == '>' {
		var s string
		s, ok = thumbConvMap[ret[1:]]
		if !ok {
			fmt.Printf("warning: broken thumbConvMap chain: %s\n", ret)
		}
		ret = s
	}
	return
}

type thumbMethodType struct {
	deftype string
	f       func(source, destdir, dest, destext, bgcolor string) error
}

var thumbMethods = map[string]thumbMethodType{
	"convert":    {deftype: "jpg", f: makeConvertThumb},
	"gm-convert": {deftype: "jpg", f: makeGmConvertThumb},
}

func runConvertCmd(gm bool, source, destdir, dest, destext, bgcolor string) error {
	tmpfile := destdir + "/" + ".tmp." + dest + "." + destext
	dstfile := destdir + "/" + dest + "." + destext

	_ = os.MkdirAll(destdir, 777)

	var runfile string
	var args []string

	if !gm {
		runfile = "convert"
	} else {
		runfile = "gm"
		args = append(args, "convert")
	}

	var convsrc string
	//	if i := strings.LastIndexByte(source, '.'); i >= 0 {
	//		var saucetype string
	//		if strings.ToUpper(source[i+1:]) != "JPG" {
	//			saucetype = strings.ToUpper(source[i+1:])
	//		} else {
	//			saucetype = "JPEG"
	//		}
	//		convsrc = saucetype + ":" + source + "[0]"
	//	} else {
	// shouldn't happen
	convsrc = source + "[0]"
	//}

	args = append(args, convsrc, "-thumbnail", fmt.Sprintf("%dx%d", thumbMaxW, thumbMaxH))
	if bgcolor != "" {
		args = append(args, "-background", bgcolor, "-flatten")
	}
	args = append(args, "-auto-orient", tmpfile)

	cmd := exec.Command(runfile, args...)
	err := cmd.Run()
	if err != nil {
		os.Remove(tmpfile)
		return err
	}

	os.Rename(tmpfile, dstfile)

	return nil
}

func makeConvertThumb(source, destdir, dest, destext, bgcolor string) error {
	return runConvertCmd(false, source, destdir, dest, destext, bgcolor)
}

func makeGmConvertThumb(source, destdir, dest, destext, bgcolor string) error {
	return runConvertCmd(true, source, destdir, dest, destext, bgcolor)
}

func makeThumb(fullname, fname, board, ext, mimetype string, isop bool) (string, error) {
	var err error

	method := findConverter(ext, mimetype)
	if method == "" || method[0] == '/' {
		return method, nil
	}

	var format string
	if i := strings.IndexByte(method, '/'); i != -1 {
		method, format = method[:i], method[i+1:]
	}

	var bgcolor string
	if isop {
		bgcolor = thumbBgOp
	} else {
		bgcolor = thumbBgReply
	}

	m, ok := thumbMethods[method]
	if !ok {
		fmt.Printf("warning: method %s not found\n", method)
		return "", nil
	}

	err = m.f(fullname, serverThumbPathDir(board), fname, m.deftype, bgcolor)
	if err != nil {
		return "", err
	}
	return fname + "." + format, nil
}
