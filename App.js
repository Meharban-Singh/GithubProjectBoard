import React from 'react'
import { StyleSheet, Text, View, Button } from 'react-native'

export default function App() {
    return (
        <View style={styles.container}>
            <Text style={styles.title}>GitHub Project Boards</Text>
            <Button
                title="Login with gitHub"
            ></Button>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 24,
        backgroundColor: "#20232a"
    },
    title: {
        paddingTop: 150,
        color: "#87C71A",
        textAlign: "center",
        fontSize: 82,
        fontWeight: "bold",
        paddingBottom: 100
    }
})