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
	TopicName   *string
}

func ToConverter(articles []Article) *pb.TopicArticles {
	topicArticles := &pb.TopicArticles{}
	topicArticles.TopicArticles = map[string]*pb.Articles{}

	for _, article := range articles {
		if article.TopicName == nil {
			continue
		}

		topicName := *article.TopicName

		if _, ok := topicArticles.TopicArticles[topicName]; !ok {
			topicArticles.TopicArticles[topicName] = &pb.Articles{}
		}
		articleList := topicArticles.TopicArticles[topicName].Articles

		protoArticle := &pb.Article{
			Title:       article.Title,
			Description: article.Description,
			Url:         article.Url,
			ImageUrl:    article.ImageUrl,
			Source:      article.Source,
		}

		articleList = append(articleList, protoArticle)
		topicArticles.TopicArticles[topicName].Articles = articleList
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

func ToArticleProtos(articles []*Article) []*pb.Article {
	var protoArticles []*pb.Article

	for _, article := range articles {
		protoArticle := &pb.Article{
			Title:       article.Title,
			Description: article.Description,
			Url:         article.Url,
			ImageUrl:    article.ImageUrl,
			Source:      article.Source,
		}
		protoArticles = append(protoArticles, protoArticle)
	}

	return protoArticles
}
