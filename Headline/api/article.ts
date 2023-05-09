import {ArticleServiceClient} from "../proto/article/ArticleServiceClientPb"
import {User, UserArticles} from "../proto/article/article_pb";
import {API_HOST} from "./constants";

class ArticleService {

    static async getArticles(userId: number): Promise<Array<Article>> {
        const articleServiceClient = new ArticleServiceClient(`http://${API_HOST}:80`);

        const user = new User();
        user.setUserid(1)

        const response: UserArticles = await articleServiceClient.getArticles(user, null)

        return response.getArticlesList().map(x => {
            const article: Article = {
                description: x.getDescription(),
                imageUrl: x.getImageurl(),
                title: x.getTitle(),
                url: x.getUrl(),
                interest: x.getInterest()
            }

            return article;
        });
    }
}

export interface Article {
    title: string
    url: string
    imageUrl: string
    description: string
    interest: string
}

export default ArticleService;

