package category

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
				var categories []model.Category
				global.DB.Model(categories).
					Preload("Articles").
					Order("created_at desc").
					Find(&categories)
				result := slice.Map(categories, func(i int, category model.Category) SimpleVO {
					return convertToSimpleVO(&category)
				})
				response.SuccessWithData(c, "查询成功", result)
			},
		},
	}
)

func convertToSimpleVO(category *model.Category) SimpleVO {
	return SimpleVO{
		ID:    category.ID,
		Name:  category.Name,
		Count: len(category.Articles),
	}
}

type SimpleVO struct {
	ID    uint
	Name  string `json:"name"`
	Count int    `json:"count"`
}
