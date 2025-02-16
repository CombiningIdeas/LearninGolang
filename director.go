package main

type Director struct {
	builder BuilderI
}

func NewDirector(b BuilderI) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(builder BuilderI) {
	d.builder = builder
}

func (d *Director) BuildComputer() Computer {
	return d.builder.Build()
}
