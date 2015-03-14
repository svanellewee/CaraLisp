package main

import "fmt"

type Any interface{}
type FuncType func(Any, Any) Any
type ConsType func(FuncType) Any

var (
	car, cdr FuncType
)

func init() {
	car = func(head, tail Any) Any { return head }
	cdr = func(head, tail Any) Any { return tail }
}

// eq [X; X] = T
// eq [X; A] = F
// eq [X; (X · A)] is undefined.
func Eq(a, b Any) bool {
	switch aVal := a.(type) {
	case int:
		bVal, bCastOk := b.(int)
		if !bCastOk {
			return false
		}

		retval := aVal == bVal
		return retval
	case string:
		bVal, bCastOk := b.(string)
		if !bCastOk {
			return false
		}

		retval := aVal == bVal
		return retval

	case bool:
		bVal, bCastOk := b.(bool)
		if !bCastOk {
			return false
		}

		retval := aVal == bVal
		return retval

	default:
		panic("I dont support EQ for this type yet")
	}
	return false
}

func Atom(e Any) bool {
	switch e.(type) {
	case string, int, bool:
		return true
	case ConsType:
		return false
	default:
		panic("Not supported type")
	}
}

func List(vals ...Any) ConsType {
	if len(vals) == 1 {
		return Cons(vals[0], nil)
	}
	return Cons(vals[0], List(vals[1:]...))
}

// pair[x; y] = [null[x]∧null[y] → NIL; ¬atom[x]∧¬atom[y] → cons[list[car[x]; car[y]]; pair[cdr[x]; cdr[y]]]
func Pair(a, b ConsType) ConsType {
	//fmt.Println("Start",a,b)
	if a == nil && b == nil {
		//fmt.Println("a nil b nil")
		return nil
	}
	if !Atom(a) && !Atom(b) {
		//fmt.Printf("2atom atom %s, %s, %T, %T\n",a,b,a,b)

		cdrA, _ := Cdr(a).(ConsType)
		cdrB, _ := Cdr(b).(ConsType)

		if cdrA == nil || cdrB == nil {
			return Cons(List(Car(a), Car(b)), nil)
		}
		return Cons(List(Car(a), Car(b)), Pair(cdrA, cdrB))
	}
	return nil
}

/*
         ((12,(14|,16|)),(13,(15|,17|)))
assoc[X; ((W ,(A, B   )),(X, (C  , D )),(Y,(E, F)))]

*/
// assoc[x; y] = eq[caar[y]; x] → cadar[y]; T → assoc[x; cdr[y]]]
func Assoc(x Any, y ConsType) Any {
	if Eq(Caar(y), x) {
		return Cdar(y) // differs from McArthy here (cadar vs cdar) not sure why ?
	}

	rest, castOk := Cdr(y).(ConsType)
	_ = castOk
	if !castOk {
		return nil
	}
	return Assoc(x, rest)
}

func (c ConsType) String() string {
	var printVal func(Any, Any) Any
	printVal = func(a, b Any) Any {
		var result string
		if aVal, casOk := a.(ConsType); casOk {
			result = fmt.Sprintf("%s", aVal(printVal))
		} else {
			result = fmt.Sprintf("%v", a)
		}
		if b == nil {
			return fmt.Sprintf("%v|", result)
		}
		if bVal, casOk := b.(ConsType); casOk {
			return fmt.Sprintf("(%v,%v)", result, bVal(printVal))
		}
		return result + "." + fmt.Sprintf("%v", b)
	}
	return fmt.Sprintf("%v", c(printVal))
}

func Cons(a, b Any) ConsType {
	return func(fn FuncType) Any {
		return fn(a, b)
	}
}
func Car(cons ConsType) Any {
	return cons(car)
}

func Cdr(cons ConsType) Any {
	return cons(cdr)
}

func Cddr(cons ConsType) Any {
	tail := Cdr(cons)
	if tail == nil {
		return nil
	}
	tcons, _ := tail.(ConsType)
	if tcons == nil {
		return tail
	}
	return tcons(cdr)
}

func Cadr(cons ConsType) Any {
	tail := Cdr(cons)
	if tail == nil {
		return nil
	}
	tcons, _ := tail.(ConsType)
	if tcons == nil {
		return tail
	}
	return tcons(car)
}

func Cdar(cons ConsType) Any {
	head := cons(car)
	if head == nil {
		return nil
	}
	tcons, _ := head.(ConsType)
	if tcons == nil {
		return head
	}
	return tcons(cdr)
}

func Caar(cons ConsType) Any {
	x := Car(cons)
	if x == nil {
		return nil
	}
	tcons, _ := x.(ConsType)
	if tcons == nil {
		return x
	}
	return tcons(car)
}

func Caddr(cons ConsType) Any {
	x := Cddr(cons)
	if x == nil {
		return nil
	}
	y, _ := x.(ConsType)
	if y == nil {
		return x
	}
	return y(car)
}

func Cadar(cons ConsType) Any {
	x := Cdar(cons)
	if x == nil {
		return nil
	}
	y, _ := x.(ConsType)
	if y == nil {
		return x
	}
	return y(car)
}

///func isAtom(cons ConsType)

func main() {
	fmt.Printf("Hello, 世界\n")
	// list := Cons(12, Cons(14, Cons(16, nil)))

	// fmt.Println(list(car), list(cdr))
	// fmt.Println(list(car), list(cdr).(ConsType)(car))
	// fmt.Println(Cddr(list),Cadr(list), Cadar(list))

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

	fmt.Println("1>>", Assoc(12, list2))
	fmt.Println("2>>", Assoc(13, list2))
	fmt.Println("3>>", Assoc(16, list2))
	fmt.Println("4>>", Assoc(18, list2))
	fmt.Println(">>..", list2)
	/*
	   list3 = ( (1, ("hello", nil)),
	             ( (2, ("byebye",nil)),
	               nil) )
	*/
	list3 := Cons(
		Cons(1, Cons("hello", nil)),
		Cons(
			Cons(2, Cons("byebye", nil)),
			nil))
	fmt.Println("NewList >>", list3, Assoc(1, list3), Assoc(2, list3))
	fmt.Println("atom test", Atom(list3), Atom(12), Atom("hello"))
	fmt.Println("List test", List(12, 15, 12, 155, 8), List(17, 20))
	fmt.Println("List test", List(List(12, 15), List(14, 155), 8), List(17, 20))
	fmt.Println("Pairtest", Pair(List(12, 13, 20), List("bla", "Yadda", "yurface")))
}

// >>.. ((12,(14|,16|)),(13,(15|,17|)))
//      ((W, (A,  B)),  (X,(C, D)),    (Y,(E, F)))
// >>.. (12,(14|,16|),13,(15|,17|))
