package routes

import (
	"fmt"
	"github.com/Kvikikpt/wow-traider/blizzardApi"
	"net/http"
	"os"
)

//index page
func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Index"))

	newcall := blizzardApi.Init(os.Getenv("BLIZZARD_CLIENT_ID"), os.Getenv("BLIZZARD_CLIENT_SECRET"))
	data := newcall.MakeRequest("https://us.battle.net/oauth/userinfo")
	fmt.Println(string(data))
}
