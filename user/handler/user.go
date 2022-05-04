package handler

import (
	"context"
	"github.com/869413421/micro-service/common/pkg/types"
	"github.com/869413421/micro-service/user/pkg/model"
	"github.com/869413421/micro-service/user/pkg/repo"
	pb "github.com/869413421/micro-service/user/proto/user"
	"github.com/869413421/micro-service/user/service"
	"github.com/micro/go-micro/v2/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

//UserServiceHandler 用户服务处理器
type UserServiceHandler struct {
	UserRepo     repo.UserRepositoryInterface
	TokenService service.Authble
}

// NewUserServiceHandler 用户服务初始化
func NewUserServiceHandler() *UserServiceHandler {
	return &UserServiceHandler{
		UserRepo:     repo.NewUserRepository(),
		TokenService: service.NewTokenService(),
	}
}

// Pagination 分页
func (srv *UserServiceHandler) Pagination(ctx context.Context, req *pb.PaginationRequest, rsp *pb.PaginationResponse) error {
	// 1.查找分页数据
	users, pagerData, err := srv.UserRepo.Pagination(req.Page, req.PerPage)
	if err != nil {
		return errors.InternalServerError("user.Pagination.Pagination.Error", err.Error())
	}

	// 2.构造用户列表
	userItems := make([]*pb.User, len(users))
	for index, user := range users {
		userItem := user.ToProtobuf()
		userItems[index] = userItem
	}

	// 3.返回用户信息
	rsp.Users = userItems
	rsp.Total = pagerData.TotalCount
	return nil
}

// Get 根据ID获取数据
func (srv *UserServiceHandler) Get(ctx context.Context, req *pb.GetRequest, rsp *pb.UserResponse) error {
	// 1.查找用户
	user, err := srv.UserRepo.GetByID(req.GetId())
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return errors.BadRequest("User.GetByID", "user not found")
	}

	// 2.返回用户信息
	rsp.User = user.ToProtobuf()
	return nil
}

// Create 创建用户
func (srv *UserServiceHandler) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.UserResponse) error {
	// 1.填充提交信息
	user := &model.User{}
	types.Fill(user, req)

	// 2.创建用户
	err := user.Store()
	if err != nil {
		return err
	}

	// 3.返回用户信息
	rsp.User = user.ToProtobuf()
	return nil
}

// Update 更新用户信息
func (srv *UserServiceHandler) Update(ctx context.Context, req *pb.UpdateRequest, rsp *pb.UserResponse) error {
	// 1.获取用户
	id := req.Id
	_user, err := srv.UserRepo.GetByID(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return errors.NotFound("User.Update.GetUserByID.Error", "user not found ,check you request id")
	}

	// 2.验证提交信息
	types.Fill(_user, req)

	// 3.更新用户
	rowsAffected, err := _user.Update()
	if rowsAffected == 0 || err != nil {
		return errors.InternalServerError("User.Update.Update.Error", err.Error())
	}

	// 4.返回更新信息
	rsp.User = _user.ToProtobuf()
	return nil
}

// Delete 删除用户
func (srv *UserServiceHandler) Delete(ctx context.Context, req *pb.DeleteRequest, rsp *pb.UserResponse) error {
	// 1.获取用户
	id := req.Id
	_user, err := srv.UserRepo.GetByID(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return errors.NotFound("User.Delete.GetUserByID.Error", "user not found ,check you request id")
	}

	// 2.删除用户
	rowsAffected, err := _user.Delete()
	if err != nil {
		return errors.InternalServerError("User.Delete.Delete.Error", err.Error())
	}
	if rowsAffected == 0 {
		return errors.BadRequest("User.Delete.Delete.Fail", "update fail")
	}

	// 3.返回更新信息
	rsp.User = _user.ToProtobuf()
	return nil
}

// Auth 认证获取token
func (srv UserServiceHandler) Auth(ctx context.Context, req *pb.AuthRequest, rsp *pb.TokenResponse) error {
	// 1.根据邮箱回去用户
	log.Println("Logging in with:", req.Email, req.Password)
	user, err := srv.UserRepo.GetByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return errors.NotFound("User.Auth.GetByEmail.Error", "user not found ,check you password or email")
	}

	// 2.检验用户密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return errors.Unauthorized("User.Auth.CheckPassword.Error", err.Error())
	}

	// 3.生成token
	token, err := srv.TokenService.Encode(user)
	if err != nil {
		return err
	}

	// 4.返回token字符串
	rsp.Token = token
	return nil
}

// ValidateToken 验证token
func (srv *UserServiceHandler) ValidateToken(ctx context.Context, req *pb.TokenRequest, rsp *pb.TokenResponse) error {
	// 1.将token字符串转换为token对象
	claims, err := srv.TokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	// 2.判断转换是否成功
	if claims.User.ID == 0 {
		return errors.BadRequest("User.ValidateToken.Error", "user invalid")
	}

	// 3.返回验证状态
	rsp.Token = req.Token
	rsp.Valid = true
	return nil
}
