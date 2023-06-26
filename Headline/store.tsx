import { create } from 'zustand';
import ArticleService from './api/article';
import TopicService from './api/topic';
import { type Article } from './proto/article/article_pb';
import { type Topic } from './proto/topic/topic_pb';

export interface UserInfo {
    name: string;
    image: string;
}

interface State {
    topics: Topic[];
    articles: Map<string, Article[]>

    userInfo: UserInfo | null;
    token: string | null;

    selectedTopic: string | null;

    setSelectedTopic: (topic: string) => void;

    setUserInfo: (userInfo: UserInfo) => void;
    setToken: (token: string | null) => void;

    fetchTopics: (userId: number) => Promise<void>;
    fetchArticles: (topics: string[]) => Promise<void>;

    deleteTopic: (topic: string) => Promise<void>;
}

const useStore = create<State>((set) => ({
    topics: [],
    articles: new Map<string, Article[]>(),
    selectedTopic: null,
    token: null,

    userInfo: null,

    setSelectedTopic: (topic: string) => {
        set({ selectedTopic: topic });
    },

    setUserInfo: (userInfo: UserInfo) => {
        set({ userInfo });
    },

    setToken: (token: string | null) => {
        set({ token });
    },

    fetchTopics: async (userId: number) => {
        const topicResponse = await TopicService.getTopics(userId);
        set({ topics: topicResponse });
    },
    fetchArticles: async (topics: string[]) => {
        const articlesResponse = await ArticleService.getArticles(topics);
        set({ articles: articlesResponse });
    },
    deleteTopic: async (topicName: string) => {
        await TopicService.deleteTopic(topicName);
        set((state) => {
            const newTopics = state.topics.filter((topic) => topic.getName() !== topicName);
            return {
                topics: newTopics,
            };
        });
    },
}));

export { useStore };
