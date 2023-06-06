import { ArticleServiceClient } from '../proto/article/ArticleServiceClientPb';
import { User, type UserArticles, type Article } from '../proto/article/article_pb';
import { API_HOST } from './constants';

// eslint-disable-next-line
class ArticleService {
    static async getArticles(userId: number): Promise<Article[]> {
        const articleServiceClient = new ArticleServiceClient(`http://${API_HOST}:80`);

        const user = new User();
        user.setUserid(1);

        const response: UserArticles = await articleServiceClient.getArticles(user, null);
        return response.getArticlesList();
    }
}

export default ArticleService;
