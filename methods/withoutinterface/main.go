package main

import "fmt"

type UserA struct {
	Name      string
	Address   string
	ContectNo int
}
type UserB struct {
	Name      string
	Address   string
	ContectNo int
}

func main() {

	var vishal UserA
	vishal.Name = "vishal singh"
	vishal.Address = "varanasi"
	vishal.ContectNo = 44961346
	Shivam := UserB{
		Name:      "shivam singh",
		Address:   "aurai",
		ContectNo: 73828957,
	}
	vishal.adduser(5)
	fmt.Println("vishal after adduser =", vishal)
	value := vishal.adduserdub(347587376)
	fmt.Println("value from adddub = ", value, "vishal =", vishal)
	Shivam.adduser(5)
	add(23255)

}

func add(a int) {
	fmt.Println("function run ho rha hai ,aur a=", a)
}

func (user UserA) adduser(a int) {
	user.Name = "24478472389754857"
	fmt.Println("User A wala method chal rha hai,user=", user)
}
func (u UserA) adduserdub(a int) int {
	u.Name = "24478472389754857"
	fmt.Println("User A Dub wala method chal rha hai,user=", u)
	return 0
}
func (user UserB) adduser(a int) {
	fmt.Println("User Bwala method chal rha hai,user=", user)
}
