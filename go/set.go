package main

import (
	"fmt"
)

type Set struct {
	Elements map[string] string
}

type ElementOfSet struct {
	Part1 int
	Part2 string
}

func (s *Set) Add(elem interface {}) bool {
	str := fmt.Sprintf("%v", elem)
	_, ok := s.Elements[str]
	if ok == true {
		return false
	}

	s.Elements[str] = str
	return true
}

func (s *Set) Remove(elem interface{}) bool {
	str := fmt.Sprintf("%v", elem)
	_, ok := s.Elements[str]
	if ok == false {
		return false
	}
	delete(s.Elements, str)
	return true
}

func (s *Set) Cardinality() int {
	return len(s.Elements)
}
func (s *Set) IsMember(elem interface{}) bool {
	str := fmt.Sprintf("%v", elem)
	_, ok := s.Elements[str]
	return ok
}

func (s Set) Union(s2 Set) Set {
	u := Set{ make(map[string] string) }
	for _, elem := range s.Elements {
		u.Add(elem)
	}
	for _, elem := range s2.Elements {
		u.Add(elem)
	}
	return u
}

func (s Set) Intersection(s2 Set) Set {
	u := Set{ make(map[string] string) }
	for key, elem := range s.Elements {
		_, ok := s2.Elements[key]
		if ok == true {
			u.Add(elem)
		}
	}

	return u
}

func (s Set) SetTheoreticDifference(s2 Set) Set {
	u := Set{ make(map[string] string) }
	for key, _ := range s.Elements {
		_, ok := s2.Elements[key]
		if ok == false {
			u.Add(key)
		}
	}
	return u
}

func (s Set) SymmetricDifference(s2 Set) Set {
	u := s.Union(s2)
	is := s.Intersection(s2)
	return u.SetTheoreticDifference(is)
}

func (s *Set) IsSubsetOf(s2 Set) bool {
	for key, _ := range s.Elements {
		_, ok := s2.Elements[key]
		if ok == false {
			return false
		}
	}
	return true
}




/*
// all the Add function for various types
func (s *Set) Add(x int) bool {
	str := strconv.Itoa(x)
	return s.Add(str)
}

func (s *Set) Add(f float32) bool {
	str := fmt.Sprintf("%f", f)
	return s.Add(str)
}
*/


func main() {
	var s1 = Set { make(map[string] string) }
	s1.Add(54)
	s1.Add("the")
	s1.Add("bonsai")

	var s2 = Set { make(map[string] string) }
	s2.Add("the")
	s2.Add(67.22)

	var s4 = s1.Intersection(s2)
	fmt.Printf("Inter %+v\n", s4)
	var s5 = s1.Union(s2)
	fmt.Printf("Union %+v\n", s5)

	var s3 = s1.SetTheoreticDifference(s2)
	fmt.Printf("Theor diff %+v\n", s3)

	var s6 = s1.SymmetricDifference(s2)
	fmt.Printf("Sym diff %+v\n", s6)

	var s7 = Set { make(map[string] string) }
	var e1 = ElementOfSet{ 66, "brist" }
	s7.Add(e1)

	fmt.Printf("Sym diff %+v\n", s7)
}
