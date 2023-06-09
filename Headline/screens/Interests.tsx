import * as React from 'react';
import { useEffect, useState } from 'react';
import { View, TouchableOpacity, StyleSheet } from 'react-native';
import { Button as NativeBaseButton, FormControl, Input, Modal, VStack } from 'native-base';
import InterestService from '../api/interest';
import Interest from '../components/Interest';
import { useStore } from '../store';
import { Svg, Rect, Line } from 'react-native-svg';

const plus = (): JSX.Element => {
    return (
        <Svg
            enable-background='new 0 0 50 50'
            height='50px'
            id='Layer_1'
            viewBox='0 0 50 50'
            width='50px'
        >
            <Rect fill='none' height='50' width='50' />
            <Line
                fill='none'
                stroke='#030D45'
                strokeMiterlimit='10'
                strokeWidth='3'
                x1='12'
                x2='38'
                y1='25'
                y2='25'
            />
            <Line
                fill='none'
                stroke='#030D45'
                strokeMiterlimit='10'
                strokeWidth='3'
                x1='25'
                x2='25'
                y1='12'
                y2='38'
            />
        </Svg>
    );
};

const styles = StyleSheet.create({
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
    button: {
        position: 'absolute',
        right: 20,
        bottom: 20,
        backgroundColor: '#14AAFF',
        borderRadius: 30,
    },
});

export default function Interests(): JSX.Element {
    const [interests, fetchInterests] = useStore((state) => [
        state.interests,
        state.fetchInterests,
    ]);
    const [modalVisible, setModalVisible] = useState(false);
    const [inputText, setInputText] = useState('');

    const fetchData = async (): Promise<void> => {
        await fetchInterests(1);
    };

    const createInterest = async (): Promise<void> => {
        await InterestService.createInterest(inputText);
        await fetchData();
    };

    useEffect(() => {
        void fetchData();
    }, []);

    return (
        <View style={styles.container}>
            <VStack space={4} alignItems='center' style={styles.list}>
                {interests.map((interest) => {
                    return (
                        <Interest
                            key={interest.getId()}
                            id={interest.getId()}
                            name={interest.getName()}
                        />
                    );
                })}
            </VStack>
            <TouchableOpacity
                style={styles.button}
                onPress={() => {
                    setModalVisible(true);
                }}
            >
                {plus()}
            </TouchableOpacity>
            <Modal isOpen={modalVisible}>
                <Modal.Content>
                    <Modal.CloseButton />
                    <Modal.Header>New Interest</Modal.Header>
                    <Modal.Body>
                        <FormControl>
                            <FormControl.Label>Interest</FormControl.Label>
                            <Input
                                onChangeText={(text) => {
                                    setInputText(text);
                                }}
                            />
                        </FormControl>
                    </Modal.Body>
                    <Modal.Footer>
                        <NativeBaseButton.Group space={2}>
                            <NativeBaseButton
                                variant='ghost'
                                colorScheme='blueGray'
                                onPress={() => {
                                    setModalVisible(false);
                                }}
                            >
                                Cancel
                            </NativeBaseButton>
                            <NativeBaseButton
                                onPress={() => {
                                    setModalVisible(false);
                                    void createInterest();
                                }}
                            >
                                Save
                            </NativeBaseButton>
                        </NativeBaseButton.Group>
                    </Modal.Footer>
                </Modal.Content>
            </Modal>
        </View>
    );
}
