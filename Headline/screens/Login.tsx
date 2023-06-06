import * as React from 'react';
import { Button, Center, Heading, VStack } from 'native-base';
import { type NativeStackScreenProps } from '@react-navigation/native-stack';
import type { RootStackParamList } from '../App';

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

export default function Login(props: Props): JSX.Element {
    return (
        <Center _dark={{ bg: 'blueGray.900' }} _light={{ bg: 'blueGray.50' }} px={4} flex={1}>
            <VStack space={5} alignItems='center'>
                <Heading size='lg'>Headline</Heading>
                <Button
                    shadow={2}
                    size='lg'
                    onPress={() => {
                        props.navigation.navigate('Main');
                    }}
                >
                    Login
                </Button>
            </VStack>
        </Center>
    );
}
