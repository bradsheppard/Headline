package model

import (
        "gorm.io/gorm"
	pb "headline/proto/article"
)

type Article struct {
        gorm.Model
	Title   string 
	Summary string 
	Link    string 
        UserID  int
}

func ToArticleProtos(articles []Article) []*pb.Article {
        var protoArticles []*pb.Article

        for _, article := range(articles) {
                protoArticle := &pb.Article{
                        Title: article.Title,
                        Summary: article.Summary,
                        Link: article.Link,
                        UserId: uint64(article.UserID),
                }
                protoArticles = append(protoArticles, protoArticle)
        }

        return protoArticles
}

func FromArticleProtos(protoArticles []*pb.Article) []*Article {
        var articles []*Article

        for _, protoArticle := range(protoArticles) {
                article := &Article{
                        Title: protoArticle.Title,
                        Summary: protoArticle.Summary,
                        Link: protoArticle.Link,
                        UserID: int(protoArticle.UserId),
                }
                articles = append(articles, article)
        }

        return articles
}

