import { ArticleServiceClient } from '../proto/article/ArticleServiceClientPb';
import { type Article, GetTopicArticlesRequest, type TopicArticles } from '../proto/article/article_pb';
import { API_HOST } from './constants';

// eslint-disable-next-line
class ArticleService {
    static async getArticles(topics: string[]): Promise<Map<string, Article[]>> {
        const articleServiceClient = new ArticleServiceClient(`http://${API_HOST}:80`);

        const request = new GetTopicArticlesRequest();
        request.setTopicsList(topics)

        const response: TopicArticles = await articleServiceClient.getTopicArticles(request, null);
        const grpcMap = response.getTopicarticlesMap();

        const result = new Map<string, Article[]>

        for (const [key, value] of grpcMap.entries()) {
            result.set(key, value.getArticlesList())
        }

        return result;
    }
}

export default ArticleService;
