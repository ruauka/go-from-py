package logic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResult_EnoughMoneyFinal(t *testing.T) {
	data := NewTestData()

	TestTable := []struct {
		enoughMoney bool
		expected    int
		testName    string
	}{
		{
			enoughMoney: true,
			expected:    1,
			testName:    "test-1-EnoughMoney=true",
		},
		{
			enoughMoney: false,
			expected:    0,
			testName:    "test-2-EnoughMoney=false",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			data.Delinquency.EnoughMoney[0] = testCase.enoughMoney
			data.EnoughMoneyFinal(data)
			actual := data.Result.EnoughMoneyByMonths[0]
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestResult_DelinquencyFinal(t *testing.T) {
	data := NewTestData()

	TestTable := []struct {
		delinquency bool
		expected    int
		testName    string
	}{
		{
			delinquency: true,
			expected:    1,
			testName:    "test-1-delinquency=true",
		},
		{
			delinquency: false,
			expected:    0,
			testName:    "test-2-delinquency=false",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			data.Delinquency.Delinquency[0] = testCase.delinquency
			data.DelinquencyFinal(data)
			actual := data.Result.DelinquencyByMonths[0]
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestResult_DelinquencyDurationFinal(t *testing.T) {
	data := NewTestData()

	TestTable := []struct {
		delinquencyDurationDays [6]int
		expected                int
		testName                string
	}{
		{
			delinquencyDurationDays: [6]int{0, 1, 0, 1, 0, 1},
			expected:                3,
			testName:                "test-1",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			data.Delinquency.DelinquencyDurationDays = testCase.delinquencyDurationDays
			data.DelinquencyDurationFinal(data)
			actual := data.Result.DelinquencyDurationDaysTotal
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestResult_DelinquencySumTotalCount(t *testing.T) {
	data := NewTestData()

	TestTable := []struct {
		delinquencySum [6]int
		expected       int
		testName       string
	}{
		{
			delinquencySum: [6]int{10, 0, 20, 30, 0, 10},
			expected:       70,
			testName:       "test-1",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			data.Delinquency.DelinquencySum = testCase.delinquencySum
			data.DelinquencySumTotalCount(data)
			actual := data.Result.DelinquencySumTotal
			require.Equal(t, testCase.expected, actual)
		})
	}
}
