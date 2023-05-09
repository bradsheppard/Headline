import {useState} from "react";
import {FlatList, Text, View} from "react-native";
import InterestService from "../api/interest";

export default function Main() {
    const [interests, setInterests] = useState<string[]>([]);

    const fetchData = async () => {
        const interestResponse = await InterestService.getInterests(1)

        setInterests(interestResponse);
    }

    return (
        <View>
            <FlatList 
                data={interests}
                renderItem={({item}) => {
                    return <Text>item</Text>
                }}
            />
        </View>
    )
}
