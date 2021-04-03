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

 //CategoriesTotal возвращает сумму платежей по каждой категории
 func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money  {
	 categories:= map[types.Category]types.Money{}
	 for _, payment := range payments {
		 categories[payment.Category]+=payment.Amount
	 }

	 return categories
 }

 func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money  {
	 categories:= map[types.Category]types.Money{}
	 paymentsInCategory:=map[types.Category]int{}
	for _, payment := range payments {		
		categories[payment.Category]+=payment.Amount
		paymentsInCategory[payment.Category]++
	}

	for key, category := range categories {
		categories[key]=category/types.Money(paymentsInCategory[key])
	}

	return categories
 }

 func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money  {
	dynamic:= map[types.Category]types.Money{} 
		for scndCategory, scndPay := range second {
			dynamic[scndCategory]=scndPay-first[scndCategory]			
	 	}
	 for frstCategory, frstPay := range first {
		 for _, _ = range dynamic {
			_, ok:=dynamic[frstCategory]
			if !ok{
				dynamic[frstCategory]=0-frstPay
			}
		 }
				
			}	
	return dynamic	 
 }