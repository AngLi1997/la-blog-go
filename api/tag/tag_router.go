package tag

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"la-blog-go/api"
	"la-blog-go/global"
	"la-blog-go/model"
	"la-blog-go/response"
)

var (
	Apis = []api.Api{
		{
			Url:    "/list_all",
			Method: "GET",
			Func: func(c *gin.Context) {
				var tags []model.Tag
				global.DB.Model(tags).
					Preload("Articles").
					Order("created_at desc").
					Find(&tags)
				result := slice.Map(tags, func(i int, tag model.Tag) SimpleVO {
					return convertToSimpleVO(&tag)
				})
				response.SuccessWithData(c, "查询成功", result)
			},
		},
	}
)

func convertToSimpleVO(tag *model.Tag) SimpleVO {
	return SimpleVO{
		ID:    tag.ID,
		Name:  tag.Name,
		Count: len(tag.Articles),
	}
}

type SimpleVO struct {
	ID    uint
	Name  string `json:"name"`
	Count int    `json:"count"`
}
