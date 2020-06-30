import React from 'react'
import { StyleSheet, Text, View } from 'react-native'

export default function Repos() {
    return (
        <View>
            <Text style={styles.text}>Code goes here</Text>
        </View>
    )
}

const styles = StyleSheet.create({
    text: {
        color: "white"
    }
})