package safer

import (
	"reflect"
	"testing"
)

func BenchmarkKindOf(b *testing.B) {
	var (
		k       reflect.Kind
		exp     = reflect.Int
		kindInt = 0
	)

	k = reflect.Invalid
	b.Run("KindOf", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			kind := KindOf(kindInt)
			if kind != exp {
				b.Fatalf(`exp %v; got %v`, exp, kind)
			}
			k = kind
		}
	})
	if k != exp {
		b.Fatalf(`exp %v; got %v`, exp, k)
	}

	b.Run("Reflection", func(b *testing.B) {

		k = reflect.Invalid
		b.Run("ValueOf", func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				kind := reflect.ValueOf(kindInt).Kind()
				if kind != exp {
					b.Fatalf(`exp %v; got %v`, exp, kind)
				}
				k = kind
			}
		})
		if k != exp {
			b.Fatalf(`exp %v; got %v`, exp, k)
		}

		k = reflect.Invalid
		b.Run("TypeOf", func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				kind := reflect.TypeOf(kindInt).Kind()
				if kind != exp {
					b.Fatalf(`exp %v; got %v`, exp, kind)
				}
				k = kind
			}
		})
		if k != exp {
			b.Fatalf(`exp %v; got %v`, exp, k)
		}
	})
}

func BenchmarkPCForFunc(b *testing.B) {
	b.Run("PCForFunc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pc := PCForFunc(KindOf)
			if pc == 0 {
				b.Fatal(`exp non zero pc`)
			}
		}
	})
	b.Run("Reflection", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			val := reflect.ValueOf(KindOf)
			if val.Kind() != reflect.Func {
				b.Fatal(`exp non zero pc`)
			}

			pc := val.Pointer()
			if pc == 0 {
				b.Fatal(`exp non zero pc`)
			}
		}
	})
}
