package common

type LoginVo struct {
	Token    string `json:"token"`
	UserInfo UserVo `json:"userInfo"`
}

type UserVo struct {
	ID           int64    `json:"id"`
	Username     string   `json:"username"`
	Nickname     string   `json:"nickname"`
	Avatar       string   `json:"avatar"`
	Introduction string   `json:"introduction"`
	Roles        []string `json:"roles"`
}
