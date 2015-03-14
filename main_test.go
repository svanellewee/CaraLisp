package main

import (
	"fmt"
	"testing"
)

func __string(v interface{}) string {
	return fmt.Sprintf("%v",v)
}

func TestSingleCons(t *testing.T) {
	testcell := Cons(1,2)
	expected := "1.2"
	strValue := __string(testcell)
	if strValue != expected {
		t.Errorf("expected ",expected, ", got ",strValue)
	}
}

func TestListCons(t *testing.T) {
	testcell := Cons(1,Cons(2,Cons(3,nil)))
	expected := "(1,(2,3|))"
	strValue := __string(testcell)
	if strValue != expected {
		t.Errorf("expected ",expected,", got ",strValue)
	}
}


func TestAccessors(t *testing.T) {
	testcells := Cons(1,Cons(2,Cons(3,nil)))
	expected := "(1,(2,3|))"
	strValue := __string(testcells)
	if strValue != expected {
		t.Errorf("expected %s got %s",expected,strValue)
	}
	
	carVal := Car(testcells)
	expected2 := 1
	if carVal != expected2 {
		t.Errorf("expected %s, got %s", expected2, carVal)
	}

	cdrVal := __string(Cdr(testcells))
	expected3 := "(2,3|)"
	if cdrVal != expected3{
		t.Errorf("expected  [%s], got [%s] ", expected3, cdrVal)
	}


	cadrVal := __string(Cadr(testcells))
	expected4 := "2"
	if cadrVal != expected4{
		t.Errorf("expected  [%s], got [%s] ", expected4, cadrVal)
	}


	cddrVal := __string(Cddr(testcells))
	expected5 := "3|"
	if cddrVal != expected5{
		t.Errorf("expected  [%s], got [%s] ", expected5, cddrVal)
	}


}
