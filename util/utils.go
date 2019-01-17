package util

import (
	"github.com/l3vick/go-pharmacy/model"
)

func GetPage (count int, intPage int)  (model.Page){
	var page model.Page
	var index int
	if (count > 10){
		if (count % 10 == 0){
			index = 1
		}else{
			index = 0
		}
		if intPage == 0 {
			page.First = 0
			page.Previous = 0
			page.Next = intPage+1
			page.Last = (count/10) - index
			page.Count = count
		} else if intPage == (count/10) - index {
			page.First = 0
			page.Previous = intPage -1
			page.Next = intPage
			page.Last = (count/10) - index
			page.Count = count
		} else {
			page.First = 0
			page.Previous = intPage-1
			page.Next = intPage+1
			page.Last = (count/10) - index
			page.Count = count
		}
	} else {
		page.First = 0
		page.Previous = 0
		page.Next = 0
		page.Last = 0
		page.Count = count
	}
	return page		
}