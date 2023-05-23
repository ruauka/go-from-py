package logic

// Result - структура выходных агрегатов.
type Result struct {
	// Признак достаточности денег для ежемесячного платежа в дату платежа в разбивке по месяцам
	// Формат: 1 - достаточно , 0 - не достаточно
	EnoughMoneyByMonths [6]int `json:"enough_money_by_months"`
	// Признак просрочки платежа в разбивке по месяцам
	// Формат: 1 - есть просрочка, 0 - нет
	DelinquencyByMonths [6]int `json:"delinquency_by_months"`
	// Количество дней в просрочке за весь период
	DelinquencyDurationDaysTotal int `json:"delinquency_duration_days_total"`
	// Итоговая сумма просрочек за весь период
	DelinquencySumTotal int `json:"delinquency_sum_total"`
}

// EnoughMoneyFinal - Расчет итоговых агрегатов достаточности денег.
func (r *Result) EnoughMoneyFinal(data *Data) {
	for index, flag := range data.Delinquency.EnoughMoney {
		if flag {
			r.EnoughMoneyByMonths[index] = 1
		} else {
			r.EnoughMoneyByMonths[index] = 0
		}
	}
}

// DelinquencyFinal - Расчет итоговых агрегатов просрочки.
func (r *Result) DelinquencyFinal(data *Data) {
	for index, flag := range data.Delinquency.Delinquency {
		if flag {
			r.DelinquencyByMonths[index] = 1
		} else {
			r.DelinquencyByMonths[index] = 0
		}
	}
}

// DelinquencyDurationFinal - Расчет длительности просрочек в днях по месяцам.
// По сути, это пример ненужной перекладки.
func (r *Result) DelinquencyDurationFinal(data *Data) {
	for _, value := range data.Delinquency.DelinquencyDurationDays {
		r.DelinquencyDurationDaysTotal += value
	}
}

// DelinquencySumTotalCount - Расчет итоговой суммы случавшихся просрочек за весь период.
func (r *Result) DelinquencySumTotalCount(data *Data) {
	for _, sum := range data.Delinquency.DelinquencySum {
		r.DelinquencySumTotal += sum
	}
}
