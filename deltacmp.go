package deltacmp

import "fmt"

type Mode int

const (
	ModeValue Mode = iota
	ModeDelta
)

type Deltacmp struct {
	fields map[string]field
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
						}
					}
				}
			}
		}
	}
	return diff
}
