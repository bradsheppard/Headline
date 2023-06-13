import { create } from 'zustand';
import ArticleService from './api/article';
import InterestService from './api/interest';
import { type Article } from './proto/article/article_pb';
import { type Interest } from './proto/interest/interest_pb';

export interface UserInfo {
    name: string;
    image: string;
}

interface State {
    interests: Interest[];
    articles: Article[];

    userInfo: UserInfo | null;
    token: string | null;

    selectedInterest: string | null;

    setSelectedInterest: (interest: string) => void;

    setUserInfo: (userInfo: UserInfo) => void;
    setToken: (token: string | null) => void;

    fetchInterests: (userId: number) => Promise<void>;
    fetchArticles: (userId: number) => Promise<void>;

    deleteInterest: (id: number) => Promise<void>;
}

const useStore = create<State>((set) => ({
    interests: [],
    articles: [],
    selectedInterest: null,
    token: null,

    userInfo: null,

    setSelectedInterest: (interest: string) => {
        set({ selectedInterest: interest });
    },

    setUserInfo: (userInfo: UserInfo) => {
        set({ userInfo });
    },

    setToken: (token: string | null) => {
        set({ token });
    },

    fetchInterests: async (userId: number) => {
        const interestResponse = await InterestService.getInterests(userId);
        set({ interests: interestResponse });
    },
    fetchArticles: async (userId: number) => {
        const articlesResponse = await ArticleService.getArticles(userId);
        set({ articles: articlesResponse });
    },
    deleteInterest: async (id: number) => {
        await InterestService.deleteInterest(id);
        set((state) => {
            const newInterests = state.interests.filter((interest) => interest.getId() !== id);
            return {
                interests: newInterests,
            };
        });
    },
}));

export { useStore };
