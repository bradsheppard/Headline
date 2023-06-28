package model

import (
	pb "headline/proto/article"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string
	Description string
	Url         string
	ImageUrl    string
	Source      string
	TopicName   string
}

func ToConverter(articles []Article) *pb.TopicArticles {
	topicArticles := &pb.TopicArticles{}
	topicArticles.TopicArticles = map[string]*pb.Articles{}

	for _, article := range articles {
		if _, ok := topicArticles.TopicArticles[article.TopicName]; !ok {
			topicArticles.TopicArticles[article.TopicName] = &pb.Articles{}
		}
		articleList := topicArticles.TopicArticles[article.TopicName].Articles

		protoArticle := &pb.Article{
			Title:       article.Title,
			Description: article.Description,
			Url:         article.Url,
			ImageUrl:    article.ImageUrl,
			Source:      article.Source,
		}

		articleList = append(articleList, protoArticle)
		topicArticles.TopicArticles[article.TopicName].Articles = articleList
	}

	return topicArticles
}

func ToTopicArticleProto(topics []Topic) *pb.TopicArticles {
	topicArticles := &pb.TopicArticles{}
	topicArticles.TopicArticles = map[string]*pb.Articles{}

	for _, topic := range topics {
		topicArticles.TopicArticles[topic.Name] = &pb.Articles{}
		articleList := topicArticles.TopicArticles[topic.Name].Articles

		for _, article := range topic.Articles {
			protoArticle := &pb.Article{
				Title:       article.Title,
				Description: article.Description,
				Url:         article.Url,
				ImageUrl:    article.ImageUrl,
				Source:      article.Source,
			}
			articleList = append(articleList, protoArticle)
		}
	}

	return topicArticles
}

func FromArticleProtos(protoArticles []*pb.Article) []*Article {
	var articles []*Article

	for _, protoArticle := range protoArticles {
		article := &Article{
			Title:       protoArticle.Title,
			Description: protoArticle.Description,
			Url:         protoArticle.Url,
			ImageUrl:    protoArticle.ImageUrl,
			Source:      protoArticle.Source,
		}
		articles = append(articles, article)
	}

	return articles
}
