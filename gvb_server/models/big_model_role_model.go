package models

// BigModelRoleModel 大模型角色表
type BigModelRoleModel struct {
	MODEL
	Name      string             `gorm:"size:16" json:"name"`
	Enable    bool               `json:"enable"`
	Icon      string             `gorm:"size:256" json:"icon"`                                                                                                         // 系统默认   自己上传
	Abstract  string             `gorm:"size:256" json:"abstract"`                                                                                                     // 角色简介
	Tags      []BigModelTagModel `gorm:"many2many:big_model_role_tag_models;joinForeignKey:big_model_role_model_id;JoinReferences:big_model_tag_model_id" json:"tags"` // 角色标签列表
	Scope     int                `json:"scope"`                                                                                                                        // 消耗积分
	Prologue  string             `gorm:"size:512" json:"prologue"`                                                                                                     // 开场白
	Prompt    string             `gorm:"size:2048" json:"prompt"`                                                                                                      // 设定词
	AutoReply bool               `json:"auto_reply"`                                                                                                                   // 是否接入自动回复
}

// BigModelTagModel 大模型标签表
type BigModelTagModel struct {
	MODEL
	Title string              `gorm:"size:16" json:"title"`                                                                                                          // 标题
	Color string              `gorm:"size:16" json:"color"`                                                                                                          // 颜色
	Roles []BigModelRoleModel `gorm:"many2many:big_model_role_tag_models;joinForeignKey:big_model_tag_model_id;JoinReferences:big_model_role_model_id" json:"roles"` // 角色列表
}

// BigModelUserRoleModel 用户选择大模型角色表
type BigModelUserRoleModel struct {
	MODEL
	UserId    uint              `json:"user_id"` // 用户ID
	RoleId    uint              `json:"role_id"` // 角色ID
	RoleModel BigModelRoleModel `gorm:"foreignKey:RoleId" json:"-"`
}

// BigModelSessionModel 大模型会话表
type BigModelSessionModel struct {
	MODEL
	Name      string              `gorm:"size:32" json:"name"` // 会话名称
	UserId    uint                `json:"user_id"`             // 用户ID
	UserModel UserModel           `gorm:"foreignKey:UserId" json:"-"`
	RoleId    uint                `json:"role_id"` // 角色ID
	RoleModel BigModelRoleModel   `gorm:"foreignKey:RoleId" json:"-"`
	ChatList  []BigModelChatModel `gorm:"foreignKey:SessionId"` // 会话列表
}

// BigModelChatModel 大模型对话表
type BigModelChatModel struct {
	MODEL
	SessionId    uint                 `json:"session_id"` // 会话ID
	SessionModel BigModelSessionModel `gorm:"foreignKey:SessionId"`
	Status       bool                 `json:"status"`      // 状态
	Content      string               `json:"content"`     // 用户聊天内容
	BotContent   string               `json:"bot_content"` // AI回复
	UserId       uint                 `json:"user_id"`     // 用户ID
	UserModel    UserModel            `gorm:"foreignKey:UserId" json:"-"`
	RoleId       uint                 `json:"role_id"` // 角色ID
	RoleModel    BigModelRoleModel    `gorm:"foreignKey:RoleId" json:"-"`
}
