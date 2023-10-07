package dob

import "time"

// func to parse dateofbirth
func ParseDOB(dobstring string) time.Time {
	dob, err := time.Parse("2006-01-02", dobstring)
	if err != nil {
		panic(err)
	}
	return dob
}
