// Package logic - main logic.
package logic

import (
	"go-app/internal/request"
)

// Data - data dict..
type Data struct {
	// input JSON
	request.Request
	// main logic
	Delinquency
	// Logic of creating output aggregates
	Result
}

// NewData - data dictionary constructor.
func NewData(inputMessage *request.Request) *Data {
	return &Data{
		Request:     *inputMessage,
		Delinquency: *NewLocal(),
	}
}

// LocalCount - call of Delinquency logic.
func (d *Data) LocalCount() {
	d.Delinquency.EnoughMoneyCount(d)
	d.Delinquency.DelinquencyCount(d)
	d.Delinquency.DelinquencyDurationCount(d)
	d.Delinquency.DelinquencySumCount(d)
}

// ResultCount - call of Result logic.
func (d *Data) ResultCount() {
	d.Result.DelinquencyFinal(d)
	d.Result.EnoughMoneyFinal(d)
	d.Result.DelinquencyDurationFinal(d)
	d.Result.DelinquencySumTotalCount(d)
}
