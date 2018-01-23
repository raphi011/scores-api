// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import Button from 'material-ui/Button';
import BackIcon from 'material-ui-icons/KeyboardArrowLeft';
import NextIcon from 'material-ui-icons/KeyboardArrowRight';
import MobileStepper from 'material-ui/MobileStepper';
import Router from 'next/router';
import withRedux from 'next-redux-wrapper';

import { validateMatch } from '../validation/match';
import Layout from '../components/Layout';
import SelectPlayers from '../components/SelectPlayers';
import SetScores from '../components/SetScores';
import withRoot from '../components/withRoot';
import initStore, { dispatchActions } from '../redux/store';
import { playersSelector, matchSelector } from '../redux/reducers/reducer';
import {
  createNewMatchAction,
  loadPlayersAction,
  userOrLoginRouteAction,
  loadMatchAction,
} from '../redux/actions/action';
import type { Match, Player } from '../types';

const styles = theme => ({
  root: {
    width: '90%',
  },
  submitButton: {
    marginRight: theme.spacing.unit,
    width: '100%',
  },
  stepContainer: {
    padding: '0 20px',
  },
  actionsContainer: {
    marginTop: theme.spacing.unit,
  },
  resetContainer: {
    marginTop: 0,
    padding: theme.spacing.unit * 3,
  },
  transition: {
    paddingBottom: 4,
  },
  button: {
    margin: theme.spacing.unit,
  },
});

type Props = {
  match: ?Match,
  classes: Object,
  playerIds: Array<number>,
  playersMap: { [number]: Player },
  createNewMatch: Match => Promise<any>,
};

type State = {
  activeStep: number,
  teamsComplete: boolean,
  match: {
    player1Id: number,
    player2Id: number,
    player3Id: number,
    player4Id: number,
    scoreTeam1: string,
    scoreTeam2: string,
    targetScore: string,
  },
  errors: {
    valid: boolean,
  },
};

class NewMatch extends React.Component<Props, State> {
  static async getInitialProps({ store, query, isServer, req, res }) {
    const actions = [loadPlayersAction(), userOrLoginRouteAction()];

    const { rematchId } = query;

    if (rematchId) {
      actions.push(loadMatchAction(Number.parseInt(rematchId, 10)));
    }

    await dispatchActions(store.dispatch, isServer, req, res, actions);

    return { rematchId };
  }

  constructor(props) {
    super(props);

    const { match } = props;

    const state = {
      activeStep: 0,
      teamsComplete: false,
      match: {
        player1Id: 0,
        player2Id: 0,
        player3Id: 0,
        player4Id: 0,
        scoreTeam1: '',
        scoreTeam2: '',
        targetScore: '15',
      },
      errors: {
        valid: true,
      },
    };

    if (match) {
      state.activeStep = 1;
      state.teamsComplete = true;
      state.match.player1Id = match.team1.player1Id;
      state.match.player2Id = match.team1.player2Id;
      state.match.player3Id = match.team2.player1Id;
      state.match.player4Id = match.team2.player2Id;
    }

    this.state = state;
  }

  onUnsetPlayer = selected => {
    const match = {
      ...this.state.match,
      [`player${selected}Id`]: 0,
    };

    this.setState({ match, teamsComplete: false });
  };

  onSetPlayer = (unassigned, id, teamsComplete) => {
    const activeStep = teamsComplete ? 1 : 0;
    const match = {
      ...this.state.match,
      [`player${unassigned}Id`]: id,
    };

    this.setState({ match, teamsComplete, activeStep });
  };

  onPrevious = () => {
    const { activeStep } = this.state;

    if (activeStep === 0) {
      Router.push('/');
      return;
    }

    this.setState({ activeStep: 0 });
  };

  onSetScores = () => {
    this.setState({ activeStep: 1 });
  };

  onChangeScore = (teamNr, score) => {
    const match = {
      ...this.state.match,
      [`scoreTeam${teamNr}`]: score,
    };

    this.setState({ match });
  };

  onChangeTargetScore = (e, targetScore) => {
    const match = {
      ...this.state.match,
      targetScore,
    };

    this.setState({ match });
  };

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
        Router.push('/');
      } catch (error) {}
    }
  };

  getMatch = (): Match => {
    const { scoreTeam1, scoreTeam2, targetScore, ...rest } = this.state.match;

    return {
      ...rest,
      scoreTeam1: Number.parseInt(scoreTeam1, 10) || 0,
      scoreTeam2: Number.parseInt(scoreTeam2, 10) || 0,
      targetScore: Number.parseInt(targetScore, 10),
    };
  };

  getPlayers = () => {
    const { playerIds, playersMap } = this.props;

    return playerIds.map(p => playersMap[p]);
  };

  render() {
    const { playersMap, classes } = this.props;
    const { teamsComplete, activeStep, match, errors } = this.state;

    const {
      scoreTeam1,
      scoreTeam2,
      player1Id,
      player2Id,
      player3Id,
      player4Id,
      targetScore,
    } = match;

    const players = this.getPlayers();

    const player1 = playersMap[player1Id];
    const player2 = playersMap[player2Id];
    const player3 = playersMap[player3Id];
    const player4 = playersMap[player4Id];

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
                {activeStep === 0 ? 'Cancel' : 'Back'}
              </Button>
            }
            nextButton={
              <Button
                onClick={this.onSetScores}
                dense
                disabled={activeStep === 1 || !teamsComplete}
              >
                Next
                <NextIcon className={classes.button} />
              </Button>
            }
          />
          <div>
            {activeStep === 0 ? (
              <SelectPlayers
                player1Id={player1Id}
                player2Id={player2Id}
                player3Id={player3Id}
                player4Id={player4Id}
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
  const { rematchId } = ownProps;
  const { playersMap, playerIds } = playersSelector(state);
  const match = rematchId ? matchSelector(state, rematchId) : null;

  return {
    playersMap,
    playerIds,
    match,
  };
}

const mapDispatchToProps = {
  loadPlayers: loadPlayersAction,
  createNewMatch: createNewMatchAction,
};

export default withStyles(styles)(
  withRedux(initStore, mapStateToProps, mapDispatchToProps)(withRoot(NewMatch)),
);
