import * as React from 'react';
import { Button, Center, Text } from 'native-base';
import { useStore } from '../store';

interface Props {
    id: number;
    name: string;
}

const Interest: React.FC<Props> = (props: Props) => {
    const deleteInterest = useStore((state) => state.deleteInterest);

    return (
        <Center w='64' h='10' bg='info.400' rounded='xl' shadow={3}>
            <Text>{props.name}</Text>
            <Button
                bgColor='red.700'
                rounded='md'
                position='absolute'
                right='5'
                size='sm'
                onPress={() => {
                    void deleteInterest(props.id);
                }}
            >
                -
            </Button>
        </Center>
    );
};

export default Interest;
