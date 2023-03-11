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
type useroperations interface{
adduser()

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

	// vishal.adduser()
	// Shivam.adduser()

	var useroperations useroperations
	useroperations=vishal
	useroperations.adduser()
	useroperations=Shivam
	useroperations.adduser()
}
func (user UserA) adduser(){
	user.Name = "24478472389754857"
	fmt.Println("User A wala method chal rha hai,user=", user)
}

func (user UserB) adduser() {
	fmt.Println("User Bwala method chal rha hai,user=", user)
}
