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
	addressBook := &AddressBook{}

	// data, err := proto.Marshal(prashant)
	data, err := proto.Marshal(addressBook)
	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}
	fmt.Println("Printing data: ", data)

	// newPrashant := &Person{}
	// err = proto.Unmarshal(data, newPrashant)
	newAddressBook := &AddressBook{}
	err = proto.Unmarshal(data, newAddressBook)
	if err != nil {
		log.Fatal("Unmarshalling error: ", err)
	}
	// fmt.Println(newPrashant.GetAge())
	// fmt.Println(newPrashant.GetName())
	// fmt.Println(newPrashant.GetEmail())
	fmt.Println("Print after Unmarshall: ", newAddressBook.GetPerson())
}
