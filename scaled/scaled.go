package scaled

import (
	"text/template"

	"github.com/embedfi/finance"
	pgs "github.com/lyft/protoc-gen-star"
)

type ScaledAmount struct{}

func (sv ScaledAmount) TypeName() string {
	return ".embedfi.finance.v1.ScaledAmount"
}

func (sv ScaledAmount) Extract(f pgs.Field) (rule interface{}, err error) {
	var rules *finance.Rules
	if _, err := f.Extension(finance.E_Rules, &rules); err != nil {
		return rule, err
	}

	if rules == nil {
		return nil, nil
	}

	return rules.ScaledAmount, nil
}

func (sv ScaledAmount) AddTemplates(lang string, tpl *template.Template) {

	switch lang {
	case "go":
		template.Must(tpl.New(sv.TypeName()).Parse(goTpl))
	default:
		panic("no template for " + lang)
	}
}
