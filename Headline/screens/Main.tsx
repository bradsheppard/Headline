import {NativeBaseProvider} from "native-base";
import {Button, FlatList, RefreshControl, Text} from "react-native";
import uuid from 'react-native-uuid'
import {View} from "react-native";
import Tags from "../components/Tags";
import InterestService from "../api/interest";
import {useEffect, useState} from "react";
import ArticleService from "../api/article";
import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {createStore} from "zustand";
import {useStore} from "../store";
import {Article as ArticleProto} from "../proto/article/article_pb";
import Article from "../components/Article";

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
    }
}

type ParamList = {
    Interests: undefined;
}

type Props = NativeStackScreenProps<ParamList, 'Interests'>;

export default function Main(props: Props) {
    const [articles, interests] = 
        useStore((state) => [state.articles, state.interests])
    const [fetchArticles, fetchInterests] = 
        useStore((state) => [state.fetchArticles, state.fetchInterests])
    const [selectedInterest, setSelectedInterest] = 
        useStore((state) => [state.selectedInterest, state.setSelectedInterest])

    const [filteredArticles, setFilteredArticles] = useState<Array<ArticleProto>>([])
    const [isLoading, setIsLoading] = useState(false)

    const fetchAndFilter = async() => {
        await fetchData()
        filterArticles()
    }

    const fetchData = async () => {
        setIsLoading(true)
        await Promise.all([fetchArticles(1), fetchInterests(1)])
        setIsLoading(false)
    }

    const filterArticles = () => {
        if (selectedInterest !== null) {
            setFilteredArticles(articles.filter(x => x.getInterest() === selectedInterest))
        }
        else {
            setFilteredArticles(articles)
        }
    }

    useEffect(() => {
        fetchData()
    }, [])

    useEffect(() => {
        filterArticles()
    }, [selectedInterest])

    return (
        <View style={[styles.container]}>
            <Tags interests={interests.map(x => x.getName())} />
            <FlatList 
                style={styles.list}
                data={filteredArticles}
                keyExtractor={() => uuid.v4() as string}
                renderItem={(entry) => {
                    return <Article article={
                        {
                            description: entry.item.getDescription(),
                            imageUrl: entry.item.getImageurl(),
                            interest: entry.item.getInterest(),
                            title: entry.item.getTitle(),
                            url: entry.item.getUrl()
                        } 
                    } />
                }}
                refreshControl={
                    <RefreshControl refreshing={isLoading} onRefresh={fetchAndFilter} />
                }
            />
            <Button title="Interests" onPress={() => props.navigation.navigate('Interests')} />
        </View>
    )
}
