package grade

type Reqeust struct {
	XNM string `form:"xnm"`
	XQM string `form:"xqm"`
}

type GradeItem struct {
	Course   string `json:"course"`   // 课程名称
	Credit   string `json:"credit"`   // 学分
	Grade    string `json:"grade"`    // 成绩
	Category string `json:"category"` // 课程类别名称，比如专业课/公共课
	Type     string `json:"type"`     // 课程归属名称，比如文/理
	Kcxzmc   string `json:"kcxzmc"`   // 课程性质名称，比如专业主干课程/通识必修课
}
