package main

import (
	"os"

	"github.com/xconstruct/go-pushbullet"
)

func main() {

	pb := pushbullet.New("xxxxxxxxxxxxxxxxxx")
	devs, err := pb.Devices()
	if err != nil {
		panic(err)
	} else {
		err := pb.PushNote(devs[0].Iden, os.Args[1], os.Args[2])
		if err != nil {
			panic(err)
		}

	}
}
