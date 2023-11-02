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
