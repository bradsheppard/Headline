import {NativeBaseProvider} from "native-base";
import {FlatList, Text} from "react-native";
import Article from "./Article";
import uuid from 'react-native-uuid'
import {View} from "react-native";
import Tags from "./Tags";
import InterestService from "../api/interest";
import {useEffect, useState} from "react";
import ArticleService from "../api/article";

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

export default function Main() {
    const [interests, setInterests] = useState<string[]>([]);
    const [articles, setArticles] = useState<Article[]>([]);
    const [selectedInterest, setSelectedInterest] = useState<string | null>(null)

    const fetchData = async () => {
        const [interestResponse, articleResponse] = await Promise.all([
            InterestService.getInterests(1),
            ArticleService.getArticles(1)
        ])

        setArticles(articleResponse);
        setInterests(interestResponse);
    }

    useEffect(() => {
        fetchData()
    }, [])

    return (
        <View style={[styles.container]}>
            <Tags interests={interests} setSelectedInterest={setSelectedInterest} />
            <FlatList 
                style={styles.list}
                data={articles}
                keyExtractor={() => uuid.v4() as string}
                renderItem={({item, index}: any) => (
                    <Article article={item} />
                )}
            />
        </View>
    )
}
