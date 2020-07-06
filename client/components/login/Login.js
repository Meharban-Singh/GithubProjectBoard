import React from 'react'
import { StyleSheet, Text, View, Button } from 'react-native'

function handleLogin() {
    fetch('http://192.168.100.105:1010/login')
        .then(res => {
            if(res.status === 200) {
                console.log('SUCCESS')
                //TODO: Redirect to repos page on client side
            }
            else 
                console.log('CANNOT LOGIN')
        })
        .catch(err => console.log(err))
}

export default function Home() {
    return (
        <View>
            <Text style={styles.title}>GitHub Project Boards</Text>
            <Button
                title="Login with gitHub"
                color="#FF5733"
                onPress= {handleLogin}
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