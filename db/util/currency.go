package util

var supportedCurrencies = [4]string{"GBP", "G", "S", "K"}

func IsSupportedCurrency(currency string) bool {
	for _, supportedCurrency := range supportedCurrencies {
		if currency == supportedCurrency {
			return true
		}
	}
	return false
}
