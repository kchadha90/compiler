package parser

import (
"testing"
)

func Test_simple(t *testing.T) {
s :=`
(* Fib code *)

MODULE mymod;

TYPE 
	string = ARRAY 11 OF CHAR;
	foo = ARRAY 10 OF INTEGER;
	foo1 = ARRAY 0 OF INTEGER;

VAR a,b,d : INTEGER;
	c : BOOLEAN;
	f1, f3 : foo;
	f2 : foo1;
	str : string;
	control : INTEGER;

PROCEDURE abc (VAR x : foo; y : BOOLEAN; VAR str1 : string);
	VAR
		a, b : BOOLEAN;
		z : INTEGER;
	
	BEGIN
		z := TRUE;

	END abc;

	BEGIN
	
	FOR control := (f1[20]+f2[2]*f3[1]-100/3/3) TO (control + f2[f1[f3[f2[10]]]]) BY 10000 DO

		f1 := f2;
		b := -f1[30]; 
		d := 2;
		f1[2] := 10;
		str := "hello world"; 
		(* The ultimate expression test *)
		c := (1 < f3[d+a]) & c OR c & ((1 + 2 + 3) < 10);
		(* A simple boolean expression *)
		c := (TRUE & TRUE) OR FALSE;
		
		FOR control := 0 TO 1 DO
			str := "lame!!"
		END;

		WHILE c DO
			(* A simple if control block tested *)
			IF str < str THEN 
				a := 2;
			ELSIF c OR (1 < 2) THEN
				a := f1[b+f2[control * 23 + f3[3]]];
					IF (c = c) OR (c # c) & (a # a) THEN
						c := FALSE;
					END;
			ELSE 
				a := 3;
			END;
		END;
	
	END;

END mymod.

`
driver(s)
}
