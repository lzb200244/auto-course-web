package request

/*
Created by 斑斑砖 on 2023/9/2.
Description：
	进行校验参数
*/

type Auths struct {
	RoleID     int   `json:"roleID" validate:"required" label:"角色ID"`
	Permission []int `json:"permission" validate:"required" label:"权限ID"`
}

type Auth struct {
	RoleID       int `json:"roleID" validate:"required" label:"角色ID"`
	PermissionID int `json:"permissionID" validate:"required" label:"权限ID"`
}

type Permission struct {
	Name string `json:"name" validate:"required" label:"权限名称"`
}
type Meta struct {
	Title       string `json:"title"  label:"标题"`
	KeepAlive   bool   `json:"keepAlive"  label:"是否缓存"`
	RequireAuth bool   `json:"requireAuth"  label:"是否需要认证"`
}
type Component struct {
	Name      string `json:"name"  validate:"required" label:"组件名称"`
	Path      string `json:"path"   validate:"required" label:"路由路径"`
	Redirect  string `json:"redirect" label:"重定向路径"`
	Component string `json:"component"  label:"组件名称" `
	Meta      Meta   `json:"meta" validate:"required" label:"元信息"`
	Disable   bool   `json:"disable"`
	Role      []int  `json:"role"   validate:"required" label:"权限限制"`
	Parent    uint   `json:"parent"  label:"父级ID"`
}

type Notice struct {
}

type Category struct {
	Name string `json:"name" validate:"required" label:"分类名称"`
	Desc string `json:"desc" validate:"required" label:"分类描述"`
}
