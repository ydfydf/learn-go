	package others

	import "fmt"

	type Base struct {
		FirstName, LastName string
		Age float32
	}

	func (base *Base) HasFeet() {
		fmt.Println(base.FirstName + base.LastName + "has feet! Base")
	}

	func (base *Base) Drink() {
		fmt.Println("Base Drink")
	}

	func (base *Base) Flying() {
		fmt.Println("Base Can flying!")
	}

	type Sub struct {
		Base
		Area string
	}

	func (sub Sub) SubFlying() {
		sub.Base.Flying()
		fmt.Println("Sub flying")
	}

	//如果“子类”重写了“基类”的成员方法，需要在子类的成员方法中调用基类的同名成员方法，一定要以sub.Base.Flying()这样显式的方法调用，
	// 而不是使用sub.Flying()这种调用继承方法的方式调用，这样会出现无限循环，即一直在调用子类的方法
	//并且，如果“子类”重写了“基类”的成员方法，会发生覆盖现象，即“子类”的同名方法覆盖了“基类”的同名方法
	func (sub Sub) Flying() {
		sub.Base.Flying()
		fmt.Println("Sub flying")
	}

	//利用组合的方式，拓展了原Base结构，使得原Base结构能正常使用，新的拓展也能正常使用
	func (sub *Sub) SetBaseInfo(firstname,lastname string,age float32,area string) {
		sub.FirstName = firstname
		sub.LastName = lastname
		sub.Age = age
		sub.Area = area
	}

	func (sub Sub) PrintBaseInfo() {
		fmt.Println(sub.FirstName)
		fmt.Println(sub.LastName)
		fmt.Println(sub.Age)
		fmt.Println(sub.Area)
	}

	func (sub *Sub) SetArea() {
		sub.Area = "cd"
	}

	func main() {
		chk := new(Sub)
		chk.Flying()//可以直接使用Base结构的变量或者方法，体现了继承特性
		chk.SubFlying()
		chk.Base.Flying()
		chk.Drink()
		fmt.Println("==============================")
		chk2 := Sub{Base{"Bob", "Steven", 2.0}, "China"}
		fmt.Println(chk2.Area)
		chk2.SetArea()
		fmt.Println(chk2.Area)

		chk.SetBaseInfo("yang","fan",25,"chengdu")
		chk.PrintBaseInfo()
	}