
package dog

import "fmt"

type Dog struct {
	Name string
	Age float32
	Dad string
}

func (d Dog) GetName() string {
	fmt.Println(d.Name)
	return d.Name
}
func (d Dog) GetAge() float32 {
	fmt.Println(d.Age)
	return d.Age
}
func (d Dog) GetDad() string {
	fmt.Println(d.Dad)
	return d.Dad
}

func (d Dog) Eat() {
	fmt.Println("咔咔咔")
}
func (d Dog) Shout() {
	fmt.Println("汪汪汪")
}