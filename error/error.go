package error

import (
	"fmt"

	"github.com/l3vick/go-pharmacy/model"
	"database/sql"
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

func HandleEmptyRowsError(result sql.Result, queryType string) model.RequestResponse{
	var response model.RequestResponse
	count, err := result.RowsAffected()  
	if err != nil {  
		response.Code = 405
		response.Message = fmt.Sprintf("Mysql Error: %s", err.Error())		
	} else {		
		if (count == 0){
			response.Code = 404
			response.Message = fmt.Sprintf("%s Treatment error", queryType)		
		} else {
			response.Code = 201
			response.Message = fmt.Sprintf("%s Treatment success", queryType)		
		}
	}
	return response
}