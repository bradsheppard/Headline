import {Button, Center, NativeBaseProvider, Text, VStack} from "native-base";
import {ViewStyle} from "react-native";
import {StyleProp, TouchableOpacity} from "react-native"
import InterestService from "../api/interest";
import {useStore} from "../store";

interface Props {
    id: number
    name: string
}

const Interest: React.FC<Props> = (props: Props) => {
    const [interests, deleteInterest] = useStore((state) => [state.interests, state.deleteInterest])

    return(
        <Center w="64" h="10" bg="info.400" rounded="xl" shadow={3}>
            <Text>{props.name}</Text>
            <Button 
                bgColor="red.700"
                rounded="md" 
                position="absolute" 
                right="5" 
                size="sm"
                onPress={() => deleteInterest(props.id)}
            >
                -
            </Button>
        </Center>
    )
}

export default Interest;
