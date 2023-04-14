import {NativeBaseProvider} from "native-base";
import Article from "./Article";

interface ArticleInput {
    interest: string
    title: string
    description: string
    url: string
    imageUrl: string
}

const data: Array<ArticleInput> = [
    {
        interest: 'Metallica',
        title: 'Walking Dead\'s Norman Reedus reveals Death Stranding 2 existence - Digital Spy',
        description: 'Test description',
        url: 'https://www.digitalspy.com/tech/a40064850/walking-dead-norman-reedus-death-stranding-2/',
        imageUrl: 'https://hips.hearstapps.com/digitalspyuk.cdnds.net/16/28/1468254183-screen-shot-2016-07-11-at-171152.jpg?crop=1xw:0.8929577464788733xh;center,top&resize=1200:*'
    }
]

export default function Main() {
    return (
        <NativeBaseProvider>
            {data.map((article: ArticleInput) => {
                return <Article article={article}></Article>
            })}
        </NativeBaseProvider>
    )
}
