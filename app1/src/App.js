import {
  Container,
  Navbar,
  Nav,
  NavDropdown,
  FormControl,
  Form,
  Button,
} from "react-bootstrap";
import logo from "./logo1.svg";
import axios from "axios";
import { useState, useEffect } from "react";

function App() {
  const [coins, setCoins] = useState([]);
  const getCoins = async () => {
    const coins = await axios.get("http://20.69.90.222:30008/albums");
    setCoins(coins.data);
    console.log(coins);
  };
  return (
    <>
      <Container>
        <Navbar bg="dark" variant="dark">
          <Container>
            <Navbar.Brand href="#home">
              <img
                alt=""
                src={logo}
                width="30"
                height="30"
                className="d-inline-block align-top"
              />
              Rythm
            </Navbar.Brand>
            <Nav className="me-auto">
              <Nav.Link href="#home">Home</Nav.Link>
              <Nav.Link href="#features">Features</Nav.Link>
              <Nav.Link href="#pricing">Pricing</Nav.Link>
              <NavDropdown title="Dropdown" id="collasible-nav-dropdown">
                <NavDropdown.Item href="#action/3.1">Action</NavDropdown.Item>
                <NavDropdown.Item href="#action/3.2">
                  Another action
                </NavDropdown.Item>
                <NavDropdown.Item href="#action/3.3">
                  Something
                </NavDropdown.Item>
                <NavDropdown.Divider />
                <NavDropdown.Item href="#action/3.4">
                  Separated link
                </NavDropdown.Item>
              </NavDropdown>
            </Nav>
            <Form className="d-flex">
              <FormControl
                type="search"
                placeholder="Search"
                className="me-2"
                aria-label="Search"
              />
              <Button variant="outline-success">Search</Button>
            </Form>
          </Container>
        </Navbar>
        <div>
          <Button onClick={() => getCoins()}>get coins</Button>
        </div>
        <br />
        {coins.map((coin) => (
          <div key={coin.id}>
            {coin.title}
            <br />
          </div>
        ))}
      </Container>
    </>
  );
}

export default App;
