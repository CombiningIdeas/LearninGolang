package main

// Builder

type Computer struct {
	CPU   string
	RAM   int
	MB    string
	Mouse string
}

//

type BuilderI interface {
	CPU(val string) BuilderI
	RAM(val int) BuilderI
	MB(val string) BuilderI
	Mouse(val string) BuilderI

	Build() Computer
}

//

type Builder struct {
	cpu   string
	ram   int
	mb    string
	mouse string
}

func NewBuilder() BuilderI {
	return Builder{}
}

func (b Builder) CPU(val string) BuilderI {
	b.cpu = val
	return b

}
func (b Builder) RAM(val int) BuilderI {
	b.ram = val
	return b

}
func (b Builder) MB(val string) BuilderI {
	b.mb = val
	return b
}
func (b Builder) Mouse(val string) BuilderI {
	b.mouse = val
	return b
}

func (b Builder) Build() Computer {
	return Computer{
		CPU:   b.cpu,
		RAM:   b.ram,
		MB:    b.mb,
		Mouse: b.mouse,
	}
}

//
//
//

type officeCompBuilder struct {
	Builder
}

func NewOfficeComputerBuilder() BuilderI {
	return officeCompBuilder{}.CPU("pentium III").RAM(2).MB("asrock").Mouse("logitech")
}

func (b officeCompBuilder) Build() Computer {
	return Computer{
		CPU:   b.cpu,
		RAM:   b.ram,
		MB:    b.mb,
		Mouse: b.mouse,
	}
}
