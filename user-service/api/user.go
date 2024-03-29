package apiUser

import (
	"context"
	"encoding/json"
	"reflect"
	"user-service/internal/model"
	protoCommon "user-service/proto/common"
	protoUser "user-service/proto/user"

	orm "github.com/hadanhtuan/go-sdk/db/orm"

	"github.com/hadanhtuan/go-sdk/common"
	"gorm.io/gorm"
)

type UserController struct {
	protoUser.UnimplementedUserServiceServer
}

func (pc *UserController) Login(ctx context.Context, req *protoUser.MessageLogin) (*protoCommon.BaseResponse, error) {
	user := &model.User{
		Username: req.Username,
		Password: req.Password,
	}

	result := QueryOne(model.UserDB, user)

	data, err := json.Marshal(result.Data)
	if err != nil {
	}
	baseResponse := &protoCommon.BaseResponse{
		Status:  result.Status,
		Message: result.Message,
		Data:    string(data),
		Total:   result.Total,
	}

	return baseResponse, nil
}

func Create(db *gorm.DB, entity interface{}) *common.APIResponse {
	// check table

	err := db.WithContext(context.TODO()).Create(entity).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.Created,
			Message: "Cannot create item in table. Error detail: " + err.Error(),
		}
	}

	return &common.APIResponse{
		Status: common.APIStatus.Created,
		Data:   entity,
	}
}

func QueryOne(i *orm.Instance, params interface{}) *common.APIResponse {

	// check table
	if i.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + " is not init.",
		}
	}

	reflect.TypeOf(i.Model)
	var entity any
	err := i.DB.WithContext(context.TODO()).Where(params).First(entity).Error

	if entity == nil || err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "Not found any matched. Error detail: " + err.Error(),
		}
	}

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Data:    []interface{}{entity},
		Message: "Query " + " successfully.",
	}
}

func Query(i *orm.Instance, params interface{}, offset int, limit int) *common.APIResponse {
	var entities []interface{}
	var total int64

	err := i.DB.WithContext(context.TODO()).Model(params).Where(params).Count(&total).Error
	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Cannot count item in table " + ". Error detail: " + err.Error(),
		}
	}
	err = i.DB.WithContext(context.TODO()).Offset((offset - 1) * limit).Limit(limit).Where(params).Find(entities).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Cannot find item in table " + ". Error detail: " + err.Error(),
		}
	}

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Data:    entities,
		Total:   total,
		Message: "Query " + " successfully.",
	}
}
