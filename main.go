package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	p := &Person{
		Name: "Zakir",
		Age:  28,
		Addr: &Address{
			Add: "Bharatpur",
			HNo: "479",
		},
		NickName: []string{"A", "B", "C", "D"},
		Records:  map[string]string{"First": "PASS", "Second": "PASS"},
	}

	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Encoded Buffer:", data)

	p1 := &Person{}

	err = proto.Unmarshal(data, p1)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Name: ", p1.GetName())
	fmt.Println("Name: ", p1.GetNickName())
	fmt.Println("Name: ", p1.GetRecords())

}
