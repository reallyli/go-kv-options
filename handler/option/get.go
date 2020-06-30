package option

import (
	. "go-options/handler"
	"go-options/model"
	"go-options/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Get(c *gin.Context) {
	module := c.Query("module")
	key := c.Query("key")
	log.Infof(module + key)

	configs, err := model.GetConfigs(module, key)
	if err != nil {
		log.Infof(err.Error())
		SendResponse(c, errno.ErrOptionNotFound, nil)
		return
	}
	SendResponse(c, nil, configs)
}
