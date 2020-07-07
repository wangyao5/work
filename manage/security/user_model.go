package security

import (
	"manage/middle"
	"manage/util"
	"time"
)

//UserModel 用户模型
type UserModel struct {
	ID        int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;comment:'主键'"`
	Username  string    `gorm:"type:varchar(50);column:username;unique_index;comment:'用户名'"`
	Password  string    `gorm:"type:varchar(64);column:password;comment:'密码'"`
	Salt      string    `gorm:"type:varchar(32);column:salt;comment:'安全符'"`
	RealName  string    `gorm:"type:varchar(100);column:realname;comment:'真实姓名'"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now();column:created_at;comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now();column:updated_at;comment:'更新时间'"`
	Mobile    string    `gorm:"type:varchar(20);column:mobile;default:null;comment:'手机号'"`
	Email     string    `gorm:"type:varchar(50);column:email;default:null;comment:'电子邮箱'"`
	Gender    int       `gorm:"type:varchar(50);column:gender;default:null;comment:'性别,NULL 未选择 1 男 2 女'"`
	Status    int       `gorm:"type:int(3);default:1;column:status;default:1;comment:'状态,0禁用,1正常,2未激活'"`
	Role      int       `gorm:"type:int(11);default:2;column:role;default:2;comment:'角色,1管理员,2推广账号'"`
}

//TableName of user.
func (UserModel) TableName() string {
	return "user"
}

//FindOne 查到指定账号.
func (UserModel) FindOne(username string) *UserModel {
	userModel := new(UserModel)
	db := middle.GormDb
	db.Where("username=?", username).First(userModel)
	return userModel
}

//RoleModel 角色模型
type RoleModel struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT;column:id;comment:'主键'"`
	RoleName  string    `gorm:"type:varchar(32);column:role_name;comment:'角色名称'"`
	RoleType  int       `gorm:"type:int(11);column:role_type;comment:'1、平台管理员 2、推广员'"`
	Remark    string    `gorm:"type:varchar(100);column:remark;comment:'备注'"`
	Status    int       `gorm:"type:int(3);column:status;default:1;comment:'1 启用 0 禁用'"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now();column:created_at;comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now();column:updated_at;comment:'更新时间'"`
}

//TableName of role.
func (RoleModel) TableName() string {
	return "role"
}

func init() {
	db := middle.GormDb
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8")
	tables := []interface{}{&UserModel{}, &RoleModel{}}
	for _, table := range tables {
		middle.CreateTable(table)
	}

	//初始化角色数据
	roles := make([]RoleModel, 0, 10)
	roles = append(roles, RoleModel{ID: 1, RoleName: "平台管理员", RoleType: 1, Remark: "平台管理员角色", Status: 1})
	roles = append(roles, RoleModel{ID: 2, RoleName: "推广员", RoleType: 2, Remark: "推广员角色", Status: 1})
	rtx := db.Begin()
	for _, role := range roles {
		if db.NewRecord(role) {
			rtx.Create(&role)
		}
	}
	rtx.Commit()

	//初始化管理员数据
	username := "admin"
	passwd := "123456"
	slat := util.RandomString(8)
	newPasswd := util.MD5(passwd + slat)
	admin := UserModel{ID: 1, Username: username, Password: newPasswd, Salt: slat,
		RealName: "系统管理员", Mobile: "18519272342", Email: "82486240@qq.com", Gender: 1, Status: 1, Role: 1}
	if db.NewRecord(admin) {
		atx := db.Begin()
		atx.Create(&admin)
		atx.Commit()
	}
}
