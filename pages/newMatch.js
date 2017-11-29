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

import { validateMatch } from "../validation/match";
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
import withWidth from "material-ui/utils/withWidth";

const styles = theme => ({
  root: {
    width: "90%"
  },
  submitButton: {
    marginRight: theme.spacing.unit,
    width: "100%"
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
    match: {
      player1ID: 0,
      player2ID: 0,
      player3ID: 0,
      player4ID: 0,
      scoreTeam1: "",
      scoreTeam2: ""
    },
    errors: {
      valid: true
    }
  };

  static async getInitialProps({ store }) {
    await store.dispatch(loadPlayersAction());
  }

  onUnsetPlayer = selected => {
    const match = {
      ...this.state.match,
      [`player${selected}ID`]: 0
    };

    this.setState({ match, teamsComplete: false });
  };

  onSetPlayer = (unassigned, ID, teamsComplete) => {
    const activeStep = teamsComplete ? 1 : 0;
    const match = {
      ...this.state.match,
      [`player${unassigned}ID`]: ID
    };

    this.setState({ match, teamsComplete, activeStep });
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
    const match = {
      ...this.state.match,
      ["scoreTeam" + teamNr]: score
    };

    this.setState({ match });
  };

  onCreateMatch = async () => {
    const { match } = this.state;
    const { createNewMatch } = this.props;

    const errors = validateMatch(match);

    if (!errors.valid) {
      this.setState({ errors });
    } else {
      await createNewMatch(match);
      Router.push("/");
    }
  };

  render() {
    const { playersMap, classes, error } = this.props;
    const { teamsComplete, activeStep, match, errors } = this.state;

    const {
      scoreTeam1,
      scoreTeam2,
      player1ID,
      player2ID,
      player3ID,
      player4ID
    } = match;

    const players = this.getPlayers();

    let body;

    const player1 = playersMap[player1ID];
    const player2 = playersMap[player2ID];
    const player3 = playersMap[player3ID];
    const player4 = playersMap[player4ID];

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
                player1ID={player1ID}
                player2ID={player2ID}
                player3ID={player3ID}
                player4ID={player4ID}
                players={players}
                onSetPlayer={this.onSetPlayer}
                onUnsetPlayer={this.onUnsetPlayer}
              />
            ) : (
              <div>
                <div className={classes.stepContainer}>
                  <SetScores
                    player1={player1}
                    player2={player2}
                    player3={player3}
                    player4={player4}
                    errors={errors}
                    scoreTeam1={scoreTeam1}
                    scoreTeam2={scoreTeam2}
                    onChangeScore={this.onChangeScore}
                  />
                </div>
                <Button
                  className={classes.submitButton}
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
