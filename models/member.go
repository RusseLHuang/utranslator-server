package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/utranslator-server/database"
	"github.com/utranslator-server/dto"
)

type IMember interface {
	Get(memberId string) (m *Member, err error)
}

type Member struct {
	ID                bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	FacebookID        string        `bson:"facebook_id" json:"facebookId"`
	GoogleID          string        `bson:"google_id" json:"googleId"`
	Email             string        `bson:"email" json:"email"`
	Username          string        `bson:"username" json:"username"`
	MobilePhoneNumber string        `bson:"mobile_phone_number" json:"mobilePhoneNumber"`
	CreatedAt         time.Time     `bson:"created_at" json:"createdAt"`
	UpdatedAt         time.Time     `bson:"updated_at" json:"updatedAt"`
}

func (member *Member) Save() (err error) {
	collection := db.GetDB().C("member")
	err = collection.Insert(member)

	if err != nil && err.Error() == "no reachable servers" {
		panic(err)
	}

	return
}

func (member *Member) CreateByGoogle(googleUserInfo dto.GoogleUserInfo) (*Member, error) {
	m := Member{
		GoogleID:  googleUserInfo.ID,
		Email:     googleUserInfo.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	m.Save()

	return &m, nil
}

func (member *Member) Get(memberId string) (m *Member, err error) {
	bsonObjId := bson.ObjectIdHex(memberId)
	collection := db.GetDB().C("member")
	err = collection.FindId(bsonObjId).One(&m)

	return m, err
}

func (member *Member) GetByGoogleId(memberGoogleId string) (m *Member, err error) {
	collection := db.GetDB().C("member")
	err = collection.Find(bson.M{"google_id": memberGoogleId}).One(&m)

	return
}

func (member *Member) Post(memberId string) (m *Member, err error) {
	bsonObjId := bson.ObjectIdHex(memberId)
	collection := db.GetDB().C("member")
	err = collection.FindId(bsonObjId).One(&m)

	return m, err
}
