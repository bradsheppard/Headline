import * as React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { NativeBaseProvider } from 'native-base';
import Interests from './screens/Interests';
import Login from './screens/Login';
import Feed from './screens/Main';
import { useStore } from './store';
import { Image, TouchableOpacity } from 'react-native';
import { MaterialCommunityIcons, AntDesign } from '@expo/vector-icons';

// eslint-disable-next-line
export type RootStackParamList = {
    Login: undefined;
    Feed: undefined;
    Interests: undefined;
};

const Stack = createNativeStackNavigator<RootStackParamList>();
const Tab = createBottomTabNavigator<RootStackParamList>();

function App(): JSX.Element {
    const [token, setToken, userInfo] = useStore((state) => [
        state.token,
        state.setToken,
        state.userInfo,
    ]);

    return (
        <NativeBaseProvider>
            <NavigationContainer>
                {token === null ? (
                    <Stack.Navigator initialRouteName='Login'>
                        <Stack.Screen
                            name='Login'
                            component={Login}
                            options={{
                                headerShown: false,
                            }}
                        />
                    </Stack.Navigator>
                ) : (
                    <Tab.Navigator
                        screenOptions={{
                            tabBarActiveBackgroundColor: 'black',
                            tabBarInactiveBackgroundColor: 'black',
                            headerStyle: {
                                backgroundColor: 'black',
                            },
                            headerTintColor: 'white',
                            headerTitleStyle: {
                                fontWeight: 'bold',
                            },
                            headerRight: () => {
                                return (
                                    <TouchableOpacity
                                        onPress={() => {
                                            setToken(null);
                                        }}
                                    >
                                        <Image
                                            style={{
                                                width: 35,
                                                height: 35,
                                                borderRadius: 35,
                                                margin: 10,
                                            }}
                                            source={{ uri: userInfo?.image }}
                                        />
                                    </TouchableOpacity>
                                );
                            },
                        }}
                    >
                        <Tab.Screen
                            name='Feed'
                            component={Feed}
                            options={{
                                tabBarIcon: ({ color, size }) => (
                                    <MaterialCommunityIcons name='home' color={color} size={size} />
                                ),
                            }}
                        />
                        <Tab.Screen
                            name='Interests'
                            component={Interests}
                            options={{
                                tabBarIcon: ({ color, size }) => (
                                    <AntDesign name='bars' color={color} size={size} />
                                ),
                            }}
                        />
                    </Tab.Navigator>
                )}
            </NavigationContainer>
        </NativeBaseProvider>
    );
}

export default App;
