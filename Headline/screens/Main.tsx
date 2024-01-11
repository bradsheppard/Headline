import * as React from 'react';
import { FlatList, RefreshControl, View } from 'react-native';
import uuid from 'react-native-uuid';
import Tags from '../components/Tags';
import { useEffect, useState } from 'react';
import { type NativeStackScreenProps } from '@react-navigation/native-stack';
import { useStore } from '../store';
import { type Article as ArticleProto } from '../proto/article/article_pb';
import type { RootStackParamList } from '../App';
import Article from '../components/Article';
import {SpecialTopics} from '../api/constants';

const styles = {
    container: {
        flex: 1,
        flexGrow: 1,
        backgroundColor: '#000',
    },
    list: {
        flex: 1,
        flexGrow: 1,
        paddingVertical: 8,
    },
};

type Props = NativeStackScreenProps<RootStackParamList, 'Feed'>;

export default function Main(props: Props): JSX.Element {
    const [topicArticles, trendingArticles, topics, selectedTopic] = useStore((state) => [
        state.topicArticles, 
        state.trendingArticles, 
        state.topics, 
        state.selectedTopic
    ]);
    const [fetchArticles, fetchTrendingArticles, fetchTopics] = useStore((state) => [
        state.fetchTopicArticles,
        state.fetchTrendingArticles,
        state.fetchTopics,
    ]);

    const [filteredArticles, setFilteredArticles] = useState<ArticleProto[]>([]);
    const [isLoading, setIsLoading] = useState(false);

    const fetchAndFilter = async (): Promise<void> => {
        await fetchData();
        filterArticles();
    };

    const fetchData = async (): Promise<void> => {
        setIsLoading(true);
        await Promise.all([
            fetchTopics(),
            fetchArticles(topics.map(x => x.getName())),
            fetchTrendingArticles()
        ])
        setIsLoading(false);
    };

    const filterArticles = (): void => {
        if (selectedTopic !== null && selectedTopic !== SpecialTopics.Trending) {
            const selectedArticles = topicArticles.get(selectedTopic);

            if (selectedArticles === undefined) {
                setFilteredArticles([])
                return
            }

            setFilteredArticles(selectedArticles);
        } else {
            setFilteredArticles(trendingArticles);
        }
    };

    useEffect(() => {
        void fetchData();
    }, []);

    useEffect(() => {
        filterArticles();
    }, [selectedTopic]);

    return (
        <View style={[styles.container]}>
            <Tags topics={topics.map((x) => x.getName())} />
            <FlatList
                style={styles.list}
                data={filteredArticles}
                keyExtractor={() => uuid.v4() as string}
                renderItem={(entry) => {
                    const date = entry.item.getDate();

                    return (
                        <Article
                            article={{
                                description: entry.item.getDescription(),
                                imageUrl: entry.item.getImageurl(),
                                title: entry.item.getTitle(),
                                url: entry.item.getUrl(),
                                source: entry.item.getSource(),
                                date: date !== undefined ? date.toDate() : new Date()
                            }}
                        />
                    );
                }}
                refreshControl={
                    <RefreshControl
                        refreshing={isLoading}
                        onRefresh={() => {
                            void fetchAndFilter();
                        }}
                    />
                }
            />
        </View>
    );
}
