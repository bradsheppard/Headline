import * as React from 'react';
import { Button, FlatList, RefreshControl, View } from 'react-native';
import uuid from 'react-native-uuid';
import Tags from '../components/Tags';
import { useEffect, useState } from 'react';
import { type NativeStackScreenProps } from '@react-navigation/native-stack';
import { useStore } from '../store';
import { type Article as ArticleProto } from '../proto/article/article_pb';
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

interface ParamList {
    Interests: undefined;
}

type Props = NativeStackScreenProps<ParamList, 'Interests'>;

export default function Main(props: Props): JSX.Element {
    const [articles, interests] = useStore((state) => [state.articles, state.interests]);
    const [fetchArticles, fetchInterests] = useStore((state) => [
        state.fetchArticles,
        state.fetchInterests,
    ]);
    const selectedInterest = useStore((state) => state.selectedInterest);

    const [filteredArticles, setFilteredArticles] = useState<ArticleProto[]>([]);
    const [isLoading, setIsLoading] = useState(false);

    const fetchAndFilter = async (): Promise<void> => {
        await fetchData();
        filterArticles();
    };

    const fetchData = async (): Promise<void> => {
        setIsLoading(true);
        await Promise.all([fetchArticles(1), fetchInterests(1)]);
        setIsLoading(false);
    };

    const filterArticles = (): void => {
        if (selectedInterest !== null) {
            setFilteredArticles(articles.filter((x) => x.getInterest() === selectedInterest));
        } else {
            setFilteredArticles(articles);
        }
    };

    useEffect(() => {
        void fetchData();
    }, []);

    useEffect(() => {
        filterArticles();
    }, [selectedInterest]);

    return (
        <View style={[styles.container]}>
            <Tags interests={interests.map((x) => x.getName())} />
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
                                interest: entry.item.getInterest(),
                                title: entry.item.getTitle(),
                                url: entry.item.getUrl(),
                            }}
                        />
                    );
                }}
                refreshControl={
                    <RefreshControl
                        refreshing={isLoading}
                        onRefresh={() => {
                            void fetchAndFilter;
                        }}
                    />
                }
            />
            <Button
                title='Interests'
                onPress={() => {
                    props.navigation.navigate('Interests');
                }}
            />
        </View>
    );
}
