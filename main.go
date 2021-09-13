package main

import (
  "net/http"

	"github.com/gin-gonic/gin"
)
var clog []Competitor = []Competitor{}

var competitors = []Competitor{
{Firstname: "Juan Manuel", Lastname: "Reyes",Username: "Seyerman", Password: "1234", ConfirmPwd: "1234",Birthday: "01/05/1994"  },
}
func loadLogout(c *gin.Context) {
	clog = []Competitor{};
	c.Redirect(http.StatusMovedPermanently, "/")
}
func loadRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
func defaultRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/rcp-auth")
}

func loadIndex(c *gin.Context) {

	if(len(clog)==0) {
    c.HTML(http.StatusOK, "login.html", gin.H {
    })
	} else {
    c.HTML(http.StatusOK, "index.html", gin.H {
      "user": clog,
      "users": competitors,
    })
    return
	}
}
func main() {
	router := gin.Default()
	router.LoadHTMLFiles("login.html", "register.html", "index.html")
	router.GET("/", defaultRedirect)
  router.GET("/rcp-auth", loadIndex)
	router.POST("/rcp-auth", loadLogin)
	router.GET("/rcp-auth/register", loadRegister)
	router.POST("/rcp-auth/register", registerComp)
	router.GET("/rcp-auth/logout", loadLogout)

	router.Run("localhost:8080")
}
func registerComp(c *gin.Context) {
  firstname := c.PostForm("Username");
  lastname  := c.PostForm("Password");
  username := c.PostForm("Username");
  password := c.PostForm("Password");
  confirmpwd := c.PostForm("ConfirmPwd");
  birthday := c.PostForm("Birthday")
  comp:=[]Competitor{
  { Firstname	: firstname,
    Lastname		: lastname,
    Username		: username,
  	Password		: password,
  	ConfirmPwd	: confirmpwd,
  	Birthday		: birthday},
  }
  if(len(firstname) > 0 || len(lastname) > 0 ){
    if( len(username) > 0 || len(password) > 0){
      if( len(confirmpwd) > 0 || len(birthday) > 0 ){
        if(password==confirmpwd){
          //competitors:={firstname,lastname,username,password,confirmpwd,birthday},
            competitors=append(competitors,comp...)
            c.HTML(http.StatusOK, "login.html", gin.H {
  					})
  					return
        }
      }
    }
  }else{
    c.HTML(http.StatusOK, "login.html", gin.H {
    })
  }
}
func loadLogin(c *gin.Context) {
	u := c.PostForm("Username");
	p := c.PostForm("Password");

	if(len(p) > 0 && len(u) > 0) {
		for _, comp := range competitors {
			if comp.Username == u {
				if comp.Password == p {
					clog := comp
					c.HTML(http.StatusOK, "index.html", gin.H {
						"username": clog.Username,
						"users": competitors,
					})
					return
				}
				c.HTML(http.StatusOK, "login.html", gin.H {
					"message": "The password is wrong",
				})
				return
			}
		}
		c.HTML(http.StatusOK, "login.html", gin.H {
			"message": "This competitor does not exist",
		})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H {
			"message": "Please fill in all the fields",
		})
	}
}
type Competitor struct {
  Firstname			string  `json:"firstname"`
  Lastname			string  `json:"lastname"`
  Username			string  `json:"username"`
	Password			string  `json:"password"`
	ConfirmPwd		string  `json:"confirmpwd"`
	Birthday			string	`json:"birthday"`
}
