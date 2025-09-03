package main

type car struct {
	color string
}

func (c *car) changeColor(color string) {
	c.color = color
}

// c *car is just a reference to a "class"(? I'm not sure if it is to strictly to a class)
// changeColor is just a method that is manipulating the car struct
// Really fun.
