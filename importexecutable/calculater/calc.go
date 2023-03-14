package calculater


type User struct {
	Name      string
	contectno int
}
type Student struct {
	Name   string
	Rollno int
}

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}
func Mul(a, b int) int {
	return a * b
}
func Div(a, b int) int {
	return a / b
}
