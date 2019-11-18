package deltacmp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Basic struct {
	Bool       bool
	String     string
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint       uint
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Byte       byte
	Rune       rune
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
}

func TestBasic(t *testing.T) {
	modes := map[string]Mode{
		"Bool":       ModeValue,
		"String":     ModeValue,
		"Int":        ModeValue,
		"Int8":       ModeValue,
		"Int16":      ModeValue,
		"Int32":      ModeValue,
		"Int64":      ModeValue,
		"Uint":       ModeValue,
		"Uint8":      ModeValue,
		"Uint16":     ModeValue,
		"Uint32":     ModeValue,
		"Uint64":     ModeValue,
		"Byte":       ModeValue,
		"Rune":       ModeValue,
		"Float32":    ModeValue,
		"Float64":    ModeValue,
		"Complex64":  ModeValue,
		"Complex128": ModeValue,
	}
	basic := Basic{}
	deltacmp1 := Load(basic)
	deltacmp2 := Load(basic)
	result := Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, 0)
	basic = Basic{
		Bool:       true,
		String:     "not empty",
		Int:        1,
		Int8:       1,
		Int16:      1,
		Int32:      1,
		Int64:      1,
		Uint:       1,
		Uint8:      1,
		Uint16:     1,
		Uint32:     1,
		Uint64:     1,
		Byte:       1,
		Rune:       1,
		Float32:    1,
		Float64:    1,
		Complex64:  1,
		Complex128: 1,
	}
	deltacmp2.Update(basic)
	result = Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, len(modes))
}

type Embedded struct {
	In int
}

type Embed struct {
	Embedded
	Out int
}

func TestEmbed(t *testing.T) {
	modes := map[string]Mode{
		"Embedded.In": ModeValue,
		"Out":         ModeValue,
	}
	embed := Embed{}
	deltacmp1 := Load(embed)
	deltacmp2 := Load(embed)
	result := Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, 0)
	embed = Embed{
		Embedded: Embedded{
			In: 1,
		},
		Out: 1,
	}
	deltacmp2.Update(embed)
	result = Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, len(modes))
}

type Delta struct {
	Int int
}

func TestDelta(t *testing.T) {
	modes := map[string]Mode{
		"Int": ModeDelta,
	}
	delta1 := Delta{Int: 1}
	delta2 := Delta{Int: 10}
	deltacmp1 := Load(delta1)
	deltacmp2 := Load(delta2)
	result := Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, 0)
	delta1.Int++
	delta2.Int++
	deltacmp1.Update(delta1)
	deltacmp2.Update(delta2)
	result = Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, 0)
	delta1.Int += 2
	delta2.Int += 3
	deltacmp1.Update(delta1)
	deltacmp2.Update(delta2)
	result = Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, len(modes))
}

type Percent struct {
	Int int
}

func TestPercent(t *testing.T) {
	modes := map[string]Mode{
		"Int": ModePercent,
	}
	percent1 := Percent{Int: 1000}
	percent2 := Percent{Int: 1009}
	deltacmp1 := Load(percent1)
	deltacmp2 := Load(percent2)
	result := Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, 0)
	percent2.Int++
	deltacmp1.Update(percent1)
	deltacmp2.Update(percent2)
	result = Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, 0)
	percent2.Int++
	deltacmp1.Update(percent1)
	deltacmp2.Update(percent2)
	result = Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, len(modes))
}

func TestZeroAndSamePercent(t *testing.T) {
	modes := map[string]Mode{
		"Int": ModePercent,
	}
	percent1 := Percent{Int: 0}
	percent2 := Percent{Int: 0}
	deltacmp1 := Load(percent1)
	deltacmp2 := Load(percent2)
	result := Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, 0)
	percent1.Int++
	deltacmp1.Update(percent1)
	deltacmp2.Update(percent2)
	result = Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, 1)
	percent2.Int++
	deltacmp1.Update(percent1)
	deltacmp2.Update(percent2)
	result = Compare(deltacmp1, deltacmp2, modes)
	require.Len(t, result, 0)
}
