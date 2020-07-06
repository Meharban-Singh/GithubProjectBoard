import React from 'react'
import { StyleSheet, View } from 'react-native'

import Login from './components/login/Login'

export default function App() {
    return (
        <View style={styles.container}>
            <Login />
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