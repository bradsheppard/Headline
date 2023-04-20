import {NativeBaseProvider} from "native-base";
import {FlatList, Text} from "react-native";
import Article from "./Article";
import uuid from 'react-native-uuid'
import {View} from "react-native";
import Tags from "./Tags";

interface ArticleInput {
    interest: string
    title: string
    description: string
    url: string
    imageUrl: string
}

const data: Array<ArticleInput> = [
    {
        interest: 'The Walking Dead',
        title: 'Walking Dead\'s Norman Reedus reveals Death Stranding 2 existence - Digital Spy',
        description: 'Test description',
        url: 'https://www.digitalspy.com/tech/a40064850/walking-dead-norman-reedus-death-stranding-2/',
        imageUrl: 'https://hips.hearstapps.com/digitalspyuk.cdnds.net/16/28/1468254183-screen-shot-2016-07-11-at-171152.jpg?crop=1xw:0.8929577464788733xh;center,top&resize=1200:*'
    },
    {
        interest: 'Metallica',
        title: 'Metallica Releases "72 Seasons" Album',
        description: 'Test description',
        url: 'https://www.google.com',
        imageUrl: 'https://sleazeroxx.com/wp-content/uploads/Metallica-album-cover-e1669765574222.webp'
    },
    {
        interest: 'Software Engineering',
        title: 'Meta Set to Layoff More Engineers',
        description: 'Test description',
        url: 'https://www.google.com',
        imageUrl: 'https://akm-img-a-in.tosshub.com/indiatoday/images/story/202304/meta_0-sixteen_nine.jpg?VersionId=g.5vwaNxKtG9caVo29sDGHgndy0J4PKh&size=690:388'
    }
]

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
    const tags = [...new Set(data.map(x => x.interest))];

    return (
        <View style={[styles.container]}>
            <Tags interests={tags} />
            <FlatList 
                style={styles.list}
                data={data}
                keyExtractor={() => uuid.v4() as string}
                renderItem={({item, index}: any) => (
                    <Article article={item} />
                )}
            />
        </View>
    )
}
