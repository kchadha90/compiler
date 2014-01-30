package testmylexer

type TestToken struct {
enum int
value string
location int
line int
column int
}

func (x TestToken) Value() string {
return x.value
}
func (x TestToken) Location() int {
return x.location
}
func (x TestToken) Line() int {
return x.line
}
func (x TestToken) Column() int {
return x.column
}
func (x TestToken) Enum() int {
return x.enum
}
