package deltacmp

import (
	"reflect"
)

func (cmp *Deltacmp) Update(obj interface{}) {
	forEachField(reflect.ValueOf(obj), "", func(field reflect.Value, name string) {
		if cmpField, ok := cmp.fields[name]; ok {
			switch realField := cmpField.(type) {
			case *fieldBool:
				realField.value = field.Bool()
			case *fieldString:
				realField.value = field.String()
			case *fieldInt:
				value := field.Int()
				realField.delta = value - realField.value
				realField.value = value
			case *fieldUint:
				value := field.Uint()
				realField.delta = value - realField.value
				realField.value = value
			case *fieldFloat:
				value := field.Float()
				realField.delta = value - realField.value
				realField.value = value
			case *fieldComplex:
				value := field.Complex()
				realField.delta = value - realField.value
				realField.value = value
			}
		}
	})
}

func Load(obj interface{}) *Deltacmp {
	deltacmp := Deltacmp{
		fields: make(map[string]field),
	}
	forEachField(reflect.ValueOf(obj), "", func(field reflect.Value, name string) {
		fieldBase := fieldBase{
			name: name,
		}
		switch field.Kind() {
		case reflect.Bool:
			deltacmp.fields[name] = &fieldBool{
				fieldBase: fieldBase,
				value:     field.Bool(),
			}
		case reflect.String:
			deltacmp.fields[name] = &fieldString{
				fieldBase: fieldBase,
				value:     field.String(),
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			deltacmp.fields[name] = &fieldInt{
				fieldBase: fieldBase,
				value:     field.Int(),
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			deltacmp.fields[name] = &fieldUint{
				fieldBase: fieldBase,
				value:     field.Uint(),
			}
		case reflect.Float32, reflect.Float64:
			deltacmp.fields[name] = &fieldFloat{
				fieldBase: fieldBase,
				value:     field.Float(),
			}
		case reflect.Complex64, reflect.Complex128:
			deltacmp.fields[name] = &fieldComplex{
				fieldBase: fieldBase,
				value:     field.Complex(),
			}
		}
	})
	return &deltacmp
}

func forEachField(value reflect.Value, name string, fn func(reflect.Value, string)) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	typ := value.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := value.Field(i)
		fieldEntry := typ.Field(i)

		fullName := name + fieldEntry.Name

		if field.Kind() == reflect.Struct {
			forEachField(field, fullName+".", fn)
		} else {
			fn(field, fullName)
		}
	}
}
