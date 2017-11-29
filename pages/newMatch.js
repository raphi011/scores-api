import React from "react";
import { withStyles } from "material-ui/styles";
import fetch from "isomorphic-unfetch";
import PersonIcon from "material-ui-icons/Person";
import Button from "material-ui/Button";
import BackIcon from "material-ui-icons/KeyboardArrowLeft";
import NextIcon from "material-ui-icons/KeyboardArrowRight";
import DoneIcon from "material-ui-icons/Done";
import MobileStepper from "material-ui/MobileStepper";
import Router from "next/router";
import withRedux from "next-redux-wrapper";

import Layout from "../components/Layout";
import CreateMatch from "../components/CreateMatch";
import SetScores from "../components/SetScores";
import withRoot from "../components/withRoot";
import initStore from "../redux/store";
import { playersSelector, statusSelector } from "../redux/reducers/reducer";
import {
  createNewMatchAction,
  loadPlayersAction
} from "../redux/actions/action";

const styles = theme => ({
  root: {
    width: "90%"
  },
  button: {
    marginRight: theme.spacing.unit
  },
  stepContainer: {
    padding: "0 20px"
  },
  actionsContainer: {
    marginTop: theme.spacing.unit
  },
  resetContainer: {
    marginTop: 0,
    padding: theme.spacing.unit * 3
  },
  transition: {
    paddingBottom: 4
  },
  button: {
    margin: theme.spacing.unit
  }
});

class NewMatch extends React.Component {
  state = {
    activeStep: 0,
    teamsComplete: false,
    player1ID: 0,
    player2ID: 0,
    player3ID: 0,
    player4ID: 0,
    scoreTeam1: "",
    scoreTeam2: ""
  };

  static async getInitialProps({ store }) {
    await store.dispatch(loadPlayersAction());
  }

  onUnsetPlayer = selected => {
    this.setState({ [`player${selected}ID`]: 0, teamsComplete: false });
  };

  onSetPlayer = (unassigned, ID, teamsComplete) => {
    const activeStep = teamsComplete ? 1 : 0;
    this.setState({ [`player${unassigned}ID`]: ID, teamsComplete, activeStep });
  };

  onSelectTeam = () => {
    this.setState({ activeStep: 0 });
  };

  onSetScores = () => {
    this.setState({ activeStep: 1 });
  };

  getPlayers = () => {
    const { playerIDs = [], playersMap } = this.props;

    return playerIDs.map(p => playersMap[p]);
  };

  onChangeScore = (teamNr, score) => {
    this.setState({ ["scoreTeam" + teamNr]: score });
  };

  onCreateMatch = async () => {
    const { formState, ...match } = this.state;
    const { createNewMatch } = this.props;

    await createNewMatch(match);
    Router.push("/");
  };

  render() {
    const { playersMap, classes, error } = this.props;
    const {
      teamsComplete,
      activeStep,
      scoreTeam1,
      scoreTeam2,
      ...selectedIDs
    } = this.state;

    const players = this.getPlayers();

    let body;

    const player1 = playersMap[selectedIDs.player1ID];
    const player2 = playersMap[selectedIDs.player2ID];
    const player3 = playersMap[selectedIDs.player3ID];
    const player4 = playersMap[selectedIDs.player4ID];

    return (
      <Layout title="New Match">
        <div>
          <MobileStepper
            position="static"
            className={classes.mobileStepper}
            steps={2}
            activeStep={activeStep}
            orientation="vertical"
            backButton={
              <Button
                dense
                disabled={activeStep == 0}
                onClick={this.onSelectTeam}
              >
                <BackIcon className={classes.button} />
                Back
              </Button>
            }
            nextButton={
              <Button
                onClick={this.onSetScores}
                dense
                disabled={activeStep == 1 || !teamsComplete}
              >
                Next
                <NextIcon className={classes.button} />
              </Button>
            }
          />
          <div>
            {activeStep == 0 ? (
              <CreateMatch
                {...selectedIDs}
                players={players}
                onSetPlayer={this.onSetPlayer}
                onUnsetPlayer={this.onUnsetPlayer}
              />
            ) : (
              <div className={classes.stepContainer}>
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
                  onClick={this.onCreateMatch}
                  raised
                  color="primary"
                >
                  <DoneIcon className={classes.leftIcon} />
                  Submit
                </Button>
              </div>
            )}
          </div>
        </div>
      </Layout>
    );
  }
}

function mapStateToProps(state) {
  const { playersMap, playerIDs } = playersSelector(state);

  return {
    playersMap,
    playerIDs
  };
}

const mapDispatchToProps = {
  loadPlayers: loadPlayersAction,
  createNewMatch: createNewMatchAction
};

export default withRedux(initStore, mapStateToProps, mapDispatchToProps)(
  withRoot(withStyles(styles)(NewMatch))
);
