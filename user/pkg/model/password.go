package model

import (
	db "github.com/869413421/micro-service/common/pkg/db"
)

// PasswordReset 重置密码模型
type PasswordReset struct {
	db.BaseModel
	Token  string `gorm:"column:token;type:varchar(255) not null;index" `
	Email  string `gorm:"column:email;type:varchar(255) not null;index" valid:"email"`
	Verify int8   `gorm:"column:verify;type:tinyint(1);not null;default:0"`
}

// Store 创建重置记录
func (model *PasswordReset) Store() (err error) {
	result := db.GetDB().Create(&model)
	err = result.Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除数据库重置记录
func (model *PasswordReset) Delete() (rowsAffected int64, err error) {
	result := db.GetDB().Delete(&model)
	err = result.Error
	if err != nil {
		return 0, err
	}
	rowsAffected = result.RowsAffected
	return rowsAffected, nil
}

// Update 更新数据库重置记录
func (model *PasswordReset) Update() (rowsAffected int64, err error) {
	result := db.GetDB().Save(&model)
	err = result.Error
	if err != nil {
		return 0, err
	}
	rowsAffected = result.RowsAffected
	return rowsAffected, nil
}
