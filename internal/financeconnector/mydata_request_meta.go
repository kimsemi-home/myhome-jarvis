package financeconnector

import (
	"fmt"
	"time"
)

func (client LiveClient) fromDate() string {
	if client.Config.FromDate != "" {
		return client.Config.FromDate
	}
	return time.Now().Format("20060102")
}

func (client LiveClient) toDate() string {
	if client.Config.ToDate != "" {
		return client.Config.ToDate
	}
	return time.Now().Format("20060102")
}

func transactionID() string {
	return fmt.Sprintf("MHJ%022d", time.Now().UnixNano())
}
