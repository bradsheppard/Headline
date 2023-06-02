import {create} from "zustand";
import ArticleService from "./api/article";
import InterestService from "./api/interest";
import {Article} from "./proto/article/article_pb";
import {Interest} from "./proto/interest/interest_pb";

interface State {
    interests: Array<Interest>
    articles: Array<Article>

    selectedInterest: string | null

    setSelectedInterest: (interest: string) => void
    fetchInterests: (userId: number) => Promise<void>
    fetchArticles: (userId: number) => Promise<void>

    deleteInterest: (id: number) => Promise<void>
}

const useStore = create<State>((set) => ({
    interests: [],
    articles: [],
    selectedInterest: null,

    setSelectedInterest: (interest: string) => {
        set({selectedInterest: interest})
    },

    fetchInterests: async(userId: number) => {
        const interestResponse = await InterestService.getInterests(userId)
        set({interests: interestResponse})
    },
    fetchArticles: async(userId: number) => {
        const articlesResponse = await ArticleService.getArticles(userId)
        set({articles: articlesResponse})
    },
    deleteInterest: async(id: number) => {
        await InterestService.deleteInterest(id)
    }
}))

export {useStore}
