import {ListRenderItemInfo, StyleSheet} from "react-native";
import {ListRenderItem} from "react-native";
import {FlatList} from "react-native";
import Tag from "./Tag";

interface Props {
    interests: Array<string>
    setSelectedInterest: (interest: string) => void;
}

const style = StyleSheet.create({
    list: {
        maxHeight: 40
    },
    contentContainer: {
        paddingHorizontal: 24,
        alignItems: 'center',
        justifyContent: 'center'
    }
})

const Tags: React.FC<Props> = (props: Props) => {
    return (
        <FlatList
            horizontal
            data={props.interests}
            showsHorizontalScrollIndicator={false}
            keyExtractor={(item: string) => item}
            renderItem={({item}: ListRenderItemInfo<string>) => (
                <Tag name={item} setSelectedInterest={props.setSelectedInterest}/>
            )}
            style={style.list}
            contentContainerStyle={style.contentContainer}
        />
    )
}

export default Tags;

