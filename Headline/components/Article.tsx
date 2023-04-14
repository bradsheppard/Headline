import React from "react"
import {Image, Text, TouchableOpacity} from "react-native"
import {LinearGradient} from "react-native-svg"

type Article = {
    title: string
    url: string
    imageUrl: string
    description: string
}

interface Props {
    article: Article
}

const styles = {
    image: {
        flex: 1,
        borderRadius: 24,
        height: 300
    },
    title: {
        position: 'absolute',
        bottom: 0,
        width: '100%',
        borderBottomLeftRadius: 24,
        borderBottomRightRadius: 24,
        height: 160,
        paddingLeft: 16,
        paddingRight: 10,
        justifyContent: 'flex-end',
        alignItems: 'flex-start',
    },
    text: {
        fontSize: 18,
        lineHeight: 24,
        color: '#fff',
        paddingBottom: 24,
    },
    timestamp: {
        position: 'absolute',
        color: '#eee',
        fontSize: 12,
        fontWeight: '300',
        right: 16,
        bottom: 8, 
    }
}

const Article: React.FC<Props> = (props: Props) => {
    return (
        <TouchableOpacity>
            <Image 
                source={{uri: props.article.imageUrl}} 
                resizeMode={'cover'}
                style={styles.image}
            />
            <LinearGradient
                colors={['#0000', '#000A', '#000']}
                style={styles.title}>
                <Text style={styles.text}>{props.article.title}</Text>
            </LinearGradient>
        </TouchableOpacity>
    )
}

export default Article
