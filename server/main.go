package main

import (
	"errors"
	"gorpcts/gorts"
)

type Calculator struct{}

type Testing struct {
	VarA int
	VarB string
}

type Args struct {
	A int
	B int
}

func (c *Calculator) Addition(args *Args, reply *int) error {
	if args == nil {
		return errors.New("args empty")
	}

	*reply = args.A + args.B
	return nil
}

func (c *Calculator) Multiply(args *Args, reply *int) error {
	if args == nil {
		return errors.New("args empty")
	}

	*reply = args.A * args.B
	return nil
}

func main() {
	calc := &Calculator{}
	server := gorts.NewGorts(8080)

	err := server.Register(calc)
	if err != nil {
		panic(err)
	}

	err = server.Initiate()
	if err != nil {
		panic(err)
	}
}
