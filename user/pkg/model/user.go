package model

import (
	db "github.com/869413421/micro-service/common/pkg/db"
	pb "github.com/869413421/micro-service/user/proto/user"
)

// User 用户模型
type User struct {
	db.BaseModel
	Name     string `gorm:"column:name;type:varchar(255);not null;unique;default:''" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(255) not null;unique;default:''" valid:"email"`
	RealName string `gorm:"column:real_name;type:varchar(255);not null;default:''" valid:"realName"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null;default:''" valid:"avatar"`
	Status   int    `gorm:"column:status;type:tinyint(1);not null;default:0" `
	Password string `gorm:"column:password;type:varchar(255) not null;;default:''" valid:"password"`
}

// ToORM protobuf转换为orm
func ToORM(protoUser *pb.User) *User {
	user := &User{}
	user.ID = protoUser.Id
	user.Email = protoUser.Email
	user.Name = protoUser.Name
	user.Avatar = protoUser.Avatar
	user.RealName = protoUser.RealName
	return user
}

// ToProtobuf orm转换为protobuf
func (model *User) ToProtobuf() *pb.User {
	user := &pb.User{}
	user.Id = model.ID
	user.Email = model.Email
	user.Name = model.Name
	user.Avatar = model.Avatar
	user.CreateAt = model.CreatedAtDate()
	user.UpdateAt = model.UpdatedAtDate()
	user.RealName = model.RealName
	return user
}

// Store 创建用户
func (model *User) Store() (err error) {
	result := db.GetDB().Create(&model)
	err = result.Error
	if err != nil {
		return err
	}
	return nil
}

// Update 更新用户
func (model *User) Update() (rowsAffected int64, err error) {
	result := db.GetDB().Save(&model)
	err = result.Error
	if err != nil {
		return 0, err
	}
	rowsAffected = result.RowsAffected
	return
}

// Delete 删除用户
func (model User) Delete() (rowsAffected int64, err error) {
	result := db.GetDB().Delete(&model)
	err = result.Error
	if err != nil {
		return
	}
	rowsAffected = result.RowsAffected
	return
}