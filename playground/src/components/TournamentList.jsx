import React from "react";
import styled from "styled-components";
import * as Theme from "../theme";

const exampleTournaments = [
    {
        id: 1,
        type: "Amateur 1",
        name: "Herren Eibesbrunn",
        when: new Date(),
        teams: "3 / 12",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 2,
        type: "Amateur 1",
        name: "Herren Wulzeshofen",
        when: new Date(),
        teams: "5 / 18",
        color: "#000",
    },
    { 
        id: 3,
        type: "Pro",
        name: "Herren Beachvolleyclub",
        when: new Date(),
        teams: "12 / 32",
        color: "#000",
    }
];

const TournamentList = () => (
  <List>
    {exampleTournaments.map(t => (
        <li key={t.id}>
            <ListItem>
                <Title>{t.name}</Title>
                <Subtitle>{t.type} • {t.teams} • {t.when.toDateString()}</Subtitle>
            </ListItem>
        </li>
    ))}
  </List>
);

const Title = styled.div`
    font-size: 20px;
`

const Subtitle = styled.div`
    color: #6B787E;
`

const List = styled.ul`
    background-color: #fff;
    color: #333;
    list-style: none;
    margin: 0;
    padding: 0;
`

const ListItem = styled.div`
    display: flex;
    flex-direction: column;
    padding: 15px 10px;
    border: 1px solid #cccccc;
    margin: 0 0 -1px 0;
    cursor: pointer;

    &:hover {
        border: 1px solid #898989;
    }
`

export default TournamentList;

