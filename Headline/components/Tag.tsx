import {StyleSheet, Text, TouchableOpacity} from "react-native"

interface Props {
    name: string
}

const styles = StyleSheet.create({
  container: {
    borderWidth: 1,
    borderColor: '#FFA500',
    borderRadius: 12,
    height: 28,
    paddingHorizontal: 14,
    justifyContent: 'center',
    alignItems: 'center',
    marginRight: 24,
    backgroundColor: '#FFA50066',
  },
  text: {
    fontSize: 14,
    fontWeight: '500',
  },
  selected: {
    backgroundColor: '#FF8800',
    borderColor: '#FF6600',
  },
})

const Tag: React.FC<Props> = (props: Props) => {
    return (
        <TouchableOpacity 
            style={[styles.container]}>
            <Text style={styles.text}>
                {props.name}
            </Text>
        </TouchableOpacity>
    )
}

export default Tag
