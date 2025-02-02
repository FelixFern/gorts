package main

import (
	"errors"
	"fmt"
	"gorpcts/gorts"
)

type Dog struct{}
type Cat struct{}
type Duck struct{}

type DogArgs struct {
	Name string
}
type CatArgs struct {
	Name string
}
type DuckArgs struct {
	Name string
}

func (c *Dog) Bark(args *DogArgs, reply *string) error {
	if args == nil {
		return errors.New("args empty")
	}

	*reply = fmt.Sprintf("%s barked", args.Name)
	return nil
}

func (c *Cat) Bark(args *CatArgs, reply *string) error {
	if args == nil {
		return errors.New("args empty")
	}

	*reply = fmt.Sprintf("%s barked", args.Name)
	return nil
}

func (c *Duck) Quack(args *DuckArgs, reply *string) error {
	if args == nil {
		return errors.New("args empty")
	}

	*reply = fmt.Sprintf("%s barked", args.Name)
	return nil
}

func main() {
	dog := &Dog{}
	cat := &Cat{}
	duck := &Duck{}
	server := gorts.NewGorts(8080)

	err := server.Register(dog)
	if err != nil {
		panic(err)
	}

	err = server.Register(cat)
	if err != nil {
		panic(err)
	}

	err = server.Register(duck)
	if err != nil {
		panic(err)
	}

	err = server.Initiate()
	if err != nil {
		panic(err)
	}
}
