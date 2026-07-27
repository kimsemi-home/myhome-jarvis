package financeconnector

import "fmt"

type myDataHTTPError struct {
	statusCode int
}

func (err myDataHTTPError) Error() string {
	return fmt.Sprintf("mydata auth request returned HTTP %d", err.statusCode)
}
