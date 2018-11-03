package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/utranslator-server/dto"
	"github.com/utranslator-server/models"
	"github.com/utranslator-server/utils"
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

		memberModel := model.Member{}
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

		tokenAuth, err := credential.GenerateToken(member)
		if err != nil {
			return
		}

		return c.JSON(http.StatusOK, tokenAuth)
	}
}
