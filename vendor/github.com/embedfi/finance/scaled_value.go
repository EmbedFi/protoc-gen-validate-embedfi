package finance

import (
	"database/sql/driver"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (sa *ScaledAmount) ToFloat() float64 {
	return float64(sa.Val) / math.Pow(10, float64(sa.Scale))
}

// Implements the Well Known Type interface for Double, allowing the validator
// to work
func (sa *ScaledAmount) GetValue() float64 {
	return sa.ToFloat()
}

func FromFloat(val float64, scale uint32) *ScaledAmount {
	v := int64(math.Round(val * math.Pow(10, float64(scale))))
	return &ScaledAmount{
		Val:   v,
		Scale: scale,
	}
}

func FromString(val string) (*ScaledAmount, error) {
	sa := &ScaledAmount{}
	return sa, sa.UnmarshalText([]byte(val))
}

func Zero() *ScaledAmount {
	return &ScaledAmount{}
}

func (sa *ScaledAmount) String() string {
	neg := false
	val := sa.Val
	if val < 0 {
		neg = true
		val = -1 * val
	}
	str := fmt.Sprintf("%d", val)
	split := len(str) - int(sa.Scale)

	scaledStr := ""
	if split == len(str) {
		scaledStr = str
	} else if split < 0 {
		scaledStr = "0." + strings.Repeat("0", -1*split) + str
	} else if split == 0 {
		scaledStr = "0." + str
	} else {
		scaledStr = str[0:split] + "." + str[split:]
	}

	if neg {
		scaledStr = "-" + scaledStr
	}
	return scaledStr
}

func (sa *ScaledAmount) MarshalText() ([]byte, error) {
	scaledStr := sa.String()
	return []byte(scaledStr), nil
}

func (sa *ScaledAmount) UnmarshalText(data []byte) error {
	vs := strings.TrimSpace(string(data))
	if vs == "" {
		sa.Val = 0
		sa.Scale = 0
		return nil
	}
	if strings.Contains(vs, "E") {
		if err := sa.unmarshalScientific(vs); err != nil {
			return fmt.Errorf("invalid scientific notation: %s: %w", vs, err)
		}
		return nil
	}

	idx := strings.Index(vs, ".")
	if idx < 0 {
		val, err := strconv.ParseInt(vs, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid number format: %s: %w", vs, err)
		}
		sa.Val = val
		sa.Scale = 0
		return nil
	}
	dotIdx := len(vs) - idx - 1

	val, err := strconv.ParseInt(strings.Replace(vs, ".", "", 1), 10, 64)
	if err != nil {
		return fmt.Errorf("invalid number format: %s: %w", vs, err)
	}
	sa.Val = val
	sa.Scale = uint32(dotIdx)

	return nil
}

func (sa *ScaledAmount) unmarshalScientific(vs string) error {
	parts := strings.Split(vs, "E")
	if len(parts) != 2 {
		return fmt.Errorf("expected one E")
	}

	base, err := FromString(parts[0])
	if err != nil {
		return err
	}
	exp, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return err
	}

	res := base.Pow10(exp)
	sa.Val = res.Val
	sa.Scale = res.Scale
	return nil

}

func (sa *ScaledAmount) UnmarshalJSON(data []byte) error {
	return sa.UnmarshalText(data)
}

func (sa *ScaledAmount) MarshalJSON() ([]byte, error) {
	v, err := sa.MarshalText()
	return []byte(v), err
}

func (sa *ScaledAmount) Value() (driver.Value, error) {
	val, err := sa.MarshalText()
	if err != nil {
		return nil, err
	}
	return []uint8(val), nil
}

func (sa *ScaledAmount) Scan(value interface{}) error {
	vsBytes, ok := value.([]uint8)
	if !ok {
		return fmt.Errorf("Invalid type %T", value)
	}
	return sa.UnmarshalText(vsBytes)
}

func addVal(a, b int64) int64 {
	// Taken from github.com/johncgriffin/overflow
	c := a + b
	if (c > a) == (b > 0) {
		return c
	}
	panic(fmt.Sprintf("Overflow Adding %d and %d as int64", a, b))
}

func mulVal(a, b int64) int64 {
	// Taken from github.com/johncgriffin/overflow
	if a == 0 || b == 0 {
		return 0
	}
	c := a * b
	if (c < 0) == ((a < 0) != (b < 0)) {
		if c/b == a {
			return c
		}
	}
	panic(fmt.Sprintf("Overflow Multiplying %d and %d as int64", a, b))
}

