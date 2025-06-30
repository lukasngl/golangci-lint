//golangcitest:args -Egodogen
package godogs

import (
	"context"
)

//godogen:given correct (number) of groups
// want pattern has 2 groups, but function has 1 regular parameters
//godogen:given (too) (many) groups
// want pattern has 0 groups, but function has 1 regular parameters
//godogen:given not enough of groups
func paramGroupMismatch(ctx context.Context, number string) error {
	return nil
}

// want parameter "dogs" has unexpected type "rune" (godogen)
// pattern has 1 groups, but function has 0 regular parameters
//godogen:given wrong (parameter) type
func wrongParameterType(ctx context.Context, dogs rune) error {
	return nil
}

// want function should return (context.Context, error) with two return values
//godogen:then we have a wrong return type
func wrongReturnType() (string, error) {
	return "", nil
}

// want regex pattern does not compile: error parsing regexp: unexpected ): `)malformed pattern(
//godogen:step )malformed pattern(
func malformedPattern(ctx context.Context) error {
	return nil
}
