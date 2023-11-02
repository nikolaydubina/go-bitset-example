package permission_test

import (
	"fmt"
	"testing"

	"github.com/nikolaydubina/go-bitset-example/permission"
)

func ExamplePermission_identity() {
	v := permission.Read
	fmt.Println(v)
	// Output: r--
}

func ExamplePermission_multiple() {
	v := permission.Read.Add(permission.Write).Add(permission.Execute)
	fmt.Println(v)
	// Output: rwe
}

func ExamplePermission_Contains_superset() {
	v := permission.Read.Contains(permission.Read.Add(permission.Write))
	fmt.Println(v)
	// Output: false
}

func ExamplePermission_Contains_subset() {
	v := permission.Read.Add(permission.Write).Contains(permission.Read)
	fmt.Println(v)
	// Output: true
}

func ExamplePermission_Contains_none() {
	v := permission.Read.Add(permission.Write).Contains(permission.Execute)
	fmt.Println(v)
	// Output: false
}

func BenchmarkPermission_Add(b *testing.B) {
	var s permission.Permission
	for i := 0; i < b.N; i++ {
		s = permission.Read.Add(permission.Write).Add(permission.Execute)
	}
	if s == permission.Execute {
		b.Error()
	}
}

func BenchmarkPermission_Contains(b *testing.B) {
	var s permission.Permission
	for i := 0; i < b.N; i++ {
		s = permission.Read.Add(permission.Write)
	}
	if s == permission.Execute {
		b.Error()
	}
}

func BenchmarkPermission_String(b *testing.B) {
	s := []permission.Permission{
		permission.Read,
		permission.Write,
		permission.Execute,
		permission.Read.Add(permission.Write),
		permission.Read.Add(permission.Execute),
		permission.Write.Add(permission.Execute),
		permission.Read.Add(permission.Write).Add(permission.Execute),
	}
	var v string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = s[i%len(s)].String()
	}
	if len(v) == 0 {
		b.Error()
	}
}
