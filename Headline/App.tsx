import * as React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { NativeBaseProvider } from 'native-base';
import Interests from './screens/Interests';
import Login from './screens/Login';
import Main from './screens/Main';

interface RootStackParamList {
    [key: string]: object;
    Login: object;
    Main: object;
    Interests: object;
}

const Stack = createNativeStackNavigator<RootStackParamList>();

function App(): JSX.Element {
    return (
        <NativeBaseProvider>
            <NavigationContainer>
                <Stack.Navigator initialRouteName='Login'>
                    <Stack.Screen name='Login' component={Login} />
                    <Stack.Screen name='Main' component={Main} />
                    <Stack.Screen name='Interests' component={Interests} />
                </Stack.Navigator>
            </NavigationContainer>
        </NativeBaseProvider>
    );
}

export default App;
