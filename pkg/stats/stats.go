package stats

import (
	"github.com/mijgona/bank/v2/pkg/types"
)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money  {
	sum:=types.Money(0)
	p:=0
	for _, payment := range payments {
		if payment.Status!=types.StatusFail{		
			sum+=payment.Amount
			p++
		}
	}
	return sum/types.Money(p)
}

//TotalInCategory находит сумму покупок в определенной категории
 func TotalInCategory(payments []types.Payment, catergory types.Category) types.Money  {
	sum:=types.Money(0) 
	for _, payment := range payments {
		 if payment.Category==catergory{
			if payment.Status!=types.StatusFail{
				sum+=payment.Amount
			}
		 }
	 }
	 return sum
 }

 //FilteByCategory возвращает платежи в указанной категории
 func FilterByCategory(payments []types.Payment, catergory types.Category) []types.Payment  {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category==catergory{
			filtered=append(filtered, payment)
		}		
	}
	return filtered	 
 }