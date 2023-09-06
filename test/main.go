package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	elliot := &Person{
		Name: "Elliot",
		Age:  24,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		return
	}

	fmt.Println(data)

	var newElliot = new(Person)

	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(newElliot.Age)
	fmt.Println(newElliot.Name)
}
