import { Button, Center, Heading, NativeBaseProvider, VStack } from "native-base"

export default function App() {
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
                    <Button shadow={2} size="lg">Login</Button>
                </VStack>
            </Center>
        </NativeBaseProvider>
    )
}
