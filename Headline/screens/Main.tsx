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
    const [articles, topics, selectedTopic] = useStore((state) => [state.articles, state.topics, state.selectedTopic]);
    const [fetchArticles, fetchTopics] = useStore((state) => [
        state.fetchArticles,
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
        await fetchTopics(1);
        await fetchArticles(topics.map(x => x.getName()))
        setIsLoading(false);
    };

    const filterArticles = (): void => {
        if (selectedTopic !== null) {
            const selectedArticles = articles.get(selectedTopic);

            if (selectedArticles === undefined) {
                setFilteredArticles([])
                return
            }

            setFilteredArticles(selectedArticles);
        } else {
            let allArticles: ArticleProto[] = []
            
            for(const [, value] of articles) {
                allArticles = allArticles.concat(value)
            }

            setFilteredArticles(allArticles);
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
                    return (
                        <Article
                            article={{
                                description: entry.item.getDescription(),
                                imageUrl: entry.item.getImageurl(),
                                title: entry.item.getTitle(),
                                url: entry.item.getUrl(),
                                source: entry.item.getSource(),
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
