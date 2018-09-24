import React from "react";
import logo from "../logo.svg";
import styled from "styled-components";
import TournamentList from "./TournamentList";
import TournamentInfo from "./TournamentInfo";
import NavBar from "./NavBar";
import TeamList from "./TeamList";

const Layout = () => (
    <Container>
        <NavBar />
        <SplitView>
            <Sidebar>
                <TournamentList />
            </Sidebar>
            <Content>
                <TournamentInfo />
                <Header>3 / 16 teams</Header>
                <TeamList />
            </Content>
        </SplitView>
    </Container>
);

const Container = styled.div`
    height: 100%;
    width: 100%;
    position: relative;
`

const SplitView = styled.div`
    display: flex;
    flex-direction: row;
    position: absolute;
`

const Sidebar = styled.div`
    width: 320px;
    overflow-x: scroll;
`

const Content = styled.main`
    flex-grow: 1;
`

const Header = styled.h2`
  text-align: center;
  font-size: 24px;
  font-weight: 300;
  margin-bottom: 5px;
  color: #222;
`

export default Layout;

