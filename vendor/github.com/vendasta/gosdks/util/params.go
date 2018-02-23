package util

import "fmt"

//requireExclusiveProperties checks to make sure only one of the possible parameters is present
func RequireExclusiveProperties(n []string) error {
	found := 0
	for _, v := range n {
		if v != "" {
			found++
		}
	}
	if found != 1 {
		return fmt.Errorf("multiple exclusive properties found %+v", n)
	}
	return nil
}
