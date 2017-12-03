import React from "react";
import { withStyles } from "material-ui/styles";
import fetch from "isomorphic-unfetch";
import PersonIcon from "material-ui-icons/Person";
import Button from "material-ui/Button";
import BackIcon from "material-ui-icons/KeyboardArrowLeft";
import NextIcon from "material-ui-icons/KeyboardArrowRight";
import MobileStepper from "material-ui/MobileStepper";
import Router from "next/router";
import withRedux from "next-redux-wrapper";

import { validateMatch } from "../validation/match";
import Layout from "../components/Layout";
import SelectPlayers from "../components/SelectPlayers";
import SetScores from "../components/SetScores";
import withRoot from "../components/withRoot";
import initStore, { dispatchActions } from "../redux/store";
import { playersSelector, matchSelector, statusSelector } from "../redux/reducers/reducer";
import {
  createNewMatchAction,
  loadPlayersAction,
  userOrLoginRouteAction,
  loadMatchAction
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
  static async getInitialProps({ store, query, isServer, req, res }) {
    const actions = [loadPlayersAction(), userOrLoginRouteAction()];

    const { rematchID } = query;

    if (rematchID) {
      actions.push(loadMatchAction(Number.parseInt(rematchID)));
    }

    await dispatchActions(store.dispatch, isServer, req, res, actions);

    return { rematchID };
  }

  constructor(props) {
    super(props);

    const { match } = props;

    const state = {
      activeStep: 0,
      teamsComplete: false,
      match: {
        player1ID: 0,
        player2ID: 0,
        player3ID: 0,
        player4ID: 0,
        scoreTeam1: "",
        scoreTeam2: "",
        targetScore: "15"
      },
      errors: {
        valid: true
      }
    };

    if (match) {
      state.activeStep = 1;
      state.teamsComplete = true;
      state.match.player1ID = match.Team1.Player1ID;
      state.match.player2ID = match.Team1.Player2ID;
      state.match.player3ID = match.Team2.Player1ID;
      state.match.player4ID = match.Team2.Player2ID;
    }

    this.state = state;
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

  onPrevious = () => {
    let { activeStep } = this.state;

    if (activeStep === 0) {
      Router.push("/");
      return;
    }

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

  getMatch = () => {
    const {
      scoreTeam1,
      scoreTeam2,
      targetScore,
      ...rest,
    } = this.state.match;

    return {
      ...rest,
      scoreTeam1: Number.parseInt(scoreTeam1) || 0,
      scoreTeam2: Number.parseInt(scoreTeam2) || 0,
      targetScore: Number.parseInt(targetScore),
    };
  }

  onCreateMatch = async e => {
    e.preventDefault();

    const { createNewMatch } = this.props;
    const match = this.getMatch();

    const errors = validateMatch(match);

    if (!errors.valid) {
      this.setState({ errors });
    } else {
      try {
        await createNewMatch(match);
        Router.push("/");
      } catch (e) {}
    }
  };

  onChangeTargetScore = (e, targetScore) => {
    const match = {
      ...this.state.match,
      targetScore
    };

    this.setState({ match });
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
      player4ID,
      targetScore
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
              <Button dense onClick={this.onPrevious}>
                <BackIcon className={classes.button} />
                {activeStep === 0 ? "Cancel" : "Back"}
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
              <SelectPlayers
                player1ID={player1ID}
                player2ID={player2ID}
                player3ID={player3ID}
                player4ID={player4ID}
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
                  errors={errors}
                  scoreTeam1={scoreTeam1}
                  scoreTeam2={scoreTeam2}
                  targetScore={targetScore}
                  onChangeScore={this.onChangeScore}
                  onChangeTargetScore={this.onChangeTargetScore}
                  onCreateMatch={this.onCreateMatch}
                />
              </div>
            )}
          </div>
        </div>
      </Layout>
    );
  }
}

function mapStateToProps(state, ownProps) {
  const { rematchID } = ownProps;
  const { playersMap, playerIDs } = playersSelector(state);
  const match = rematchID ? matchSelector(state, rematchID) : null;

  return {
    playersMap,
    playerIDs,
    match,
  };
}

const mapDispatchToProps = {
  loadPlayers: loadPlayersAction,
  createNewMatch: createNewMatchAction
};

export default withRedux(initStore, mapStateToProps, mapDispatchToProps)(
  withRoot(withStyles(styles)(NewMatch))
);
