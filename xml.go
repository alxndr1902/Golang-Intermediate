package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city,omitempty"`
	Email   string   `xml:"-"`
	Address Address  `xml:"address,omitempty"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	person := Person{
		Name:  "John",
		Age:   30,
		City:  "New York",
		Email: "john@example.com",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling data into XML:", err)
		return
	}
	fmt.Println("XML Data:", string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data into XML:", err)
		return
	}
	fmt.Println(string(xmlData1))

	xmlRaw := `<person><name>John</name><age>30</age><address><city>"New York</city><state>"NY"</state></address></person>`
	var personXml Person

	err = xml.Unmarshal([]byte(xmlRaw), &personXml)
	if err != nil {
		fmt.Println("Error unmarshalling data into XML:", err)
		return
	}
	fmt.Println(personXml)
	fmt.Println(personXml.XMLName.Local)
	fmt.Println(personXml.XMLName.Space)

	book := Book{
		Title:  "Book",
		Author: "alex",
	}

	xmlData2, err := xml.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data into XML:", err)
		return
	}
	fmt.Println(string(xmlData2))
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
}
