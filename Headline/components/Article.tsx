import React from "react"
import {Image, Platform, Text, TouchableOpacity, StyleSheet} from "react-native"
import {LinearGradient} from 'expo-linear-gradient'

type Article = {
    title: string
    url: string
    imageUrl: string
    description: string
}

interface Props {
    article: Article
}

const boxShadow: any = Platform.select({
  ios: {
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 0,
    },
    shadowOpacity: 0.4,
    shadowRadius: 4,
  },
  android: {elevation: 6},
});

const styles = StyleSheet.create({
    image: {
        flex: 1,
        borderRadius: 24,
        height: 300
    },
    container: {
        height: 240,
        marginBottom: 18,
        backgroundColor: '#eee',
        borderRadius: 24,
        marginHorizontal: 16,
        ...boxShadow,
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
});

const Article: React.FC<Props> = (props: Props) => {
    return (
        <TouchableOpacity activeOpacity={1} style={styles.container}>
            <Image 
                source={{uri: props.article.imageUrl}} 
                resizeMode={'cover'}
                style={styles.image}
            />
            <LinearGradient
                colors={['#0000', '#000A', '#000']}
                style={styles.title}>
                <Text style={styles.text}>{props.article.title}</Text>
                <Text style={styles.timestamp}>May 5, 2023</Text>
            </LinearGradient>
        </TouchableOpacity>
    )
}

export default Article
