package ctype

type Role string

const (
	PermissionAdmin       Role = "管理员"
	PermissionUser        Role = "普通用户"
	PermissionVisitor     Role = "游客"
	PermissionDisableUser Role = "禁用用户"
)
