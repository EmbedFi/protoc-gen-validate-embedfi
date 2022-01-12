package scaled

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/embedfi/finance"
	pgs "github.com/lyft/protoc-gen-star"
	"google.golang.org/protobuf/proto"
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

	asBytes, err := proto.Marshal(rules.ScaledAmount)
	if err != nil {
		return nil, err
	}

	bParts := make([]string, len(asBytes))
	for idx, b := range asBytes {
		bParts[idx] = fmt.Sprintf("%#0x", b)
	}

	return fmt.Sprintf("[]byte{%s}", strings.Join(bParts, ",")), nil
}

func (sv ScaledAmount) AddTemplates(lang string, tpl *template.Template) {

	switch lang {
	case "go":
		template.Must(tpl.New(sv.TypeName()).Parse(goTpl))
	default:
		panic("no template for " + lang)
	}
}
