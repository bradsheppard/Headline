import * as React from 'react';
import { useTheme } from 'native-base';
import { StyleSheet, Text, TouchableOpacity } from 'react-native';
import { useStore } from '../store';

interface Props {
    name: string;
}

const createStyles = (color: string): StyleSheet.NamedStyles<any> =>
    StyleSheet.create({
        container: {
            borderWidth: 1,
            borderColor: color,
            backgroundColor: color,
            borderRadius: 12,
            height: 28,
            paddingHorizontal: 14,
            justifyContent: 'center',
            alignItems: 'center',
            marginRight: 24,
        },
        selected: {
            backgroundColor: '#FF8800',
            borderColor: '#FF6600',
        },
    });

const Tag: React.FC<Props> = (props: Props) => {
    const [selectedInterest, setSelectedInterest] = useStore((store) => [
        store.selectedInterest,
        store.setSelectedInterest,
    ]);

    const { colors } = useTheme();

    const isSelected = selectedInterest === props.name;
    let styles;

    if (isSelected) {
        styles = createStyles(colors.warning[400]);
    } else {
        styles = createStyles(colors.info[400]);
    }

    return (
        <TouchableOpacity
            onPress={() => {
                setSelectedInterest(props.name);
            }}
            style={[styles.container]}
        >
            <Text>{props.name}</Text>
        </TouchableOpacity>
    );
};

export default Tag;
