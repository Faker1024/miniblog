package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/miniblog/internal/pkg/core"
	"github.com/marmotedu/miniblog/internal/pkg/errno"
	"github.com/marmotedu/miniblog/internal/pkg/log"
	v1 "github.com/marmotedu/miniblog/internal/pkg/miniblog/v1"
)

// Create 创建一个新用户
func (ctrl *UserController) Create(ctx *gin.Context) {
	log.C(ctx).Infow("Create user function called")
	var r v1.CreateUserRequest
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}
	err = ctrl.b.User().Create(ctx, &r)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, errno.OK, nil)
}
