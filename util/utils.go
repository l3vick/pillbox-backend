package util

import (
	"fmt"
	"github.com/l3vick/go-pharmacy/model"
	"golang.org/x/crypto/bcrypt"
)

func GetPage(count int, intPage int) model.Page {
	var page model.Page
	var index int
	if count > 10 {
		if count%10 == 0 {
			index = 1
		} else {
			index = 0
		}
		if intPage == 0 {
			page.First = 0
			page.Previous = 0
			page.Next = intPage + 1
			page.Last = (count / 10) - index
			page.Count = count
		} else if intPage == (count/10)-index {
			page.First = 0
			page.Previous = intPage - 1
			page.Next = intPage
			page.Last = (count / 10) - index
			page.Count = count
		} else {
			page.First = 0
			page.Previous = intPage - 1
			page.Next = intPage + 1
			page.Last = (count / 10) - index
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

func ByteToBool(b byte) bool {
	if b == 1 {
		return true
	}
	return false
}

func BoolToByte(b bool) byte {
	if b == true {
		return 1
	}
	return 0
}

func CheckErr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func HashPassword(password * string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
