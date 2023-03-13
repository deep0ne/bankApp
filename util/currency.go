package util

// константы для поддерживаемой валюты
const (
	USD = "USD"
	EUR = "EUR"
	RUB = "RUB"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, RUB:
		return true
	}
	return false
}
