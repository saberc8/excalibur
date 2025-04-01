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
POST /api/user/changePassword

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
