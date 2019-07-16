package gin_log_stream

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpcloud/tail"
	"io"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET,DELETE,PUT,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func read(ch chan string) {
	fileName := "/var/log/nginx/access.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		ch <- err.Error()
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		fmt.Println(msg)
		if !ok {
			ch <- "不ok"
		} else {
			ch <- msg.Text
		}
	}

}
func main() {
	app := gin.Default()
	app.Use(Cors())
	app.GET("/event", func(context *gin.Context) {
		ch := make(chan string, 10)
		go read(ch)
		clientGone := context.Writer.CloseNotify()
		context.Stream(func(w io.Writer) bool {
			select {
			case <-clientGone:
				return false
			case message := <-ch:
				context.SSEvent("message", message)
				return true
			}
		})
	})
	app.Run("0.0.0.0:8856")
}
