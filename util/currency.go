package util

const (
	TWD = "TWD"
	JPY = "JPY"
	USD = "USD"
)

func IsSuportedCurrency(currency string) bool {
	switch currency {
	case TWD, JPY, USD:
		return true
	default:
		return false
	}
}
