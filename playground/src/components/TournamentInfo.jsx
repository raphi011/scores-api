
import React from "react";
import styled from "styled-components";
import * as Theme from "../theme";
import Button from "./Button";

const TournamentInfo = () => (
    <Container>
        <Details>
            <Title>Amateur Herren Eibesbrunn</Title>
            <Subtitle>Double-Elimination</Subtitle>
            <Button>Signup, 2 weeks remaining</Button>
        </Details>
        {/* <Maps src="eibesbrunn.png" /> */}
    </Container>
);


const Details = styled.div`
    padding: 10px;
`

const Title = styled.h1`
    font-size: 30px;
    font-weight: 400;
    margin-bottom: 5px;
`

const Subtitle = styled.span`
    font-weight: 100;
    color: #6B787E;
`

const Maps = styled.img`
    width: 100%;
    border-radius: 0 0 5px 5px;
`

const Container = styled.header`
    background-color: #fff;
    margin: 20px 5px;
    border-radius: 5px;
    display: flex;
    justify-content: space-between;
    flex-direction: column;
`

export default TournamentInfo;
