import { Button, Center, Heading, NativeBaseProvider, VStack } from "native-base"
import { NavigationProp } from "@react-navigation/native"
import {NativeStackScreenProps} from "@react-navigation/native-stack";

type ParamList = {
    Main: undefined;
}

type Props = NativeStackScreenProps<ParamList, 'Main'>;

export default function Login(props: Props) {
    return (
        <NativeBaseProvider>
            <Center
                _dark={{ bg: "blueGray.900" }}
                _light={{ bg: "blueGray.50" }}
                px={4}
                flex={1}
            >
                <VStack space={5} alignItems="center">
                    <Heading size="lg">Headline</Heading>
                    <Button 
                        shadow={2} 
                        size="lg"
                        onPress={() => props.navigation.navigate('Main')}>Login</Button>
                </VStack>
            </Center>
        </NativeBaseProvider>
    )
}

