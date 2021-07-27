 package main 
import "time"
type AllMember struct {
	AllMembers []Member
}
type MemberPage struct {
	TargetMem Member
}
type Member struct {
	ID                 int
	Name            string
	Email            string
	JoiningDate  time.Time
	Phone            int
}
type Credentials struct {
	password      string
	username      string 
}
type Query struct {
	Email            string
	Query          string
}
type ErrorPage struct {
	ErrorMsg     string
}
