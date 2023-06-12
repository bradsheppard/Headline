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

// eslint-disable-next-line
export type RootStackParamList = {
    Login: undefined;
    Main: undefined;
    Interests: undefined;
};

const Stack = createNativeStackNavigator<RootStackParamList>();

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
};

function App(): JSX.Element {
    const token = useStore((state) => state.token);

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
