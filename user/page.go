package user

import "fmt"

type pageForm struct {
	currentPage int
	orderType	string
	orderCol	string
	pageFlag	bool
	showRecord	int
}

func (pageForm *pageForm) toString()  string {
	return fmt.Sprintf("pageForm.currentPage=%d&pageForm.orderType=%s&pageForm.orderCol=%s&pageForm.pageFlag=%t&pageForm.showRecord=%d",
		pageForm.currentPage,pageForm.orderType,pageForm.orderCol,pageForm.pageFlag,pageForm.showRecord)
}

