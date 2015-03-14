package main

import (
	"fmt"
	"testing"
)

func __string(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func TestSingleCons(t *testing.T) {
	testcell := Cons(1, 2)
	expected := "1.2"
	strValue := __string(testcell)
	if strValue != expected {
		t.Errorf("expected ", expected, ", got ", strValue)
	}
}

func TestListCons(t *testing.T) {
	testcell := Cons(1, Cons(2, Cons(3, nil)))
	expected := "(1,(2,3|))"
	strValue := __string(testcell)
	if strValue != expected {
		t.Errorf("expected ", expected, ", got ", strValue)
	}
}

// Accessors
func TestCar(t *testing.T) {
	testcells := Cons(1, Cons(2, Cons(3, nil)))
	expected := "(1,(2,3|))"
	strValue := __string(testcells)
	if strValue != expected {
		t.Errorf("expected %s got %s", expected, strValue)
	}

	carVal := Car(testcells)
	expected2 := 1
	if carVal != expected2 {
		t.Errorf("expected %s, got %s", expected2, carVal)
	}
}

func TestCdr(t *testing.T) {
	testcells := Cons(1, Cons(2, Cons(3, nil)))
	expected := "(1,(2,3|))"
	strValue := __string(testcells)
	if strValue != expected {
		t.Errorf("expected %s got %s", expected, strValue)
	}
	cdrVal := __string(Cdr(testcells))
	expected3 := "(2,3|)"
	if cdrVal != expected3 {
		t.Errorf("expected  [%s], got [%s] ", expected3, cdrVal)
	}

}

func TestCadr(t *testing.T) {
	testcells := Cons(1, Cons(2, Cons(3, nil)))
	expected := "(1,(2,3|))"
	strValue := __string(testcells)
	if strValue != expected {
		t.Errorf("expected %s got %s", expected, strValue)
	}

	cadrVal := __string(Cadr(testcells))
	expected4 := "2"
	if cadrVal != expected4 {
		t.Errorf("expected  [%s], got [%s] ", expected4, cadrVal)
	}

}

func TestCddr(t *testing.T) {
	testcells := Cons(1, Cons(2, Cons(3, nil)))
	expected := "(1,(2,3|))"
	strValue := __string(testcells)
	if strValue != expected {
		t.Errorf("expected %s got %s", expected, strValue)
	}

	cddrVal := __string(Cddr(testcells))
	expected5 := "3|"
	if cddrVal != expected5 {
		t.Errorf("expected  [%s], got [%s] ", expected5, cddrVal)
	}
}

func TestAtom(t *testing.T) {
	if !Atom("string") {
		t.Errorf("string not considered atom anymore")
	}

	if !Atom(12) {
		t.Errorf("int not considered atom anymore")
	}

	if !Atom(true) {
		t.Errorf("bool not considered atom anymore")
	}

	if Atom(Cons(1, 2)) {
		t.Errorf("ConsType is now considered an atom")
	}

}

func TestPair(t *testing.T) {
	list1, list2 := List(12, 13, 20), List("this", "is", "SPARTAAAA!!!")
	expected := "((12,this|),((13,is|),(20,SPARTAAAA!!!|)|))"
	value := fmt.Sprintf("%v", Pair(list1, list2))
	if value != expected {
		t.Errorf("Pairs not generated correctly [%s] instead of [%s]", value, expected)
	}
}

func TestEq(t *testing.T) {
	var a, b interface{}
	a = 12
	b = 13
	if Eq(a, b) {
		t.Errorf("Values are not equal Eq failed")
	}

}

func TestAssoc(t *testing.T) {
	// works like a map in a way
	list2 := Cons(
		Cons(12,
			Cons(
				Cons(14, nil),
				Cons(16, nil),
			)),
		Cons(
			Cons(13,
				Cons(
					Cons(15, nil),
					Cons(17, nil),
				)),
			Cons(
				Cons(16,
					Cons(
						Cons(20, nil),
						Cons(38, nil),
					)),
				nil)))
	// _ = list2

	tests := map[int]string{
		12: "(14|,16|)",
		13: "(15|,17|)",
		16: "(20|,38|)",
		18: "<nil>",
	}

	for key, expected := range tests {
		value := __string(Assoc(key, list2))
		if value != expected {
			t.Errorf("Found %s instead of %s", value, expected)
		}
	}

}

// Test Assoc!
