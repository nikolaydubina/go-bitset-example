### How to make bitset in Go?

* *zero overhead*
* short syntax for creating sets from values
* set membership, equality, subset
* compile time
   * block of accidental arithmetics
   * block of implicit cast of untyped constants
   * block of all operators except `==` and `!=`
   * block of creating new values
* extending on the idea of [Go Enum](https://github.com/nikolaydubina/go-enum-example)

```go
package permission

var (
	Read    = Permission{0x1}
	Write   = Permission{0x2}
	Execute = Permission{0x4}
)

// Permission is a set implemented as a bitset of single byte.
type Permission struct{ v uint8 }

var ss = map[Permission]string{
	{}:                           "---",
	Read:                         "r--",
	Write:                        "-w-",
	Execute:                      "--e",
	Read.Add(Write):              "rw-",
	Read.Add(Execute):            "r-e",
	Write.Add(Execute):           "-we",
	Read.Add(Write).Add(Execute): "rwe",
}

func (a Permission) String() string { return ss[a] }

func (a Permission) Add(b Permission) Permission { return Permission{a.v | b.v} }

func (a Permission) Contains(b Permission) bool { return b.v&(a.v&b.v) == b.v }
```

### Benchmarks

```bash
$ go test -bench=. -benchmem ./permission
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/go-bitset-example/permission
BenchmarkPermission_Add-10         	1000000000	         0.7463 ns/op	       0 B/op	       0 allocs/op
BenchmarkPermission_Contains-10    	1000000000	         0.4688 ns/op	       0 B/op	       0 allocs/op
BenchmarkPermission_String-10      	132117963	         9.080 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/nikolaydubina/go-bitset-example/permission	3.552s
```

### References

* https://en.wikipedia.org/wiki/Bit_array
* https://en.cppreference.com/w/cpp/utility/bitset
* https://docs.oracle.com/javase/8/docs/api/java/util/BitSet.html
* https://godbolt.org
* https://github.com/nikolaydubina/go-enum-example