func addScale(a, b uint32) uint32 {
	// Not taken from the lib, needs verification
	c := a + b
	if c < a || c < b {
		panic(fmt.Sprintf("Overflow Adding Scale %d and %d as uint32", a, b))
	}
	return c
}

func matchScales(list ...*ScaledAmount) []*ScaledAmount {
	if len(list) < 2 {
		return list
	}
	maxScale := list[0].Scale
	for _, s := range list {
		if s.Scale > maxScale {
			maxScale = s.Scale
		}
	}

	out := make([]*ScaledAmount, len(list))
	for idx, unscaled := range list {
		out[idx] = unscaled.ToScale(maxScale)
	}
	return out
}

func (sa *ScaledAmount) Add(other *ScaledAmount) *ScaledAmount {
	vals := matchScales(sa, other)
	return &ScaledAmount{
		Val:   addVal(vals[0].Val, vals[1].Val),
		Scale: vals[0].Scale,
	}
}

func (sa *ScaledAmount) Sub(other *ScaledAmount) *ScaledAmount {
	return sa.Add(&ScaledAmount{
		Val:   other.Val * -1,
		Scale: other.Scale,
	})
}

func (sa *ScaledAmount) MultiplyRound(other *ScaledAmount, scale uint32) *ScaledAmount {
	return sa.multiplyRound(other, scale, RoundHalfUp)
}

// MultiplySimplest multiplies the numbers, the output will have as many decimal
// places are required to represent the number
func (sa *ScaledAmount) MultiplySimplest(sb *ScaledAmount) *ScaledAmount {
	sa = sa.SimplestDecimal()
	sb = sb.SimplestDecimal()
	n := &ScaledAmount{
		Val:   mulVal(sa.Val, sb.Val),
		Scale: addScale(sa.Scale, sb.Scale),
	}
	return n.SimplestDecimal()
}

// Divides A by B, limits the result to scale places to handle repeating decimals, after which the value is truncated / floored (e.g. 2/3 to 2 is 0.66, not 0.67)
func (sa *ScaledAmount) Divide(other *ScaledAmount, maxScale uint32) *ScaledAmount {
	// Here's 28 lines of code which say approximately ROUND(a / b, scale)
	vals := matchScales(sa, other)

	// Can discard the scale, as a/b == (a*1000)/(b*1000)
	num := vals[0].Val
	den := vals[1].Val
	nExp := uint32(0) // The Negative exponent (X10^-e)

	who := num / den // Floors, as is int
	rem := num % den

	// Val = who + (rem / den) X 10^exp
	// Val = (who * 10) + ((rem * 10) / den) X 10^(exp - 1)

	for rem != 0 && nExp < maxScale {
		// Add a Zero - 1 becomes 1.0
		who = who * 10
		rem = rem * 10
		nExp = nExp + 1

		// Bump in another digit of the remainder
		who += rem / den
		rem = rem % den
	}

	return &ScaledAmount{
		Val:   who,
		Scale: nExp,
	}
}

func (sa *ScaledAmount) MultiplyFloor(other *ScaledAmount, scale uint32) *ScaledAmount {
	return sa.multiplyRound(other, scale, RoundFloor)
}

func (sa *ScaledAmount) MultiplyCeil(other *ScaledAmount, scale uint32) *ScaledAmount {
	return sa.multiplyRound(other, scale, RoundCeil)
}

func (sa *ScaledAmount) multiplyRound(sb *ScaledAmount, scale uint32, method int) *ScaledAmount {
	sa = sa.SimplestDecimal()
	sb = sb.SimplestDecimal()
	n := &ScaledAmount{
		Val:   mulVal(sa.Val, sb.Val),
		Scale: addScale(sa.Scale, sb.Scale),
	}
	return n.round(scale, method)
}

// MultiplyInt is a special case of multiplication, by an integer, which doesn't need rounding
func (sa *ScaledAmount) MultiplyInt(by int64) *ScaledAmount {
	return &ScaledAmount{
		Val:   mulVal(sa.Val, by),
		Scale: sa.Scale,
	}
}

