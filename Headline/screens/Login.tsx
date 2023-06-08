import * as React from 'react';
import { StatusBar } from 'expo-status-bar';
import { type NativeStackScreenProps } from '@react-navigation/native-stack';
import type { RootStackParamList } from '../App';
import { StyleSheet, Text, Image, TextInput, TouchableOpacity, View } from 'react-native';

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

export default function Login(props: Props): JSX.Element {
    return (
        <View style={styles.container}>
            <Image style={styles.image} source={require('../assets/HL.png')} />
            <StatusBar style='auto' />
            <View style={styles.inputView}>
                <TextInput
                    style={styles.TextInput}
                    placeholder='Email.'
                    placeholderTextColor='#003f5c'
                />
            </View>
            <View style={styles.inputView}>
                <TextInput
                    style={styles.TextInput}
                    placeholder='Password.'
                    placeholderTextColor='#003f5c'
                    secureTextEntry={true}
                />
            </View>
            <TouchableOpacity>
                <Text style={styles.forgot_button}>Forgot Password?</Text>
            </TouchableOpacity>
            <TouchableOpacity style={styles.loginBtn} onPress={() => {props.navigation.navigate('Main')}}>
                <Text>LOGIN</Text>
            </TouchableOpacity>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
    },
    image: {
        marginBottom: 40,
        height: 200,
        width: 200
    },
    inputView: {
        backgroundColor: '#C0DDFF',
        borderRadius: 30,
        width: '70%',
        height: 45,
        marginBottom: 20,
        alignItems: 'center',
    },
    TextInput: {
        height: 50,
        flex: 1,
        padding: 10,
        marginLeft: 20,
    },
    forgot_button: {
        height: 30,
        marginBottom: 30,
    },
    loginBtn: {
        width: '80%',
        borderRadius: 25,
        height: 50,
        alignItems: 'center',
        justifyContent: 'center',
        marginTop: 40,
        backgroundColor: '#14AAFF',
    }
});
