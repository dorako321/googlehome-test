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
	e = echo.New()
)

func postToSlack(message string) {
	data := url.Values{}
	data.Set("token",token)
	data.Add("channel","#general")  // change channel
	data.Add("username","robo-jiro")
	data.Add("text", fmt.Sprintf("%s", message))

	client := &http.Client{}
	r, _ := http.NewRequest("POST",  fmt.Sprintf("%s",apiUrl), bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)

}

func main() {
	// Echoのインスタンス作る


	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/hello", MainPage())

	e.POST("/", GoogleAssistant())

	// サーバー起動
	e.Start(":"+os.Getenv("PORT"))
}

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		//e.Logger.Error("Hello World")
		//e.Logger.Error(c.Request())
		return c.String(http.StatusOK, "Hello World")
	}
}

type Parameters struct {
	Any string `json:any`
}

type Result struct {
	Source string `json:source`
	Parameters Parameters `json:parameters`
}

type GARequest struct {
	Id   string `json:"id"`
	Timestamp string `json:"timestamp"`
	Result Result `json:result`

}

func GoogleAssistant() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(GARequest)
		c.Bind(u);

		e.Logger.Error("post to slack")
		e.Logger.Error(u.Result.Parameters.Any)
		postToSlack(u.Result.Parameters.Any)

		return c.String(http.StatusOK, "OK")
	}
}
