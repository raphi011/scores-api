import React from "react";
import logo from "../logo.svg";
import styled from "styled-components";

const NavBar = () => (
  <header>
    <Nav>
        <Logo src={logo} alt="logo" />
        <Title>Scores</Title>
    </Nav>
  </header>
);

const Title = styled.h1`
    font-size: 32px;
    font-weight: 200;
    color: #96CEB4;
`

const Logo = styled.img`
    height: 32px;
    padding: 10px;
`

const Nav = styled.nav`
    background-color: #fff;
    height: 60px;
    display: flex;
    align-items: center;
`

export default NavBar;
