package main

type Book struct {
	Id     int
	Name   string
	Author string
}

var Books = []Book{
	Book{
		Id:     1,
		Name:   "Mastering Linux",
		Author: "RedHat",
	},
	Book{
		Id:     2,
		Name:   "Mastering Golang",
		Author: "Google",
	},
	Book{
		Id:     3,
		Name:   "Mastering Python",
		Author: "Golsing",
	},
	Book{
		Id:     4,
		Name:   "Mastering AWS",
		Author: "Amazon",
	},
	Book{
		Id:     5,
		Name:   "Mastering PERL",
		Author: "PERLER",
	},
}
