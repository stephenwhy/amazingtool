package decimal

import "github.com/shopspring/decimal"

// Sub 浮点金额减法
func Sub(a, b float64) float64 {
	da := decimal.NewFromFloat(a)
	db := decimal.NewFromFloat(b)
	res, _ := da.Sub(db).Round(2).Float64()
	return res
}

// Add 浮点金额加法
func Add(num ...float64) float64 {
	var sum = decimal.Zero
	for _, v := range num {
		dv := decimal.NewFromFloat(v)
		sum = sum.Add(dv)
	}
	res, _ := sum.Round(2).Float64()
	return res
}

// Mul 浮点乘法
func Mul(a, b float64) float64 {
	da := decimal.NewFromFloat(a)
	db := decimal.NewFromFloat(b)
	res, _ := da.Mul(db).Round(2).Float64()
	return res
}

// Div 浮点除法
func Div(a, b float64, places int) float64 {
	if b == 0 {
		return 0
	}
	da := decimal.NewFromFloat(a)
	db := decimal.NewFromFloat(b)
	res, _ := da.Div(db).Round(int32(places)).Float64()
	return res
}
