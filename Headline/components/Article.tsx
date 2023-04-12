import React from "react"
import {Image, TouchableOpacity} from "react-native"

type Article = {
    title: string
    url: string
    description: string
    imageUrl: string
}

interface Props {
    article: Article
}

const style = {
    flex: 1,
    borderRadius: 24,
    height: 300
}

const Article: React.FC<Props> = (props: Props) => {
    return (
        <TouchableOpacity>
            <Image 
                source={{uri: props.article.imageUrl}} 
                resizeMode={'cover'}
                style={style}
            />
        </TouchableOpacity>
    )
}

export default Article
