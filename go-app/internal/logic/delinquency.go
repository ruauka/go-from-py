package logic

import "go-app/internal/utils/functions"

// Delinquency - структура для расчета внутренних агрегатов.
type Delinquency struct {
	// Признак достаточности денег для ежемесячного платежа в дату платежа в разбивке по месяцам
	EnoughMoney [6]bool
	// Признак просрочки платежа в разбивке по месяцам
	Delinquency [6]bool
	// Количество дней в просрочке в разбивке по месяцам
	DelinquencyDurationDays [6]int
	// Сумма просрочек в разбивке по месяцам
	DelinquencySum [6]int
}

// NewLocal - Конструктор структуры Delinquency.
func NewLocal() *Delinquency {
	return &Delinquency{}
}

// EnoughMoneyCount - расчет флага достаточности денег для ежемесячного платежа по месяцам.
func (d *Delinquency) EnoughMoneyCount(data *Data) {
	for i := 0; i < len(data.Request.Loan); i++ {
		if *data.Request.Account[i].PaymentDateBalance < data.Request.Loan[i].Payment {
			d.EnoughMoney[i] = false
		} else {
			d.EnoughMoney[i] = true
		}
	}
}

// DelinquencyCount - Расчет количества просрочек по месяцам.
func (d *Delinquency) DelinquencyCount(data *Data) {
	for i := 0; i < len(data.Request.Loan); i++ {
		if data.Request.Loan[i].ActualPaymentDate.After(data.Request.Loan[i].PaymentDate) {
			d.Delinquency[i] = true
		}
	}
}

// DelinquencyDurationCount - Расчет длительности просрочек в днях по месяцам.
func (d *Delinquency) DelinquencyDurationCount(data *Data) {
	for i := 0; i < len(data.Request.Loan); i++ {
		d.DelinquencyDurationDays[i] = functions.DateDiff(
			data.Request.Loan[i].PaymentDate,
			data.Request.Loan[i].ActualPaymentDate,
		)
	}
}

// DelinquencySumCount - Расчет суммы просрочки в разбивке по месяцам.
func (d *Delinquency) DelinquencySumCount(data *Data) {
	for index, value := range d.DelinquencyDurationDays {
		if value != 0 {
			d.DelinquencySum[index] =
				data.Request.Loan[index].Payment - *data.Request.Account[index].PaymentDateBalance
		}
	}
}
