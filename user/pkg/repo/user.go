package repo

import (
	baseDb "github.com/869413421/micro-service/common/pkg/db"
	"github.com/869413421/micro-service/common/pkg/pagination"
	"github.com/869413421/micro-service/user/pkg/model"
	"gorm.io/gorm"
)

// UserRepositoryInterface 用户CURD仓库接口
type UserRepositoryInterface interface {
	GetFirst(where map[string]interface{}) (*model.User, error)
	GetByID(uint642 uint64) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Pagination(page uint64, perPage uint32) (users []model.User, viewData pagination.ViewData, err error)
}

// UserRepository 用户仓库
type UserRepository struct {
	Db *gorm.DB
}

// NewUserRepository 初始化仓库
func NewUserRepository() UserRepositoryInterface {
	db := baseDb.GetDB()
	return &UserRepository{Db: db}
}

// GetByID 根据ID获取用户
func (repo UserRepository) GetByID(id uint64) (*model.User, error) {
	user := &model.User{}
	err := repo.Db.First(&user, id).Error
	return user, err
}

// Pagination 获取分页数据
func (repo UserRepository) Pagination(page uint64, perPage uint32) (users []model.User, viewData pagination.ViewData, err error) {
	//1.初始化分页实例
	DB := repo.Db.Model(model.User{}).Order("created_at desc")
	_pager := pagination.New(DB, page, perPage)

	// 2. 获取分页构建数据
	viewData = _pager.Paging()

	// 3. 获取数据
	_pager.Results(&users)

	return users, viewData, nil
}

// GetByEmail 根据email获取用户
func (repo UserRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := repo.Db.Where("email = ?", email).First(&user).Error
	return user, err
}

// GetFirst 根据自定义条件获取用户
func (repo UserRepository) GetFirst(where map[string]interface{}) (*model.User, error) {
	user := &model.User{}
	for key, val := range where {
		repo.Db.Where(key+"=?", val)
	}
	err := repo.Db.First(&user).Error
	return user, err
}

