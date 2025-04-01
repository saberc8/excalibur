# rule: 除了获取列表相关接口使用 GET 请求外，其余接口均使用 POST 请求

POST /api/user/create # 创建用户
GET /api/user # 获取单个用户信息，可通过请求参数指定用户 ID
GET /api/user/list # 获取用户列表
POST /api/user/update # 更新用户信息
POST /api/user/delete # 删除用户
POST /api/user/bindRole # 为用户绑定角色
POST /api/user/unbindRole # 为用户解绑角色
GET /api/user/roles # 获取用户拥有的角色列表
GET /api/user/menus # 获取用户可访问的菜单列表
POST /api/user/changePassword # 修改用户密码

POST /api/role/create # 创建角色
GET /api/role # 获取单个角色信息，可通过请求参数指定角色 ID
GET /api/role/list # 获取角色列表
POST /api/role/update # 更新角色信息
POST /api/role/delete # 删除角色
POST /api/role/bindMenu # 为角色绑定菜单
POST /api/role/unbindMenu # 为角色解绑菜单
GET /api/role/menus # 获取角色拥有的菜单列表
GET /api/role/users # 获取拥有该角色的用户列表

POST /api/menu/create # 创建菜单
GET /api/menu # 获取单个菜单信息，可通过请求参数指定菜单 ID
GET /api/menu/list # 获取菜单列表
POST /api/menu/update # 更新菜单信息
POST /api/menu/delete # 删除菜单
GET /api/menu/roles # 获取拥有该菜单访问权限的角色列表
GET /api/menu/tree # 获取菜单的树形结构

type RoleEntity struct {
	global.GModel
	Name   string `gorm:"column:name;type:varchar(50);comment:名称" json:"name"`
	Remark string `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
	Status int64  `gorm:"column:status;type:smallint;comment:状态" json:"status"` // 0 正常 1 禁用

}

type MenuEntity struct {
	global.GModel
	Name      string `gorm:"column:name;type:varchar(50);comment:菜单名称" json:"name"`
	ParentId  int64  `gorm:"column:parent_id;type:int;comment:父级菜单ID" json:"parentId"`
	OrderNum  int64  `gorm:"column:order_num;type:int;comment:排序" json:"orderNum"`
	Path      string `gorm:"column:path;type:varchar(100);comment:路径" json:"path"`
	Component string `gorm:"column:component;type:varchar(100);comment:组件" json:"component"`
	Query     string `gorm:"column:query;type:varchar(100);comment:参数" json:"query"`
	IsFrame   int64  `gorm:"column:is_frame;type:smallint;comment:是否外链" json:"isFrame"`
	MenuType  string `gorm:"column:menu_type;type:varchar(100);comment:菜单类型" json:"menuType"` // 菜单类型: C目录 M菜单 B按钮
	IsCatch   int64  `gorm:"column:is_catch;type:smallint;comment:是否缓存" json:"isCatch"`
	IsHidden  int64  `gorm:"column:is_hidden;type:smallint;comment:是否可见" json:"isHidden"`
	Perms     string `gorm:"column:perms;type:varchar(100);comment:权限标识" json:"perms"`
	Icon      string `gorm:"column:icon;type:varchar(100);comment:图标" json:"icon"`
	Status    int64  `gorm:"column:status;type:smallint;comment:状态" json:"status"` // 0 正常 1 禁用
	Remark    string `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
}

type UserEntity struct {
	global.GModel
	Username          string    `gorm:"column:username;type:varchar(50);comment:账号" json:"username"`                                                        // 账号
	Password          string    `gorm:"column:password;type:varchar(100);not null;comment:密码" json:"password"`                                              // 密码
	Avatar            string    `gorm:"column:avatar;type:varchar(100);comment:头像" json:"avatar"`                                                           // 头像
	Nickname          string    `gorm:"column:nickname;type:varchar(50);comment:昵称" json:"nickname"`                                                        // 昵称
	UserType          int64     `gorm:"column:user_type;type:smallint;comment:账号类型:0普通账号,1是超管" json:"userType"`                                             // 账号类型:0普通账号,1是超管
	Email             string    `gorm:"column:email;type:varchar(50);comment:电邮地址" json:"email"`                                                            // 电邮地址
	Mobile            string    `gorm:"column:mobile;type:varchar(30);comment:手机号码" json:"mobile"`                                                          // 手机号码
	Sort              int64     `gorm:"column:sort;type:int;default:1;comment:排序" json:"sort"`                                                              // 排序
	Status            int64     `gorm:"column:status;type:smallint;comment:状态0是正常,1是禁用" json:"status"`                                                      // 状态0是正常,1是禁用
	LastLoginIP       string    `gorm:"column:last_login_ip;type:varchar(30);comment:最后登录ip地址" json:"lastLoginIp"`                                          // 最后登录ip地址
	LastLoginNation   string    `gorm:"column:last_login_nation;type:varchar(100);comment:最后登录国家" json:"lastLoginNation"`                                   // 最后登录国家
	LastLoginProvince string    `gorm:"column:last_login_province;type:varchar(100);comment:最后登录省份" json:"lastLoginProvince"`                               // 最后登录省份
	LastLoginCity     string    `gorm:"column:last_login_city;type:varchar(100);comment:最后登录城市" json:"lastLoginCity"`                                       // 最后登录城市
	LastLoginDate     time.Time `gorm:"column:last_login_date;type:timestamp(6);not null;default:CURRENT_TIMESTAMP(6);comment:最后登录时间" json:"lastLoginDate"` // 最后登录时间
	Salt              string    `gorm:"column:salt;type:varchar(30);comment:密码盐" json:"salt"`                                                               // 密码盐
}


