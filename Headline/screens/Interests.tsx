import {useEffect, useState} from "react";
import {Button, FlatList, Text, View} from "react-native";
import {Button as NativeBaseButton} from 'native-base'
import InterestService from "../api/interest";
import uuid from 'react-native-uuid'
import Interest from "../components/Interest";
import {FormControl, Input, Modal, VStack} from "native-base";
import {Interest as InterestProto} from "../proto/interest/interest_pb";
import {useStore} from "../store";

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
    const [interests, fetchInterests] = useStore((state) => [state.interests, state.fetchInterests])
    const [modalVisible, setModalVisible] = useState(false)
    const [inputText, setInputText] = useState('')

    const fetchData = async () => {
        await fetchInterests(1)
    }

    const createInterest = async () => {
        await InterestService.createInterest(inputText)
        await fetchData()
    }

    useEffect(() => {
        fetchData()
    }, [])

    return (
        <View style={styles.container}>
            <VStack space={4} alignItems="center" style={styles.list}>
                {interests.map(interest => {
                    return <Interest key={interest.getId()} id={interest.getId()} name={interest.getName()} />
                })}
            </VStack>
            <Button title="+" onPress={() => setModalVisible(true)}/>
            <Modal isOpen={modalVisible}>
                <Modal.Content>
                    <Modal.CloseButton />
                    <Modal.Header>New Interest</Modal.Header>
                    <Modal.Body>
                        <FormControl>
                            <FormControl.Label>Interest</FormControl.Label>
                            <Input onChangeText={text => setInputText(text)}/>
                        </FormControl>
                    </Modal.Body>
                    <Modal.Footer>
                        <NativeBaseButton.Group space={2}>
                            <NativeBaseButton variant="ghost" colorScheme="blueGray" onPress={() => {
                                setModalVisible(false);
                            }}>
                                Cancel
                            </NativeBaseButton>
                            <NativeBaseButton onPress={() => {
                                setModalVisible(false);
                                createInterest()
                            }}>
                                Save
                            </NativeBaseButton>
                        </NativeBaseButton.Group>
                    </Modal.Footer>
                </Modal.Content>
            </Modal>
        </View>
    )
}

