import React from 'react'
import { StyleSheet, View } from 'react-native'

import Home from './Home.js'

export default function App() {
    return (
        <View style={styles.container}>
            <Home />
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 24,
        backgroundColor: "#20232a"
    }
})