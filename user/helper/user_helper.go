package helper

import (
	db "github.com/rabit/database"
	dbModel "github.com/rabit/database/models"
	userModel "github.com/rabit/user/models"
	//"fmt"
	commonUtils "github.com/rabit/common"
)

func ListUsers() (error, []dbModel.User){
	err,users := db.GetUserDetail()
	return err,users
}

func AuthenticateUser(requestData userModel.LoginDetail) userModel.LoginResponse{
	var response userModel.LoginResponse
	response.Status = "Error"
	e,u := db.GetUserFromEmail(requestData.Email)
	if e != nil {
		response.Error = e.Error()
		return response
	}
	if u.Email == "" {
		response.Error = "User Not Exist"
		return response
	}

	uniqueId,err := commonUtils.NewUUID()
	if err != nil {
		response.Error = err.Error()
		return response

	}
if err != nil {
		response.Error = err.Error()
		return response

	}
	err = db.SetKey(uniqueId,u.UserName,100000)
	response.Status = "OK"
	response.UniqueKey = uniqueId
	return response
}

