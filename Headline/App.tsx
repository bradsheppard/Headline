import {NavigationContainer} from '@react-navigation/native'
import { createNativeStackNavigator } from '@react-navigation/native-stack'
import Interests from './screens/Interests'
import Login from './screens/Login'
import Main from './screens/Main'

type RootStackParamList = {
    Login: undefined;
    Main: undefined;
    Interests: undefined;
}

const Stack = createNativeStackNavigator<RootStackParamList>()

function App() {
    return (
        <NavigationContainer>
            <Stack.Navigator initialRouteName='Login'>
                <Stack.Screen name='Login' component={Login} />
                <Stack.Screen name='Main' component={Main} />
                <Stack.Screen name='Interests' component={Interests} />
            </Stack.Navigator>
        </NavigationContainer>
    )
}

export default App;

