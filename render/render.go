package render

/*
 * simple wrapper around golang templating mechanism
 * - users give file names to load
 * - we load+parse them
 * - we execute them with data given by user when user requests it
 * - user could possibly tell us to reload some files at runtime. be prepared for that too
 */

import (
	cfg "../configmgr"
	"io"
	"io/ioutil"
	"sync"
	"text/template"
)

var templates *template.Template
var templatesLock sync.RWMutex

func initialize() {
	templates = template.New("root")
	var funcs = make(template.FuncMap)
	funcs["date"] = date
	funcs["inc"] = inc
	templates.Funcs(funcs)
}

func init() {
	templatesLock.Lock()
	defer templatesLock.Unlock()

	if templates == nil {
		initialize()
	}
}

func parseFromFile(t *template.Template, fname string) (*template.Template, error) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return t.Parse(string(buf))
}

func parseSubFromFile(t *template.Template, tname, fname string) (*template.Template, error) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var nt *template.Template
	if t != nil {
		nt = t.New(tname)
	} else {
		nt = template.New(tname)
	}

	return nt.Parse(string(buf))
}

func Load(name, filename string) {
	templatesLock.Lock()
	defer templatesLock.Unlock()

	if templates == nil {
		initialize()
	}

	if tmpldir, _ := cfg.GetOption("render.templatedir"); tmpldir != "" {
		filename = tmpldir + "/" + filename
	}

	if name != "root" {
		template.Must(parseSubFromFile(templates, name, filename))
	} else {
		template.Must(parseFromFile(templates, filename))
	}
}

func Execute(w io.Writer, name string, data interface{}) {
	templatesLock.RLock()
	defer templatesLock.RUnlock()

	if templates == nil {
		panic("templates not yet initialized!!!")
	}

	if err := templates.ExecuteTemplate(w, name, data); err != nil {
		panic(err)
	}
}
