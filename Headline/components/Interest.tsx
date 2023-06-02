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
        <Center w="64" h="10" bg="indigo.300" rounded="md" shadow={3}>
            <Text>{props.name}</Text>
            <Button 
                rounded="md" 
                position="absolute" 
                right="5" 
                size="xs"
                onPress={() => deleteInterest(props.id)}
            >
                -
            </Button>
        </Center>
    )
}

function deleteInterest(id: number) {
    InterestService.deleteInterest(id)
}

export default Interest;
