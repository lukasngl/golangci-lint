//golangcitest:args -Egodogen
package testdata

/*
 #include <stdio.h>
 #include <stdlib.h>

 void myprint(char* s) {
 	printf("%d\n", s);
 }
*/
import "C"

import (
	"unsafe"
	"context"
)

func _() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

//godogen:given correct (number) of groups
//godogen:given (too) (many) groups // want `pattern has 2 groups, but function has 1 regular parameters`
//godogen:given not enough of groups // want `pattern has 0 groups, but function has 1 regular parameters`
func paramGroupMismatch(ctx context.Context, number string) error {
	return nil
}

//godogen:given wrong (parameter) type // want `pattern has 1 groups, but function has 0 regular parameters`
func wrongParameterType(ctx context.Context, dogs rune) error { // want `parameter "dogs" has unexpected type "rune"`
	return nil
}

//godogen:then we have a wrong return type
func wrongReturnType() (string, error) { // want `function should return \(context.Context, error\) with two return values`
	return "", nil
}

//godogen:step )malformed pattern( // want `regex pattern does not compile: error parsing regexp: .*`
func malformedPattern(ctx context.Context) error {
	return nil
}
