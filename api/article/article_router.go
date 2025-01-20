package article

import (
	"github.com/duke-git/lancet/v2/convertor"
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
			Dst:       &DTO{},
			ParamType: "body",
			Url:       "/save",
			Method:    "POST",
			Func: func(c *gin.Context, dst interface{}) {
				dto := dst.(*DTO)
				md := convertToArticle(dto)
				global.DB.Create(&md)
				response.Success(c, "保存成功")
			},
		}, {
			Url:    "/list_all",
			Method: "GET",
			Func: func(c *gin.Context, dst interface{}) {
				var articles []model.Article
				global.DB.Model(articles).Order("created_at desc").Preload("Categories").Preload("Tags").Find(&articles)
				result := slice.Map(articles, func(i int, article model.Article) VO {
					return convertToVO(&article)
				})
				response.SuccessWithData(c, "查询成功", result)
			},
		}, {
			Url:    "/list_top_10",
			Method: "GET",
			Func: func(c *gin.Context, dst interface{}) {
				var articles []model.Article
				global.DB.Model(articles).Order("created_at desc").Limit(10).Find(&articles)
				result := slice.Map(articles, func(i int, article model.Article) SimpleVO {
					return convertToSimpleVO(&article)
				})
				response.SuccessWithData(c, "查询成功", result)
			},
		}, {
			Url:    "/get_by_id",
			Method: "GET",
			Func: func(c *gin.Context, dst interface{}) {
				var article model.Article
				id := c.Query("id")
				if id == "" {
					response.Fail(c, "参数错误")
					return
				}
				result := global.DB.Model(article).Where("id = ?", id).Preload("Categories").Preload("Tags").First(&article)
				if result.RowsAffected == 0 {
					response.Fail(c, "未找到文章数据")
					return
				}
				response.SuccessWithData(c, "查询成功", convertToVO(&article))
			},
		}, {
			Url:    "/search",
			Method: "GET",
			Func: func(c *gin.Context, dst interface{}) {
				var articles []model.Article
				global.DB.Model(articles).Preload("Categories", "name in ", nil).Preload("Tags", "name in ", nil).Find(&articles)
				result := slice.Map(articles, func(i int, article model.Article) SimpleVO {
					return convertToSimpleVO(&article)
				})
				response.SuccessWithData(c, "查询成功", result)
			},
		},
	}
)

type DTO struct {
	Title         string   `json:"title"`
	SubTitle      string   `json:"sub_title"`
	Content       string   `json:"content"`
	CategoryNames []string `json:"category_names"`
	TagNames      []string `json:"tag_names"`
}

type VO struct {
	ID            uint
	Title         string   `json:"title"`
	SubTitle      string   `json:"sub_title"`
	Content       string   `json:"content"`
	CategoryNames []string `json:"category_names"`
	TagNames      []string `json:"tag_names"`
	CreatedAt     string   `json:"created_at"`
}

type SimpleVO struct {
	ID    uint
	Title string `json:"title"`
}

func convertToArticle(dto *DTO) model.Article {
	md := model.Article{
		Title:    dto.Title,
		SubTitle: dto.SubTitle,
		Content:  dto.Content,
		Status:   model.StatusDraft,
	}
	var categories []model.Category
	global.DB.Find(&categories, "name in ?", dto.CategoryNames)
	categoriesMap := convertor.ToMap(categories, func(c model.Category) (string, model.Category) {
		return c.Name, c
	})
	for _, name := range dto.CategoryNames {
		var c model.Category
		if category, ok := categoriesMap[name]; ok {
			c = category
		} else {
			c = model.Category{Name: name}
		}
		md.Categories = append(md.Categories, c)
	}
	var tags []model.Tag
	global.DB.Find(&tags, "name in ?", dto.TagNames)
	tagsMap := convertor.ToMap(tags, func(t model.Tag) (string, model.Tag) {
		return t.Name, t
	})
	for _, name := range dto.TagNames {
		var t model.Tag
		if tag, ok := tagsMap[name]; ok {
			t = tag
		} else {
			t = model.Tag{Name: name}
		}
		md.Tags = append(md.Tags, t)
	}
	return md
}

func convertToVO(article *model.Article) VO {
	return VO{
		ID:       article.ID,
		Title:    article.Title,
		SubTitle: article.SubTitle,
		Content:  article.Content,
		CategoryNames: slice.Map(article.Categories, func(i int, category model.Category) string {
			return category.Name
		}),
		TagNames: slice.Map(article.Tags, func(i int, tag model.Tag) string {
			return tag.Name
		}),
		CreatedAt: article.CreatedAt.Format("January 2, 2006"),
	}
}

func convertToSimpleVO(article *model.Article) SimpleVO {
	return SimpleVO{
		ID:    article.ID,
		Title: article.Title,
	}
}
