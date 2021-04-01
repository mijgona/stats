package stats

import (
	"github.com/mijgona/bank/pkg/types"
)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money  {
	sum:=types.Money(0)
	for _, payment := range payments {
		sum+=payment.Amount
	}
	return sum/types.Money(len(payments))
}

//TotalInCategory находит сумму покупок в определенной категории
 func TotalInCategory(payments []types.Payment, catergory types.Category) types.Money  {
	sum:=types.Money(0) 
	for _, payment := range payments {
		 if payment.Category==catergory{
			sum+=payment.Amount
		 }
	 }
	 return sum
 }