import React from "react";
import { withStyles } from "material-ui/styles";
import fetch from "isomorphic-unfetch";
import PersonIcon from "material-ui-icons/Person";
import Button from "material-ui/Button";
import BackIcon from "material-ui-icons/KeyboardArrowLeft";
import DoneIcon from "material-ui-icons/Done";

import Layout from "../components/Layout";
import CreateMatch from "../components/CreateMatch";
import SetScores from "../components/SetScores";

const styles = theme => ({
  button: {
    margin: theme.spacing.unit
  },
  leftIcon: {
    marginRight: theme.spacing.unit
  }
});

class NewMatch extends React.Component {
  state = {
    formState: 1,
    player1: 0,
    player2: 0,
    player3: 0,
    player4: 0,
    scoreTeam1: 21,
    scoreTeam2: 21,
  };

  static async getInitialProps() {
    try {
      const playersResponse = await fetch("http://localhost:3000/api/players");
      const players = await playersResponse.json();

      const playerMap = {};
      const playerIDs = [];
      players.forEach(p => {
        playerIDs.push(p.ID);
        playerMap[p.ID] = p;
      });

      return { playerMap, playerIDs };
    } catch (error) {
      return { error };
    }
  }

  onUnsetPlayer = selected => {
    this.setState({ ["player" + selected]: 0 });
  };

  onSetPlayer = (unassigned, ID, teamsComplete) => {
    console.log(teamsComplete);
    const formState = teamsComplete ? 2 : 1;

    this.setState({ ["player" + unassigned]: ID, formState });
  };

  onSelectTeam = () => {
    this.setState({ formState: 1 });
  };

  getPlayers = () => {
    const { playerIDs = [], playerMap } = this.props;

    return playerIDs.map(p => playerMap[p]);
  };

  onChangeScore = (teamNr, score) => {
    this.setState({ ['scoreTeam'+teamNr]: score });
  }

  render() {
    const { error, playerMap, classes } = this.props;
    const { formState, scoreTeam1, scoreTeam2, ...selectedIDs } = this.state;

    const players = this.getPlayers();

    let body;

    if (formState === 1) {
      body = (
        <CreateMatch
          {...selectedIDs}
          players={players}
          onSetPlayer={this.onSetPlayer}
          onUnsetPlayer={this.onUnsetPlayer}
        />
      );
    } else {
      const player1 = playerMap[selectedIDs.player1];
      const player2 = playerMap[selectedIDs.player2];
      const player3 = playerMap[selectedIDs.player3];
      const player4 = playerMap[selectedIDs.player4];

      body = (
        <div>
          <Button
            className={classes.button}
            onClick={this.onSelectTeam}
            raised
            color="accent"
          >
            <BackIcon className={classes.leftIcon} />
            Back
          </Button>
          <SetScores
            player1={player1}
            player2={player2}
            player3={player3}
            player4={player4}
            scoreTeam1={scoreTeam1}
            scoreTeam2={scoreTeam2}
            onChangeScore={this.onChangeScore}
          />
          <Button
            className={classes.button}
            raised
            color="primary"
          >
            <DoneIcon className={classes.leftIcon} />
            Submit
          </Button>

        </div>
      );
    }

    return (
      <Layout>
        <h1>New Match</h1>
        {body}
      </Layout>
    );
  }
}

export default withStyles(styles)(NewMatch);
