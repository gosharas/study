package main

import "fmt"

type isStuff interface {
	DoStuff()
}

type Stuff struct {
	isStuff
	Name string
}

type strStuff string

type intStuff int

func (str strStuff) DoStuff() {
	fmt.Println("strStuff")
}

func (in intStuff) DoStuff() {
	fmt.Println("intStuff")
}

func (s Stuff) SomeComplex() {
	fmt.Println(s.Name)
	s.DoStuff()
}

type Flyer interface {
	Fly()
	Greet()
}

type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println("Fly ", b.Name)
}

func (b Bird) Greet() {
	fmt.Println("Greeting")
}

func DoFly(f Flyer) {
	f.Fly()
	f.Greet()
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{X: 1, Y: 2}
	v2 = Vertex{2, 3}
	v3 = Vertex{X: 1}
	v4 = Vertex{}
	v5 = &Vertex{1, 2}
)

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	v5.X = 3
	fmt.Println(v1, v2, v3, v4, *v5)
	i, j := 44, 55
	fmt.Println(j)
	p := &i
	fmt.Println(*p)
	*p = 11
	fmt.Println(i)

	primes := [...]int{1, 2, 3, 4, 5}
	fmt.Println(primes)

	var pow = []int{1, 2, 3, 4, 5, 6}
	for i, j = range pow {
		fmt.Println(i, j)
	}

	m := make(map[string]int)
	m["lol"] = 1
	m["lol1"] = 2
	elem, ok := m["lol"]
	fmt.Println(elem, ok)

	bird := Bird{"loler"}
	DoFly(bird)

	strst := strStuff("lol")
	intstr := intStuff(8)

	r1 := Stuff{strst, "r1"}
	r2 := Stuff{intstr, "r2"}

	r1.DoStuff()
	r2.DoStuff()

}
