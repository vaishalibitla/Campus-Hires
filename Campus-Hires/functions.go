package main  
import (
	"fmt"
	"time"
	"github.com/lib/pq"
)
func (m Member) JoiningDateStr() string {
	return m.JoiningDate.Format("2006-01-02")
}
func getMember(memberID int) (Member, error) {
	res := Member{}
	var id int
	var name string
	var email string
	var phone int
	var joiningDate pq.NullTime
	err := db.QueryRow(`SELECT id, name, email, phone, joining_date FROM members where id = $1`, memberID).Scan(&id, &name, &email, &phone, &joiningDate)
	if err == nil {
		res = Member{ID: id, Name: name, Email: email, Phone: phone, JoiningDate: joiningDate.Time}
	}
	return res, err
}
func allMembers() ([]Member, error) {
	members := []Member{}
	rows, err := db.Query(`SELECT id, name, email, phone, joining_date FROM members order by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var email string
		var phone int
		var joiningDate pq.NullTime
		err = rows.Scan(&id, &name, &email, &phone, &joiningDate)
		if err != nil {
			return members, err
		}
		currentMember := Member{ID: id, Name: name, Email: email, Phone: phone}
		if joiningDate.Valid {
			currentMember.JoiningDate = joiningDate.Time
		}
		members = append(members, currentMember)
	}
	return members, err
}
func insertMember(name, email string, phone int, joiningDate time.Time) (int, error) {
	var memberID int
	err := db.QueryRow(`INSERT INTO members(name, email, phone, joining_date) VALUES($1, $2, $3, $4) RETURNING id`, name, email, phone, joiningDate).Scan(&memberID)
	if err != nil {
		return 0, err
	}
	fmt.Printf("Last inserted Member ID: %v\n", memberID)
	return memberID, err
}

func updateMember(id int, name, email string, phone int, joiningDate time.Time) (int, error) {
	res, err := db.Exec(`UPDATE members set name=$1, email=$2, phone=$3, joining_date=$4 where id=$5 RETURNING id`, name, email, phone, joiningDate, id)
	if err != nil {
		return 0, err
	}
	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
                fmt.Printf("Last updated Member ID: %v\n", id)
	return int(rowsUpdated), err
}
func removeMember(memberID int) (int, error) {
	res, err := db.Exec(`DELETE FROM members WHERE id = $1`, memberID)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
                fmt.Printf("Last deleted Member ID: %v\n", memberID)
	return int(rowsDeleted), nil
}
func insertQuery(email,query string)(int, error){
	res,err := db.Exec(`INSERT INTO queries( email, query) VALUES($1,$2)`, email,query)
                if err != nil {
		return 0,err
	}
                rowsInserted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
                fmt.Printf("Last query sent by: %v\n", email)
	return int(rowsInserted), nil
}