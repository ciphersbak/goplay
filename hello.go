package main

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Printf("hello, world\n")

	// prashant := &Person{
	// 	Name:  "Prashant",
	// 	Age:   35,
	// 	Email: "abc@abc.com",
	// }
	// addressBook := &AddressBook{}
	person := &Person{
		Name:  "Prashant",
		Age:   35,
		Email: "abc@abc.com",
	}

	// data, err := proto.Marshal(prashant)
	// data, err := proto.Marshal(addressBook)
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalln("Marshalling error: ", err)
	}
	fmt.Println("Printing data: ", data)

	// newPrashant := &Person{}
	// err = proto.Unmarshal(data, newPrashant)
	// newAddressBook := &AddressBook{}
	newPerson := &Person{}
	// err = proto.Unmarshal(data, newAddressBook)
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatalln("Unmarshalling error: ", err)
	}
	// fmt.Println(newPrashant.GetAge())
	// fmt.Println(newPrashant.GetName())
	// fmt.Println(newPrashant.GetEmail())
	// fmt.Println("Print after Unmarshall: ", newAddressBook.GetPerson())
	fmt.Println("Print after unmarshall: ", newPerson.GetEmail())
	fmt.Println("Print after unmarshall: ", newPerson.GetAge())
	fmt.Println("Print after unmarshall: ", newPerson.GetName())
}
