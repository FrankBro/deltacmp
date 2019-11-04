package deltacmp

type field interface {
	isField()
}

type fieldBase struct {
	name string
}

func (f *fieldBase) isField() {}

type fieldBool struct {
	fieldBase
	value bool
}

type fieldString struct {
	fieldBase
	value string
}

type fieldInt struct {
	fieldBase
	value int64
	delta int64
}

type fieldUint struct {
	fieldBase
	value uint64
	delta uint64
}

type fieldFloat struct {
	fieldBase
	value float64
	delta float64
}

type fieldComplex struct {
	fieldBase
	value complex128
	delta complex128
}