// Pow10 multiplies the value by 10 ^ exp
func (sa *ScaledAmount) Pow10(exp int64) *ScaledAmount {
	out := &ScaledAmount{
		Val:   sa.Val,
		Scale: sa.Scale,
	}

	if exp > 0 {
		for times := int64(0); times < exp; times++ {
			if out.Scale == 0 {
				out.Val = mulVal(out.Val, 10)
			} else {
				out.Scale = out.Scale - 1
			}
		}
	} else if exp < 0 {
		out.Scale = out.Scale + uint32(exp*-1)
	}
	return out
}

// SimplestDecimal removes trailing 0s on the right of the decimal place
func (sa *ScaledAmount) SimplestDecimal() *ScaledAmount {
	out := sa
	for out.Scale > 0 && out.Val%10 == 0 {
		out.Scale--
		out.Val = out.Val / 10
	}
	return sa
}

// LTZero: Greater Than Zero
func (sa *ScaledAmount) LTZero() bool {
	return sa.Val < 0
}

// LTEZero: Greater Than, or Equal To, Zero
func (sa *ScaledAmount) LTEZero() bool {
	return sa.Val <= 0
}

func (sa *ScaledAmount) LT(other *ScaledAmount) bool {
	m := matchScales(sa, other)
	return m[0].Val < m[1].Val
}

func (sa *ScaledAmount) LTE(other *ScaledAmount) bool {
	if sa.Equals(other) {
		return true
	}
	return sa.LT(other)
}

// GTZero: Greater Than Zero
func (sa *ScaledAmount) GTZero() bool {
	return sa.Val > 0
}

// GTEZero: Greater Than, or Equal To, Zero
func (sa *ScaledAmount) GTEZero() bool {
	return sa.Val >= 0
}

func (sa *ScaledAmount) GT(other *ScaledAmount) bool {
	m := matchScales(sa, other)
	return m[0].Val > m[1].Val
}

func (sa *ScaledAmount) GTE(other *ScaledAmount) bool {
	if sa.Equals(other) {
		return true
	}
	return sa.GT(other)
}

// Zero: Equal To Zero
func (sa *ScaledAmount) Zero() bool {
	return sa.Val == 0
}

// ToScale is just rounding, Half Up.
func (sa *ScaledAmount) ToScale(scale uint32) *ScaledAmount {
	return sa.round(scale, RoundHalfUp)
}

// Floor round to the given scale, to the nearest whole number, closer to 0
func (sa *ScaledAmount) Floor(scale uint32) *ScaledAmount {
	return sa.round(scale, RoundFloor)
}

// Ceil round to the given scale, to the nearest whole number, away from 0
func (sa *ScaledAmount) Ceil(scale uint32) *ScaledAmount {
	return sa.round(scale, RoundCeil)
}

const (
	RoundHalfUp int = iota
	RoundFloor
	RoundCeil
)

func (sa *ScaledAmount) round(scale uint32, method int) *ScaledAmount {
	if sa.Scale == scale {
		return sa
	}

	if sa.Scale < scale {
		val := sa.Val
		for i := uint32(0); i < scale-sa.Scale; i++ {
			val = mulVal(val, 10)
		}
		return (&ScaledAmount{
			Val:   val,
			Scale: scale,
		})
	}

	val := sa.Val

	switch method {
	case RoundHalfUp:
		// Floor Round to scale + 1, one extra place
		for i := uint32(0); i < sa.Scale-scale-1; i++ {
			val = val / 10
		}
		newVal := val / 10
		remainder := val % 10
		if val > 0 {
			if remainder >= 5 {
				newVal = newVal + 1
			}
		} else {
			if remainder <= -5 {
				newVal = newVal - 1
			}
		}
		val = newVal

	case RoundFloor:
		for i := uint32(0); i < sa.Scale-scale; i++ {
			// Go truncates, i.e. floor, ignore the remainder
			val = val / 10
		}

	case RoundCeil:
		isNeg := val < 0
		for i := uint32(0); i < sa.Scale-scale; i++ {
			newVal := val / 10
			remainder := val % 10
			if remainder != 0 {
				if isNeg {
					newVal -= 1
				} else {
					newVal += 1
				}
			}
			val = newVal
		}
	}

	return (&ScaledAmount{
		Val:   val,
		Scale: scale,
	})
}

func (sa *ScaledAmount) Equals(other *ScaledAmount) bool {
	vals := matchScales(sa, other)
	return vals[0].Val == vals[1].Val
}
