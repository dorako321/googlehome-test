package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"bytes"
	"os"
	"net/url"
	"fmt"
)

var(
	token  string = "" // change token
	apiUrl string = "https://slack.com/api/chat.postMessage"
	e             = echo.New()
)

/**
 * 指定した文字列をSlackにポストする
 */
func postToSlack(message string) {
	data := url.Values{}
	data.Set("token", token)
	data.Add("channel", "#general")
	data.Add("username", "robo-jiro")
	data.Add("text", fmt.Sprintf("%s", message))

	client := &http.Client{}
	r, _ := http.NewRequest("POST", fmt.Sprintf("%s", apiUrl), bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)

}

func main() {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// ルーティング
	e.POST("/", GoogleAssistant())

	e.Start(":" + os.Getenv("PORT"))
}

type Parameters struct {
	Any string `json:any`
}

type Result struct {
	Parameters Parameters `json:parameters`
}

type GARequest struct {
	Id        string `json:"id"`
	Result    Result `json:result`
}

func GoogleAssistant() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(GARequest)
		c.Bind(u);
		postToSlack(u.Result.Parameters.Any)
		return c.String(http.StatusOK, "")
	}
}
