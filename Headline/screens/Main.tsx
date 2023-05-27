import {NativeBaseProvider} from "native-base";
import {Button, FlatList, Text} from "react-native";
import Article from "../components/Article";
import uuid from 'react-native-uuid'
import {View} from "react-native";
import Tags from "../components/Tags";
import InterestService from "../api/interest";
import {useEffect, useState} from "react";
import ArticleService from "../api/article";
import {NativeStackScreenProps} from "@react-navigation/native-stack";

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
    const [interests, setInterests] = useState<string[]>([]);
    const [articles, setArticles] = useState<Article[]>([]);
    const [selectedInterest, setSelectedInterest] = useState<string | null>(null)

    const fetchData = async () => {
        let [interestResponse, articleResponse] = await Promise.all([
            InterestService.getInterests(1),
            ArticleService.getArticles(1)
        ])

        if (selectedInterest !== null) {
            articleResponse = articleResponse.filter(x => x.interest === selectedInterest)
        }

        setArticles(articleResponse);
        setInterests(interestResponse);
    }

    useEffect(() => {
        fetchData()
    }, [])

    const filteredArticles = selectedInterest !== null ? articles.filter(x => x.interest === selectedInterest) : articles

    return (
        <View style={[styles.container]}>
            <Tags interests={interests} setSelectedInterest={setSelectedInterest} />
            <FlatList 
                style={styles.list}
                data={filteredArticles}
                keyExtractor={() => uuid.v4() as string}
                renderItem={({item, index}: any) => (
                    <Article article={item} />
                )}
            />
            <Button title="Interests" onPress={() => props.navigation.navigate('Interests')} />
        </View>
    )
}
