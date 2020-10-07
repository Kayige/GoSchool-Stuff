package main

type print interface {
	print()
}

type list []print

func (l list) print() {
	for _, item := range l {
		item.print()
	}
}
