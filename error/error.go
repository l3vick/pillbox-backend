package error

import (
	"fmt"

	"github.com/l3vick/go-pharmacy/model"
	
	"github.com/go-sql-driver/mysql"
)

const Insert string = "Insert"
const Update string = "Update"
const Delete string = "Delete"

func HandleMysqlError(err error) model.RequestResponse {
	var response model.RequestResponse
	me, ok := err.(*mysql.MySQLError)
	if ok {
		response.Code = 405
		response.Message = fmt.Sprintf("Myqsl Error: Type: %d.  Message: %s ", me.Number, me.Message)
	}
	return response
}

func HandleEmptyRowsError(rowsAffected int64, queryType string, title string) model.RequestResponse{
	var response model.RequestResponse	
		if (rowsAffected == 0){
			response.Code = 404
			response.Message = fmt.Sprintf("%s %s error", queryType, title)	
		} else {
			response.Code = 201
			response.Message = fmt.Sprintf("%s %s success", queryType, title)		
		}
	return response
}