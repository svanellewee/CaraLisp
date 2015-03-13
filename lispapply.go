package lisp

const (
	QUOTE = "'"
	ATOM = "atom"
	EQ = "="
	COND = "cond"
	CAR = "car"
	CDR = "cdr"
	CONS = "cons"
	LABEL = "label"
	LAMBDA = "lambda"
)



func Eval(e, a interface{}) interface{} {
	if isAtom(e) {
		return Assoc(e,a)
	}
	carE := Car(e)
	if isAtom(carE) {
		switch(carE) {
		case QUOTE:
			return Cadr(e)
		case ATOM:
			return isAtom(Eval(Cadr(e), a))
		case EQ:
			return Eval(Cadr(e,a)) == Eval(Caddr(e),a)
		case COND:
			return EvCon(Cdr(e), a)
		case CAR:
			return Car(Eval(Cadr(e), a))
		case CDR:
			return Cdr(Eval(Cadr(e), a))
		case CONS:
			return Cons(Eval(Cadr(e),a), Eval(Caddr(e),a))
		default:
			return Eval(Cons(Assoc(Car(e),a), EvLis(Cdr(e),a)), a)
		}
	} else {
		if Car(CarE) == LABEL {
			newE := Cons(Caddr(e), Cdr(e))
			newA := Cons(List(Cadar(e), Car(e), a))
			return Eval(newE, newA)
		}
		if Car(CarE) == LAMBDA {
			newE := Caddar(e)
			newA :=  Append( Pair( Cadar(e), EvLis(Cdr(e),a)), a)
			return Eval(newE,newA) // ---> brackets are wrong here I think
		}
	}
	
}

// evcon[c; a] = [eval[caar[c]; a] → eval[cadar[c]; a]; T → evcon[cdr[c]; a]]
func EvCon(c, a interface{}) interface{} {
	if Eval(Caar(c),a) == true {
		return Eval(Cadar(c),a)
	} else {
		return EvCon(Cdr(c),a)
	}	
}
// and
// evlis[m; a] = [null[m] → NIL; T → cons[eval[car[m]; a]; evlis[cdr[m]; a]]]
func EvLis(m, a interface{}) interface{} {
	if isNull(m) {
		return NIL
	} else {
		return Cons(Eval(Car(m), a), EvLis(Cdr(m), a))
	}
}
/*
eval[e; a] = [
     atom [e] → assoc [e; a];
     atom [car [e]] → [
          eq [car [e]; QUOTE] → cadr [e];
          eq [car [e]; ATOM] → atom [eval [cadr [e]; a]];
          eq [car [e]; EQ] → [eval [cadr [e]; a] = eval [caddr [e]; a]];
          eq [car [e]; COND] → evcon [cdr [e]; a];
          eq [car [e]; CAR] → car [eval [cadr [e]; a]];
          eq [car [e]; CDR] → cdr [eval [cadr [e]; a]];
          eq [car [e]; CONS] → cons [eval [cadr [e]; a]; eval [caddr [e];a]];
          T → eval [cons [assoc [car [e]; a]; evlis [cdr [e]; a]]; a]
      ];
      eq [caar [e]; LABEL] → eval [cons [caddar [e]; cdr [e]]; cons [list [cadar [e]; car [e]; a]];
      eq [caar [e]; LAMBDA] → eval [caddar [e];append [pair [cadar [e]; evlis [cdr [e]; a]; a]]]
*/
