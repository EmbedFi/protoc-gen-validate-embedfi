package finance

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

// ValidateProto accepts a proto encoded message so that the validator
// extension can use it without importing a type
func (v *ScaledAmount) ValidateProto(raw []byte) error {
	rules := &ScaledAmountRules{}
	if err := proto.Unmarshal(raw, rules); err != nil {
		return err
	}
	return v.Validate(rules)
}

func (v *ScaledAmount) Validate(rules *ScaledAmountRules) error {

	if rules.MaxScale != nil {
		if v.Scale > rules.MaxScale.Value {
			return fmt.Errorf("scale must be less than or equal to %d", rules.MaxScale.Value)
		}
	}

	if rules.MinScale != nil {
		if v.Scale < rules.MinScale.Value {
			return fmt.Errorf("scale must be greater than or equal to %d", rules.MinScale.Value)
		}
	}

	if rules.ExactScale != nil {
		if v.Scale < rules.ExactScale.Value {
			return fmt.Errorf("scale must be exactly %d", rules.ExactScale.Value)
		}
	}

	if rules.Lt != nil {
		if v.GTE(rules.Lt) {
			return fmt.Errorf("must be less than %s", rules.Lt)
		}
	}

	if rules.Gt != nil {
		if v.LTE(rules.Lt) {
			return fmt.Errorf("must be greater than %s", rules.Gt)
		}
	}

	if rules.Lte != nil {
		if v.GT(rules.Lt) {
			return fmt.Errorf("must be less than or equal to %s", rules.Lte)
		}
	}

	if rules.Gte != nil {
		if v.LT(rules.Lt) {
			return fmt.Errorf("must be greater than or equal to %s", rules.Gte)
		}
	}
	return nil

}
