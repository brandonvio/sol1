import axios from 'axios'
import { Container, Button } from 'react-bootstrap'
import AppNav from './AppNav'
import { useState, useEffect } from 'react'

function App() {
    useEffect(() => {
        console.log('use-effect')
    })
    const [coins, setCoins] = useState([])
    const getCoins = async () => {
        const coins = await axios.get('http://20.69.90.222:30008/albums')
        setCoins(coins.data)
        console.log(coins)
    }
    return (
        <>
            <Container>
                <AppNav />
                <br />
                <div>
                    <Button onClick={() => getCoins()}>get coins</Button>
                </div>
                <br />
                <div>
                    {coins.map((coin) => (
                        <div key={coin.id}>
                            {coin.title}
                            <br />
                        </div>
                    ))}
                </div>
            </Container>
        </>
    )
}

export default App
