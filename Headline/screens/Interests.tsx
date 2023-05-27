import {useEffect, useState} from "react";
import {FlatList, Text, View} from "react-native";
import InterestService from "../api/interest";
import uuid from 'react-native-uuid'

const styles = {
    list: {
        flex: 1,
        flexGrow: 1,
        paddingVertical: 8,
    },
    container: {
        flex: 1,
        flexGrow: 1,
        backgroundColor: '#000',
    },
    text: {
        fontSize: 18,
        lineHeight: 24,
        color: '#fff',
        paddingBottom: 24,
    },
}

export default function Interests() {
    const [interests, setInterests] = useState<string[]>([]);

    const fetchData = async () => {
        const interestResponse = await InterestService.getInterests(1)
        setInterests(interestResponse);
    }

    useEffect(() => {
        fetchData()
    }, [])

    return (
        <View style={styles.container}>
            <FlatList 
                style={styles.list}
                data={interests}
                keyExtractor={() => uuid.v4() as string}
                renderItem={({item}) => {
                    return <Text style={styles.text}>{item}</Text>
                }}
            />
        </View>
    )
}
