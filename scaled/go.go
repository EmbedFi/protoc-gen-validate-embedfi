package scaled

var goTpl = `
	// Scaled
	{{ $f := .Field}}{{ $r := index .CustomRules ".embedfi.finance.v1.ScaledAmount" }}
		{{ if $r }} 
		if financeErr := {{ accessor . }}.ValidateProto({{ $r }}); financeErr != nil {
			err := {{ err . "" }}
			err.reason = financeErr.Error()
			if !all { return err }
			errors = append(errors, err)
		}
		{{ end }}
`

/*
	{{ if $r.MaxScale }}
		if {{ accessor . }}.Scale > {{ $r.MaxScale.Value}} {
			err := {{ err . "scale must be less than or equal to " $r.MaxScale.Value }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.MinScale }}
		if {{ accessor . }}.Scale < {{ $r.MinScale.Value}} {
			err := {{ err . "scale must be greater than or equal to " $r.MinScale.Value }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.ExactScale }}
		if {{ accessor . }}.Scale < {{ $r.ExactScale.Value}} {
			err := {{ err . "scale must be exactly " $r.ExactScale.Value }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Lt }}
		if {{ accessor . }}.GTE(&finance.ScaledValue{Value: {{ $r.Lt.Val}}, Scale: {{ $r.Lt.Scale}}}) {
			err := {{ err . "must be less than " $r.Lt }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Gt }}
		if {{ accessor . }}.LTE({{ $r.Lt}}) {
			err := {{ err . "must be greater than " $r.Gt }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Lte }}
		if {{ accessor . }}.GT({{ $r.Lt}}) {
			err := {{ err . "must be less than or equal to " $r.Lte }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Gte }}
		if {{ accessor . }}.LT({{ $r.Lt}}) {
			err := {{ err . "must be greater than or equal to" $r.Gte }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
	// /Scaled
`*/
