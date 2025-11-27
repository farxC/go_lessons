package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func UsingStructEmbedding() {

	// When creating with literals we have to initalize the embedding explicitly; Here the embedded type serves as the fild name
	co := container{
		base: base{
			num: 1,
		},
		str: "Some some",
	}
	// So we can acess base's fields directly
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// But also we can access the full path using the embedded type name.
	fmt.Println("also num:", co.base.num)

	// `base` methods also become methods of a `container`
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Powerful..
	var d describer = co
	fmt.Println("describer:", d.describe())
}
