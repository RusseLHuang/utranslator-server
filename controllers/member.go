package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/utranslator-server/dto"
	model "github.com/utranslator-server/models"
	credential "github.com/utranslator-server/utils"
)

type Member struct {
	GoogleAccessToken string `json:"accessToken"`
}

func GetMember() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		memberID := c.Param("id")
		member, err := (&model.Member{}).Get(memberID)

		return c.JSON(http.StatusOK, member)
	}
}

func CreateGoogleToken() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		m := Member{}
		if err = c.Bind(&m); err != nil {
			return
		}

		accessToken := m.GoogleAccessToken
		userInfo := credential.ExchangeCode(accessToken)

		googleUserInfo := dto.GoogleUserInfo{}
		err = json.Unmarshal(userInfo, &googleUserInfo)
		if err != nil {
			return
		}

		memberModel := model.MemberByGoogle{}
		member, err := memberModel.GetByGoogleId(googleUserInfo.ID)
		if err != nil && err.Error() != "not found" {
			return
		}

		if member == nil {
			member, err = memberModel.CreateByGoogle(googleUserInfo)
			if err != nil {
				return
			}
		}

		memberToken := &model.Member{
			MemberID: member.MemberID,
			Email:    member.Email,
		}
		tokenAuth, err := credential.GenerateToken(memberToken)
		if err != nil {
			return
		}

		return c.JSON(http.StatusOK, tokenAuth)
	}
}
