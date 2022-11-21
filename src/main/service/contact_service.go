package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"star-im/src/main/models"
	"star-im/src/main/utils"
)

func SearchFriends(context *gin.Context) {
	//id, _ := strconv.Atoi(c.Request.FormValue("userId"))
	body, _ := io.ReadAll(context.Request.Body)
	user := models.User{}
	err := json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	users := models.SearchFriend(user.ID)
	utils.RespOKList(context.Writer, users, len(users))
}
