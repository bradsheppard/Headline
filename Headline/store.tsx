import { StoreApi, create } from 'zustand';
import { type Article, TopicNames, TopicArticles } from './proto/article/article_pb';
import {TopicServiceClient} from './proto/topic/TopicServiceClientPb';
import { API_HOST } from './api/constants';
import empty from 'google-protobuf/google/protobuf/empty_pb';
import {
    AddTopicsRequest, RemoveTopicsRequest, Topic
} from './proto/topic/topic_pb';
import {ArticleServiceClient} from './proto/article/ArticleServiceClientPb';

export interface UserInfo {
    name: string;
    image: string;
}

interface State {
    topics: Topic[];
    articles: Map<string, Article[]>

    userInfo: UserInfo | null;
    accessToken: string | null;
    idToken: string | null;

    selectedTopic: string | null;

    setSelectedTopic: (topic: string) => void;

    setUserInfo: (userInfo: UserInfo) => void;
    setAccessToken: (token: string | null) => void;
    setIdToken: (token: string | null) => void;

    createTopic: (topic: string) => Promise<void>;

    fetchTopics: () => Promise<void>;
    fetchArticles: (topics: string[]) => Promise<void>;

    deleteTopic: (topic: string, userId: number) => Promise<void>;
}

const getTopicClient = () => {
    return new TopicServiceClient(`http://${API_HOST}:80`);
}

const getArticleClient = () => {
    return new ArticleServiceClient(`http://${API_HOST}:80`);
}

const getMetadata = (token: string) => {
    return {
        'id_token': token
    }
}

const useStore = create<State>((set, get) => ({
    topics: [],
    articles: new Map<string, Article[]>(),
    selectedTopic: null,
    accessToken: null,
    idToken: null,

    userInfo: null,

    setSelectedTopic: (topic: string) => {
        set({ selectedTopic: topic });
    },

    setUserInfo: (userInfo: UserInfo) => {
        set({ userInfo });
    },

    setAccessToken: (token: string | null) => {
        set({ accessToken: token });
    },

    setIdToken: (token: string | null) => {
        set({ idToken: token });
    },

    createTopic: async (topicName: string) => {
        const idToken = get().idToken;

        if (idToken === null)
            return;

        const metadata = getMetadata(idToken);

        const topicServiceClient = getTopicClient();

        const topic = new Topic();
        topic.setName(topicName);

        const request = new AddTopicsRequest();
        request.setTopicsList([topic])

        const response = await topicServiceClient.addTopics(request, metadata);
        set({ topics: response.getTopicsList() });
    },

    fetchTopics: async () => {
        const idToken = get().idToken;

        if (idToken === null)
            return;

        const metadata = getMetadata(idToken);

        const topicServiceClient = getTopicClient()

        const response = await topicServiceClient.getTopics(new empty.Empty(), metadata);
        set({ topics: response.getTopicsList() });
    },
    fetchArticles: async (topics: string[]) => {
        const idToken = get().idToken;

        if (idToken === null)
            return;

        const articleServiceClient = getArticleClient();

        const request = new TopicNames();
        request.setTopicsList(topics)

        const metadata = getMetadata(idToken);

        const response: TopicArticles = await articleServiceClient.getTopicArticles(request, metadata);
        const grpcMap = response.getTopicarticlesMap();

        const result = new Map<string, Article[]>

        for (const [key, value] of grpcMap.entries()) {
            result.set(key, value.getArticlesList())
        }

        set({ articles: result });
    },
    deleteTopic: async (topicName: string, userId: number) => {
        const idToken = get().idToken;

        if (idToken === null)
            return;

        const topicServiceClient = getTopicClient();

        const request = new RemoveTopicsRequest();
        request.setTopicnamesList([topicName]);

        const metadata = getMetadata(idToken);

        await topicServiceClient.removeTopics(request, metadata);
        set((state) => {
            const newTopics = state.topics.filter((topic) => topic.getName() !== topicName);
            return {
                topics: newTopics,
            };
        });
    },
}));

export { useStore };
