package model

import (
        "gorm.io/gorm"
	pb "headline/proto"
)

type Article struct {
        gorm.Model
	Title   string 
	Summary string 
	Link    string 
}

func FromProto(article *pb.Article) *Article {
        return &Article{
                Title: article.Title,
                Summary: article.Summary,
                Link: article.Link,
        }
}

func ToProto(article Article) *pb.Article {
        return &pb.Article{
                Title: article.Title,
                Summary: article.Summary,
                Link: article.Link,
        }
}
