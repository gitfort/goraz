package main

import (
	"github.com/gitfort/goraz/pkg/mapper"
	"log"
)

type Person struct {
	Parents   []*Person
	FirstName string
	LastName  string
	Rank      int
	Tags      [][]string
}
type Actor struct {
	Parents   []*Actor
	FirstName string
	LastName  string
	Rank      int32
	Tags      [][]string
}

func main() {
	persons := []*Person{
		{
			Parents: []*Person{
				{
					Parents:   nil,
					FirstName: "mohsen",
					LastName:  "samiei",
				},
				{
					Parents:   nil,
					FirstName: "monire",
					LastName:  "besharati",
				},
			},
			FirstName: "amirabbas",
			LastName:  "samiei",
			Rank:      32,
			Tags: [][]string{
				{
					"1", "2", "3",
				},
				{
					"A", "B", "C",
				},
			},
		},
	}
	var actors []*Actor
	mapper.MapSlice(persons, &actors)
	log.Printf("%+v", actors[0])

	array := [][]string{
		{
			"1", "2", "3",
		},
		{
			"A", "B", "C",
		},
	}
	var array2 [][]string
	mapper.MapSlice(array, &array2)
	log.Printf("%+v", array2)
}
