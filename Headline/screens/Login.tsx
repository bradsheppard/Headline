import * as React from 'react';
import { StatusBar } from 'expo-status-bar';
import { type NativeStackScreenProps } from '@react-navigation/native-stack';
import type { RootStackParamList } from '../App';
import { StyleSheet, Text, Image, TextInput, TouchableOpacity, View } from 'react-native';
import GoogleLogo from '../assets/google_logo';

import * as WebBrowser from 'expo-web-browser';
import * as Google from 'expo-auth-session/providers/google';
import { type UserInfo, useStore } from '../store';

WebBrowser.maybeCompleteAuthSession();

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

export default function Login(props: Props): JSX.Element {
    const [token, setToken, setUserInfo] = useStore((state) => [
        state.token,
        state.setToken,
        state.setUserInfo,
    ]);

    const [, response, promptAsync] = Google.useAuthRequest({
        clientId: '769516741930-9r4ush4ft2j15c3ng1d49idvd3qthb9r.apps.googleusercontent.com',
        androidClientId: '769516741930-t24q0a64288iahilf2t558irfn8vbvim.apps.googleusercontent.com',
    });

    React.useEffect(() => {
        if (response?.type === 'success') {
            const accessToken = response.authentication?.accessToken;

            if (accessToken === undefined) return;

            void getUserInfo(accessToken);
        }
    }, [response, token]);

    const getUserInfo = async (token: string): Promise<void> => {
        if (token === null) {
            return;
        }

        const response = await fetch('https://www.googleapis.com/userinfo/v2/me', {
            headers: { Authorization: `Bearer ${token}` },
        });

        const user = await response.json();

        const userInfo: UserInfo = {
            name: user.name,
            image: user.picture,
        };

        setUserInfo(userInfo);
        setToken(token);
    };

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
            <TouchableOpacity
                style={styles.loginBtn}
                onPress={() => {
                    void promptAsync();
                }}
            >
                <View style={styles.googleLogo}>
                    <GoogleLogo />
                </View>
                <Text>Sign in with Google</Text>
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
        width: 200,
    },
    inputView: {
        backgroundColor: '#C0DDFF',
        borderRadius: 30,
        width: '70%',
        height: 45,
        marginBottom: 20,
        alignItems: 'center',
    },
    googleLogo: {
        position: 'absolute',
        left: 20,
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
    },
});
