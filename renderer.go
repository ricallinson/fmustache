package fmustache

import(
	"github.com/hoisie/mustache"
)

type Fmustache struct {
    cache map[string]*mustache.Template
}

func Make() (*Fmustache) {
	this := &Fmustache{}
	this.cache = map[string]*mustache.Template{}
	return this
}

func (this *Fmustache) Render(f string, o ...interface{}) (string, error) {

	t, ok := this.cache[f]

    if ok == false {
        tmpl, err := mustache.ParseFile(f)
        if err != nil {
            return "", err
        }
        this.cache[f] = tmpl
        t = this.cache[f]
    }

    return t.Render(o...), nil
}