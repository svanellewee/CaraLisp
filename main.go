package main

import "fmt"

type Any interface{}
type FuncType func(Any, Any) Any
type ConsType func(FuncType) Any

var (
	car,cdr FuncType
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
	case int :
		bVal, bCastOk := b.(int)
		//fmt.Println("...", aVal , "vs", bVal)
		if !bCastOk {
			//fmt.Println(">>CastBad")
			return false
		}
		
		retval := aVal == bVal
		//fmt.Println(">>T",retval)
		return retval
	default:
		panic("I dont support EQ for this type yet")
	}
	return false
}

// assoc[x; y] = eq[caar[y]; x] → cadar[y]; T → assoc[x; cdr[y]]]
func Assoc(x Any, y ConsType) Any {
	fmt.Println("start",Caar(y), x)
	if Eq(Caar(y), x) {
		fmt.Println("EQUAL",Caar(y), x, "stoppping", Cadar(y))
		return Cadar(y)
	}
	fmt.Println("NOT EQUAL",Caar(y), x)

	rest, castOk :=Cdr(y).(ConsType)
	_ = castOk
	if !castOk {
		fmt.Println("Cast not okay",Cdr(y))
		return nil
	}
	fmt.Println("rest >", rest)
	fmt.Println("Stop2")

	return Assoc(x, rest)
}

// func PrintCons(c ConsType) {
// 	c(func(a,b Any ) Any {
// 		fmt.Printf("Cons[%v,%v]\n",a,b)
// 		return nil
// 	})
// }

func (c ConsType) String() string {
	s := "("
	accum := func(input interface{}) {
		//fmt.Println("{{", s)
		if input == nil {
			s += "nil)"
			return
		}
		
		s += fmt.Sprintf("%v,",input)
	}

	var recurse func (ConsType)
	recurse = func( c ConsType) {
		c(func (a, b Any) Any {
			accum(a)
			bCons, isCons := b.(ConsType);
			if  !isCons{
				accum(b)
			}
			if bCons == nil {
				return nil
			}
				recurse(bCons)
			return nil
		})
	}
	recurse(c)
	return s
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
	if tail == nil{
		return nil
	}
	tcons,_ := tail.(ConsType)
	if tcons == nil{
		return tail
	}
	return tcons(cdr)
}

func Cadr(cons ConsType) Any {
	tail := Cdr(cons)
	if tail == nil{
		return nil
	}
	tcons,_ := tail.(ConsType)
	if tcons == nil{
		return tail
	}
	return tcons(car)
}

func Cdar(cons ConsType) Any {
	head := cons(car)
	if head == nil {
		return nil
	}
	tcons,_ := head.(ConsType)
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
	tcons,_ := x.(ConsType)
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
	y,_ := x.(ConsType)
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
	y,_ := x.(ConsType)
	if y == nil {
		return x
	}
	return y(car)
}

///func isAtom(cons ConsType) 

func main() {
	fmt.Printf("Hello, 世界\n")
	list := Cons(12, Cons(14, Cons(16, nil)))

	fmt.Println(list(car), list(cdr))
	fmt.Println(list(car), list(cdr).(ConsType)(car))
	fmt.Println(Cddr(list),Cadr(list), Cadar(list))

	list2 := Cons(
		Cons(12, Cons(
			Cons(14, nil), 
			Cons( 16, nil),
		)), 
		Cons(13, 
			Cons(
				Cons(15, nil),
				Cons(17,nil),
			)))
	_ = list2
	fmt.Println(">>",Assoc(13, list2), "..",list.String())
}

