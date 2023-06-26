import * as React from 'react';
import { useEffect, useState } from 'react';
import { View, TouchableOpacity, StyleSheet } from 'react-native';
import { Button as NativeBaseButton, FormControl, Input, Modal, VStack } from 'native-base';
import TopicService from '../api/topic';
import Topic from '../components/Topic';
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

export default function Topics(): JSX.Element {
    const [topics, fetchTopics] = useStore((state) => [
        state.topics,
        state.fetchTopics,
    ]);
    const [modalVisible, setModalVisible] = useState(false);
    const [inputText, setInputText] = useState('');

    const fetchData = async (): Promise<void> => {
        await fetchTopics(1);
    };

    const createTopic = async (): Promise<void> => {
        await TopicService.createTopic(inputText);
        await fetchData();
    };

    useEffect(() => {
        void fetchData();
    }, []);

    return (
        <View style={styles.container}>
            <VStack space={4} alignItems='center' style={styles.list}>
                {topics.map((topic) => {
                    return (
                        <Topic
                            key={topic.getName()}
                            name={topic.getName()}
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
                    <Modal.Header>New Topic</Modal.Header>
                    <Modal.Body>
                        <FormControl>
                            <FormControl.Label>Topic</FormControl.Label>
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
                                    void createTopic();
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
