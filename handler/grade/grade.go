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
	Xnm      string `json:"xnm"`      // 学年名称
	JxbID    string `json:"jxb_id"`   // 教学班 id
}

// GradeDetailRequest ... 成绩详情请求参数
type GradeDetailRequest struct {
	Course  string `form:"course"`                      // 课程名
	ClassID string `form:"class_id" binding:"required"` // 教学班 id
	Year    string `form:"year" binding:"required"`     // 学年
	Term    string `form:"term" binding:"required"`     // 学期
}

// GradeDetailResponse ... 成绩详情响应数据
type GradeDetailResponse struct {
	Course          string `json:"course"`
	UsualGrade      string `json:"usual_grade"`      // 平时成绩
	FinalGrade      string `json:"final_grade"`      // 期末成绩
	UsualPercentage string `json:"usual_percentage"` // 平时分占比 %
	FinalPercentage string `json:"final_percentage"` // 期末分占比 %
}
