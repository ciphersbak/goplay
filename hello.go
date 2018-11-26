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
	var addressBook *AddressBook
	addressBook = nil
	addressBook = &AddressBook{
		Person: &Person{
			Name:  "Prashant",
			Age:   35,
			Email: "abc@abc.com",
		},
	}
	// data, err := proto.Marshal(prashant)
	data, err := proto.Marshal(addressBook)
	if err != nil {
		log.Fatalln("Marshalling error: ", err)
	}
	fmt.Println("Printing data: ", data)

	// newPrashant := &Person{}
	// err = proto.Unmarshal(data, newPrashant)
	newAddressBook := &AddressBook{}
	err = proto.Unmarshal(data, newAddressBook)
	// err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatalln("Unmarshalling error: ", err)
	}
	// fmt.Println(newPrashant.GetAge())
	// fmt.Println(newPrashant.GetName())
	// fmt.Println(newPrashant.GetEmail())
	// fmt.Println("Print after Unmarshall: ", newAddressBook.GetPerson())
	fmt.Println("Print after Unmarshall: ", newAddressBook.GetPerson().GetName())
	fmt.Println("Print after Unmarshall: ", newAddressBook.GetPerson().GetAge())
	fmt.Println("Print after Unmarshall: ", newAddressBook.GetPerson().GetEmail())
}
