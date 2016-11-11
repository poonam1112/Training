package interfaces

import "fmt"

type Planet interface {
	Name() string
	Mass() int
}

type pluto struct {
	name   string
	radius int
}

type mercury struct {
	name   string
	radius int
}

type earth struct {
	name   string
	radius int
}

type jupitor struct {
	name   string
	radius int
}

func (p pluto) Name() string {
	fmt.Println(" Planet with radius: %s", p)
	return p.name
}
func (p pluto) Mass() int {
	if p.radius < 1 {
		fmt.Println("invalid radius")
		return 0
	}
	return p.radius * p.radius * 314
}

func (m mercury) Name() string {
	fmt.Println(" Planet with radius: %s", m)
	return m.name
}
func (m mercury) Mass() int {
	if m.radius < 1 {
		fmt.Println("invalid radius")
		return 0
	}
	return m.radius * m.radius * 314
}

func (e earth) Name() string {
	fmt.Println(" Planet with radius: %s", e)
	return e.name
}
func (e earth) Mass() int {
	if e.radius < 1 {
		fmt.Println("invalid radius")
		return 0
	}
	return e.radius * e.radius * 314
}

func (j jupitor) Name() string {
	fmt.Println(" Planet with radius: %s", j)
	return j.name
}
func (j jupitor) Mass() int {
	if j.radius < 1 {
		fmt.Println("invalid radius")
		return 0
	}
	return j.radius * j.radius * 314
}

func Name(name string) string {
	return name
}
func Mass(radius int) int {
	if radius < 1 {
		fmt.Println("invalid radius")
		return 0
	}
	return radius * radius * 314
}

func main() {
	p := pluto{name: "Pluto", radius: 2}
	m := mercury{name: "mercury", radius: 7}
	e := earth{name: "earth", radius: 4}
	j := jupitor{name: "jupitor", radius: 8}

	p.Name()
	fmt.Println("Mass in Kg : ", p.Mass())
	m.Name()
	fmt.Println("Mass in Kg : ", m.Mass())
	e.Name()
	fmt.Println("Mass in Kg: ", e.Mass())
	j.Name()
	fmt.Println("Mass in Kg: ", j.Mass())
}
