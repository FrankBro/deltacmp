package deltacmp

import (
	"fmt"
	"math"
)

type Mode int

const (
	ModeValue Mode = iota
	ModeDelta
	ModePercent
)

func (mode Mode) String() string {
	switch mode {
	case ModeValue:
		return "value"
	case ModeDelta:
		return "delta"
	case ModePercent:
		return "percent"
	default:
		panic("unknown mode")
	}
}

type Deltacmp struct {
	fields map[string]field
}

func difference(a, b float64) float64 {
	diff := math.Abs(a - b)
	avg := (a + b) / 2
	return diff / avg
}

func Compare(a, b *Deltacmp, modes map[string]Mode) map[string]string {
	diff := make(map[string]string)
	for name, fa := range a.fields {
		if mode, ok := modes[name]; ok {
			if fb, ok := b.fields[name]; ok {
				switch ta := fa.(type) {
				case *fieldBool:
					switch mode {
					case ModeValue:
						if tb, ok := fb.(*fieldBool); ok {
							if ta.value != tb.value {
								diff[name] = fmt.Sprintf("%s value was different. %v vs %v", name, ta.value, tb.value)
							}
						}
					case ModeDelta:
						panic("cannot do delta diff for boolean")
					case ModePercent:
						panic("cannot do percent diff for boolean")
					}
				case *fieldString:
					switch mode {
					case ModeValue:
						if tb, ok := fb.(*fieldString); ok {
							if ta.value != tb.value {
								diff[name] = fmt.Sprintf("%s value was different. %v vs %v", name, ta.value, tb.value)
							}
						}
					case ModeDelta:
						panic("cannot do a delta diff for a string")
					case ModePercent:
						panic("cannot do percent diff for boolean")
					}
				case *fieldInt:
					if tb, ok := fb.(*fieldInt); ok {
						switch mode {
						case ModeValue:
							if ta.value != tb.value {
								diff[name] = fmt.Sprintf("%s value was different. %v vs %v", name, ta.value, tb.value)
							}
						case ModeDelta:
							if ta.delta != tb.delta {
								diff[name] = fmt.Sprintf("%s delta was different. %v vs %v", name, ta.delta, tb.delta)
							}
						case ModePercent:
							a := float64(ta.value)
							b := float64(tb.value)
							if a > 0 || b > 0 {
								if d := difference(a, b); d > 0.01 {
									diff[name] = fmt.Sprintf("%s percent was different. %v vs %v (%.1f%%)", name, ta.value, tb.value, d*100)
								}
							}
						}
					}
				case *fieldUint:
					if tb, ok := fb.(*fieldUint); ok {
						switch mode {
						case ModeValue:
							if ta.value != tb.value {
								diff[name] = fmt.Sprintf("%s value was different. %v vs %v", name, ta.value, tb.value)
							}
						case ModeDelta:
							if ta.delta != tb.delta {
								diff[name] = fmt.Sprintf("%s delta was different. %v vs %v", name, ta.delta, tb.delta)
							}
						case ModePercent:
							a := float64(ta.value)
							b := float64(tb.value)
							if a > 0 || b > 0 {
								if d := difference(a, b); d > 0.01 {
									diff[name] = fmt.Sprintf("%s percent was different. %v vs %v (%.1f%%)", name, ta.value, tb.value, d*100)
								}
							}
						}
					}
				case *fieldFloat:
					if tb, ok := fb.(*fieldFloat); ok {
						switch mode {
						case ModeValue:
							if ta.value != tb.value {
								diff[name] = fmt.Sprintf("%s value was different. %v vs %v", name, ta.value, tb.value)
							}
						case ModeDelta:
							if ta.delta != tb.delta {
								diff[name] = fmt.Sprintf("%s delta was different. %v vs %v", name, ta.delta, tb.delta)
							}
						case ModePercent:
							a := ta.value
							b := tb.value
							if a > 0 || b > 0 {
								if d := difference(a, b); d > 0.01 {
									diff[name] = fmt.Sprintf("%s percent was different. %v vs %v (%.1f%%)", name, ta.value, tb.value, d*100)
								}
							}
						}
					}
				case *fieldComplex:
					if tb, ok := fb.(*fieldComplex); ok {
						switch mode {
						case ModeValue:
							if ta.value != tb.value {
								diff[name] = fmt.Sprintf("%s value was different. %v vs %v", name, ta.value, tb.value)
							}
						case ModeDelta:
							if ta.delta != tb.delta {
								diff[name] = fmt.Sprintf("%s delta was different. %v vs %v", name, ta.delta, tb.delta)
							}
						case ModePercent:
							panic("cannot do percent diff for complex")
						}
					}
				}
			}
		}
	}
	return diff
}
