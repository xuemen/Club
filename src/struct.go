// struct.go
package main

type InformationAssets struct {
	Subject   string
	Submitter string
	Content   string
	Sig       string
}

//俱乐部运营计划（含预算计划、高级会员值班表）
//Club operation plan
type COP struct {
}

//Administrate Report
//管理员工作报告（含决算报告）
type AdminReport struct {
}

//Administrate system
//管理制度（含部门设置、报酬方案、各类信息传递方式）
type AdminSystem struct {
}

//会员名录（含会员的角色、贡献值、现金、可提供的服务等）

//Member Information
//会员信息
type MemberInfo struct {
	MemberID     string
	Role         string
	MemberType   string
	UnderCoach   bool   //是否出于辅导期
	CoachID      string //辅导者memberid
	CoachEndTime string //辅导期结束时间
	ActiveMonth  string //活跃时间
	Email        string
	ISWeight     int //累计内务权重
	P2coin       int //累计贡献值
	Business     []Service
}

type Service struct {
	Title    string
	Contract string
	Price    float32
}
