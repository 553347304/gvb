package config

type ModelOption struct {
	Label   string `yaml:"label" json:"label"`
	Value   string `yaml:"value" json:"value"`
	Disable bool   `yaml:"disable" json:"disable"`
}

type Setting struct {
	Name      string `yaml:"name" json:"name"`
	Enable    bool   `yaml:"enable" json:"enable"`
	Order     int    `yaml:"order" json:"order"` // 菜单序号
	ApiKey    string `yaml:"api_key" json:"api_key"`
	ApiSecret string `yaml:"api_secret" json:"api_secret"`
	Title     string `yaml:"title" json:"title"`
	Prompt    string `yaml:"prompt" json:"prompt"` // 加强提示词
	Slogan    string `yaml:"slogan" json:"slogan"` // slogan
	Help      string `json:"help"`
}
type SessionSetting struct {
	ChatScore    int `yaml:"chat_scope" json:"chat_scope"`       // 对话积分消耗
	SessionScore int `yaml:"session_score" json:"session_score"` // 会话积分消耗
	DayScore     int `yaml:"day_score" json:"day_score"`         // 每日赠送积分
}

type BigModel struct {
	Setting   *Setting        `yaml:"setting" json:"setting"`
	Session   *SessionSetting `yaml:"session" json:"session"`
	ModelList *[]ModelOption  `yaml:"model_list" json:"model_list"`
}
