import React from "react";
import styled from "styled-components";
import * as Theme from "../theme";

const exampleTeams = [
    { 
        rank: 1,
        totalPoints: 435,
        player1: { firstName: "Raphael", lastName: "Gruber", points: 220 },
        player2: { firstName: "Richard", lastName: "Bosse", points: 215 },
    },
    { 
        rank: 2,
        totalPoints: 425,
        player1: { firstName: "Lukas", lastName: "Wimmer", points: 200 },
        player2: { firstName: "Dominik", lastName: "Rieder", points: 225 },
    },
    { 
        rank: 3,
        totalPoints: 383,
        player1: { firstName: "Sebastian", lastName: "Duda", points: 188},
        player2: { firstName: "Roman", lastName: "Gutleber", points: 195},
    }
];

const TeamList = () => (
  <List>
    {exampleTeams.map(t => (
        <li key={t.rank}>
            <ListItem>
                <Rank>{t.rank}</Rank>
                <Players>
                    <span><Surname>{t.player1.lastName}</Surname> {t.player1.firstName}</span> <PlayerPoints>{t.player1.points}p</PlayerPoints><br />
                    <span><Surname>{t.player2.lastName}</Surname> {t.player2.firstName}</span> <PlayerPoints>{t.player2.points}p</PlayerPoints>
                </Players>
                <TotalPoints>{t.totalPoints} Points</TotalPoints>
            </ListItem>
        </li>
    ))}
  </List>
);

const PlayerPoints = styled.span`
    color: #6B787E;
    padding-left: 5px;
`

const Surname = styled.span`
    text-transform: uppercase;
`

const Rank = styled.span`
    font-size: 1.8em;
    padding: 0 15px 0 10px;
    color: ${Theme.c1};
`

const TotalPoints = styled.span`
    padding: 0 10px;
`

const List = styled.ul`
    background-color: #fff;
    color: #333;
    list-style: none;
    margin: 0;
    padding: 0;
`

const Players = styled.span`
    flex-grow: 1;
`

const ListItem = styled.div`
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 15px 0;
    border: 1px solid #cccccc;
    margin: 0 0 -1px 0;
    cursor: pointer;

    &:hover {
        border: 1px solid #898989;
    }
`

export default TeamList;
