package local

import "database/sql/driver"

type Currency string

const (
	CurrencyEUR Currency = "EUR"
	CurrencyUSD Currency = "USD"
)

func (c Currency) String() string {
	return string(c)
}
func (c *Currency) Scan(value interface{}) error {
	*c = Currency(value.(string))
	return nil
}
func (c Currency) Value() (driver.Value, error) {
	return string(c), nil
}
