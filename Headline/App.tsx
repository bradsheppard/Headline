import * as React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import {
    createNativeStackNavigator,
    type NativeStackNavigationOptions,
} from '@react-navigation/native-stack';
import { NativeBaseProvider } from 'native-base';
import Interests from './screens/Interests';
import Login from './screens/Login';
import Main from './screens/Main';
import { useStore } from './store';
import {type HeaderButtonProps} from '@react-navigation/native-stack/lib/typescript/src/types';
import {Image, TouchableOpacity} from 'react-native';

// eslint-disable-next-line
export type RootStackParamList = {
    Login: undefined;
    Main: undefined;
    Interests: undefined;
};

const Stack = createNativeStackNavigator<RootStackParamList>();


function App(): JSX.Element {
    const [token, setToken, userInfo ] = useStore((state) => [state.token, state.setToken, state.userInfo]);

    const screenOptions: NativeStackNavigationOptions = {
        headerStyle: {
            backgroundColor: 'black',
        },
        headerTintColor: 'white',
        headerTitleStyle: {
            fontWeight: 'bold',
        },
        statusBarStyle: 'light',
        statusBarColor: 'black',

        headerRight: (props: HeaderButtonProps) => {
            return (
                <TouchableOpacity onPress={() => {setToken(null)}}>
                    <Image 
                        style={{width: 35, height: 35, borderRadius: 35}} 
                        source={{uri: userInfo?.image}} 
                    />
                </TouchableOpacity>
            )
        }
    };

    return (
        <NativeBaseProvider>
            <NavigationContainer>
                <Stack.Navigator initialRouteName='Login'>
                    {token === null ? (
                        <Stack.Screen
                            name='Login'
                            component={Login}
                            options={{
                                headerShown: false,
                            }}
                        />
                    ) : (
                        <>
                            <Stack.Screen
                                name='Main'
                                component={Main}
                                options={{ ...screenOptions, ...{ headerTitle: 'Feed' } }}
                            />
                            <Stack.Screen
                                name='Interests'
                                component={Interests}
                                options={screenOptions}
                            />
                        </>
                    )}
                </Stack.Navigator>
            </NavigationContainer>
        </NativeBaseProvider>
    );
}

export default App;
