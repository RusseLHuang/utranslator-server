package model

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	db "github.com/utranslator-server/database"
	"github.com/utranslator-server/dto"
)

type Member struct {
	MemberID gocql.UUID `json:"memberId,omitempty"`
	Email    string     `json:"email"`
	Username string     `json:"username"`
}

type MemberByFacebook struct {
	MobilePhoneNumber string     `json:"mobilePhoneNumber"`
	FacebookID        string     `json:facebookId`
	MemberID          gocql.UUID `json:"memberId,omitempty"`
}

type MemberByGoogle struct {
	MemberID gocql.UUID `json:"memberId,omitempty"`
	GoogleID string     `json:googleId`
	Email    string     `json:"email"`
}

func (m *Member) Save() (id gocql.UUID, err error) {
	id, err = gocql.RandomUUID()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.GetSession().Query(`INSERT INTO member (member_id, email, username) VALUES (?, ?, ?)`,
		id, m.Email, m.Username).Exec(); err != nil {
		log.Fatal(err)
	}
	return
}

func (m *MemberByGoogle) SaveByGoogle() {
	if err := db.GetSession().Query(`INSERT INTO member_by_google (member_id, google_id, email) VALUES (?, ?, ?)`,
		m.MemberID, m.GoogleID, m.Email).Exec(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func (m *MemberByGoogle) CreateByGoogle(googleUserInfo dto.GoogleUserInfo) (googleMember *MemberByGoogle, err error) {
	member := Member{
		Username: "",
		Email:    googleUserInfo.Email,
	}
	memberID, err := member.Save()

	if err != nil {
		return googleMember, err
	}

	googleMember = &MemberByGoogle{
		MemberID: memberID,
		GoogleID: googleUserInfo.ID,
		Email:    googleUserInfo.Email,
	}
	googleMember.SaveByGoogle()

	return googleMember, nil
}

func (m *Member) Get(memberId string) (member *Member, err error) {
	mi := map[string]interface{}{}
	uuid, err := gocql.ParseUUID(memberId)
	if err != nil {
		log.Fatal(err)
	}

	query := db.GetSession().Query(`SELECT member_id, email, username FROM member WHERE member_id = ? LIMIT 1`,
		uuid).Consistency(gocql.One).Iter()

	for query.MapScan(mi) {
		member = &Member{
			MemberID: mi["member_id"].(gocql.UUID),
			Email:    mi["email"].(string),
			Username: mi["username"].(string),
		}
	}

	return member, err
}

func (m *MemberByGoogle) GetByGoogleId(memberGoogleId string) (member *MemberByGoogle, err error) {
	mi := map[string]interface{}{}
	query := db.GetSession().Query(`SELECT google_id, member_id, email FROM member_by_google WHERE google_id = ? LIMIT 1`,
		memberGoogleId).Consistency(gocql.One).Iter()

	for query.MapScan(mi) {
		member = &MemberByGoogle{
			GoogleID: mi["google_id"].(string),
			MemberID: mi["member_id"].(gocql.UUID),
			Email:    mi["email"].(string),
		}
	}

	return member, err
}

func (member *Member) Post(memberId string) (m *Member, err error) {
	fmt.Println(memberId)
	mi := map[string]interface{}{}
	query := db.GetSession().Query(`SELECT * FROM member WHERE member_id = ? LIMIT 1`,
		memberId).Consistency(gocql.One)

	err = query.Scan(&mi)
	fmt.Println(mi)
	if err != nil {
		log.Fatal(err)
	}

	return nil, err
}
