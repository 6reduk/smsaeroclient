package stub

import "fmt"

func GetErrorStub() error {
	return fmt.Errorf("%s", RandomDefaultStr())
}
