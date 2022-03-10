import axios from 'axios'
import { Card, Container, Button, CardGroup } from 'react-bootstrap'
import AppNav from './AppNav'
import { useState, useEffect } from 'react'

function App() {
    const [prices, setPrices] = useState([])
    const getPrices = async () => {
        const apiBase = process.env.REACT_APP_API_BASE;
        const result = await axios.get(`${apiBase}/prices`)
        setPrices(result.data)
        console.log(result)
    }
    useEffect(async () => {
        await getPrices();
    },[])
    return (
        <>
            <Container>
                <AppNav />
                <br />
                <CardGroup>
                    {prices.map((price) => (
                        <div key={price.symbol}>
                            <Card border="success" style={{ width: '18rem' }}>
                                <Card.Body>
                                    <Card.Title>{price.symbol}</Card.Title>
                                    <Card.Text>
                                        {price.price}
                                    </Card.Text>
                                    <Button variant="info">Analysis</Button>
                                </Card.Body>
                            </Card>
                        </div>
                    ))}
                </CardGroup>
            </Container>
        </>
    )
}

export default App
