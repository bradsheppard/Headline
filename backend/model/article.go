package model

import (
	"gorm.io/gorm"
	pb "headline/proto/article"
)

type Article struct {
	gorm.Model
	Title       string
	Description string
	Url         string
	ImageUrl    string
	Source      string
	UserID      int `gorm:"index"`
	Interest    string
}

func ToArticleProtos(articles []Article) []*pb.Article {
	var protoArticles []*pb.Article

	for _, article := range articles {
		protoArticle := &pb.Article{
			Title:       article.Title,
			Description: article.Description,
			Url:         article.Url,
			ImageUrl:    article.ImageUrl,
			Source:      article.Source,
			Interest:    article.Interest,
		}
		protoArticles = append(protoArticles, protoArticle)
	}

	return protoArticles
}

func FromArticleProtos(protoArticles []*pb.Article, userId uint64) []*Article {
	var articles []*Article

	for _, protoArticle := range protoArticles {
		article := &Article{
			Title:       protoArticle.Title,
			Description: protoArticle.Description,
			Url:         protoArticle.Url,
			UserID:      int(userId),
			ImageUrl:    protoArticle.ImageUrl,
			Source:      protoArticle.Source,
			Interest:    protoArticle.Interest,
		}
		articles = append(articles, article)
	}

	return articles
}
