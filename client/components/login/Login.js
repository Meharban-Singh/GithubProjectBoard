import React from 'react'
import { StyleSheet, Text, View, Button } from 'react-native'

export default function Home() {
    return (
        <View>
            <Text style={styles.title}>GitHub Project Boards</Text>
            <Button
                title="Login with gitHub"
                color="#FF5733"
                //onPress= go to Repos page
            />
        </View>
    )
}

const styles = StyleSheet.create({
    title: {
        paddingTop: 150,
        color: "#87C71A",
        textAlign: "center",
        fontSize: 82,
        fontWeight: "bold",
        paddingBottom: 100
    }
}) 