import * as React from 'react';
import { type ListRenderItemInfo, StyleSheet, FlatList } from 'react-native';
import Tag from './Tag';

interface Props {
    topics: string[];
}

const style = StyleSheet.create({
    list: {
        maxHeight: 40,
    },
    contentContainer: {
        paddingHorizontal: 24,
        alignItems: 'center',
        justifyContent: 'center',
    },
});

const Tags: React.FC<Props> = (props: Props) => {
    return (
        <FlatList
            horizontal
            data={props.topics}
            showsHorizontalScrollIndicator={false}
            keyExtractor={(item: string) => item}
            renderItem={({ item }: ListRenderItemInfo<string>) => <Tag name={item} />}
            style={style.list}
            contentContainerStyle={style.contentContainer}
        />
    );
};

export default Tags;
