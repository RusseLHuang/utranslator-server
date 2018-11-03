package credential

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var conf *oauth2.Config

func Init() {
	file, err := ioutil.ReadFile("./auth/credential.json")

	if err != nil {
		fmt.Println("File error: %v", err)
		os.Exit(1)
	}

	conf, err = google.ConfigFromJSON(file, "https://www.googleapis.com/auth/userinfo.email")
	if err != nil {
		panic(err)
	}

	fmt.Println("Initialization done")
}

func ExchangeCode(code string) []byte {
	ctx := context.Background()
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		panic(err)
	}

	client := conf.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	return data
}

func Authorize() echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfstate, err := c.Cookie("csrfstate")
		if err != nil {
			return err
		}

		if c.FormValue("state") != csrfstate.Value {
			log.Println("invalid oauth google state")
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		fmt.Println("EXCHANGE CODE : ", c.FormValue("code"))
		return c.String(http.StatusOK, "Hello store auth code")
	}
}

func HandleMain() echo.HandlerFunc {
	return func(c echo.Context) error {
		var htmlIndex = `<html>
			<body>
				<a href="/login">Google Log In</a>
			</body>
		</html>`

		return c.HTML(http.StatusOK, htmlIndex)
	}
}

// Testing Token
func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		oauthState := generateCSRFToken(c)
		u := conf.AuthCodeURL(oauthState)
		return c.Redirect(http.StatusTemporaryRedirect, u)
	}
}

func generateCSRFToken(c echo.Context) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{
		Name:    "csrfstate",
		Value:   state,
		Expires: expiration,
	}

	c.SetCookie(&cookie)

	return state
}
