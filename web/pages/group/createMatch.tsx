import React, { SyntheticEvent } from 'react';

import Button from '@material-ui/core/Button';
import MobileStepper from '@material-ui/core/MobileStepper';
import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
import BackIcon from '@material-ui/icons/KeyboardArrowLeft';
import NextIcon from '@material-ui/icons/KeyboardArrowRight';
import Router from 'next/router';

import SelectPlayers from '../../components/SelectPlayers';
import SetScores from '../../components/SetScores';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import {
  createNewMatchAction,
  loadGroupAction,
  loadMatchAction,
} from '../../redux/entities/actions';
import {
  groupPlayersSelector,
  matchSelector,
} from '../../redux/entities/selectors';
import { setStatusAction } from '../../redux/status/actions';
import { Match, NewMatch, Player } from '../../types';
import { IMatchValidation, validateMatch } from '../../validation/match';

const styles = (theme: Theme) =>
  createStyles({
    actionsContainer: {
      marginTop: theme.spacing.unit,
    },
    button: {
      margin: theme.spacing.unit,
    },
    resetContainer: {
      marginTop: 0,
      padding: theme.spacing.unit * 3,
    },
    root: {
      width: '90%',
    },
    stepContainer: {
      padding: '0 20px',
    },
    submitButton: {
      marginRight: theme.spacing.unit,
      width: '100%',
    },
    transition: {
      paddingBottom: 4,
    },
  });

function calcWinnerScore(loserScore: number, targetScore: number): number {
  const winnerScore =
    loserScore <= targetScore - 2 ? targetScore : loserScore + 2;

  return winnerScore;
}

interface IProps {
  rematch?: Match;
  players: Player[];
  createNewMatch: (NewMatch) => Promise<any>;
  setStatus: (status: string) => void;
  /* eslint-disable react/no-unused-prop-types */
  groupId: number;
  rematchId: number;
  classes: any;
}

interface IState {
  activeStep: number;
  teamsComplete: boolean;
  match: {
    groupId: number;
    player1?: Player;
    player2?: Player;
    player3?: Player;
    player4?: Player;
    scoreTeam1: string;
    scoreTeam2: string;
    targetScore: string;
  };
  errors: IMatchValidation;
}

class CreateMatch extends React.Component<IProps, IState> {
  static mapDispatchToProps = {
    createNewMatch: createNewMatchAction,
    setStatus: setStatusAction,
  };
  static getParameters(query) {
    let { rematchId, groupId } = query;

    rematchId = Number.parseInt(rematchId, 10) || 0;
    groupId = Number.parseInt(groupId, 10) || 0;

    return { groupId, rematchId };
  }

  static buildActions({ rematchId, groupId }) {
    const actions = [loadGroupAction(groupId)];

    if (rematchId) {
      actions.push(loadMatchAction(rematchId));
    }

    return actions;
  }

  static mapStateToProps(state, ownProps: IProps) {
    const { rematchId, groupId } = ownProps;

    const players = groupPlayersSelector(state, groupId);
    const rematch = rematchId ? matchSelector(state, rematchId) : null;

    return {
      players,
      rematch,
    };
  }

  state = {
    activeStep: 0,
    errors: {
      valid: true,
    },
    match: {
      groupId: 0,
      player1: null,
      player2: null,
      player3: null,
      player4: null,
      scoreTeam1: '',
      scoreTeam2: '',
      targetScore: '15',
    },
    teamsComplete: false,
  };

  rematchPlayersSet = false;

  constructor(props) {
    super(props);

    const state = this.setRematch(props);

    if (state) {
      this.state = {
        ...this.state,
        ...state,
      };
    } else {
      this.state.match.groupId = props.groupId;
    }
  }

  componentWillReceiveProps(nextProps) {
    const state = this.setRematch(nextProps);

    if (state) {
      this.setState(state);
    }
  }

  onUnsetPlayer = (selected: number) => {
    const match = {
      ...this.state.match,
    };

    switch (selected) {
      case 1:
        match.player1 = null;
        break;
      case 2:
        match.player2 = null;
        break;
      case 3:
        match.player3 = null;
        break;
      case 4:
        match.player4 = null;
        break;
      default:
        throw new Error(`Can't unset player: ${selected}`);
    }

    this.setState({ match, teamsComplete: false });
  };

  setRematch = (props: IProps) => {
    const { rematchId, rematch, players } = props;

    if (rematchId && !this.rematchPlayersSet) {
      if (!rematch || !players.length) {
        return null; // rematch or players not loaded yet
      }

      const newState = {
        activeStep: 1,
        match: {
          groupId: rematch.groupId,
          player1: rematch.team1.player1,
          player2: rematch.team1.player2,
          player3: rematch.team2.player1,
          player4: rematch.team2.player2,
          scoreTeam1: '',
          scoreTeam2: '',
          targetScore: '15',
        },
        teamsComplete: true,
      };

      this.rematchPlayersSet = true;

      return newState;
    }

    return null;
  };

  onSetPlayer = (playerNr: number, player: Player, teamsComplete: boolean) => {
    const activeStep = teamsComplete ? 1 : 0;
    const match = {
      ...this.state.match,
    };

    switch (playerNr) {
      case 1:
        match.player1 = player;
        break;
      case 2:
        match.player2 = player;
        break;
      case 3:
        match.player3 = player;
        break;
      case 4:
        match.player4 = player;
        break;
      default:
        throw new Error(`Can't set player: ${playerNr}`);
    }

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

  onChangeScore = (teamNr: number, score: string) => {
    const match = {
      ...this.state.match,
    };

    switch (teamNr) {
      case 1:
        match.scoreTeam1 = score;
        break;
      case 2:
        match.scoreTeam2 = score;
        break;
      default:
        throw new Error(`Can't set score for team: ${teamNr}`);
    }

    this.setState({ match });
  };

  onChangeTargetScore = (_, targetScore: string) => {
    const match = {
      ...this.state.match,
      targetScore,
    };

    this.setState({ match });
  };

  onScoreLoseFocus = (teamNr: number) => {
    const match = {
      ...this.state.match,
    };

    const scoreTeam1 = Number.parseInt(match.scoreTeam1, 10);
    const scoreTeam2 = Number.parseInt(match.scoreTeam2, 10);
    const targetScore = Number.parseInt(match.targetScore, 10);

    switch (teamNr) {
      case 1: {
        if (Number.isNaN(scoreTeam1) || Number.isInteger(scoreTeam2)) {
          return;
        }

        match.scoreTeam2 = calcWinnerScore(scoreTeam1, targetScore).toString();
        break;
      }
      case 2: {
        if (Number.isNaN(scoreTeam2) || Number.isInteger(scoreTeam1)) {
          return;
        }

        match.scoreTeam1 = calcWinnerScore(scoreTeam2, targetScore).toString();
        break;
      }
      default:
        throw new Error(`Can't set score for team: ${teamNr}`);
    }

    this.setState({ match });
  };

  onCreateMatch = async (e: SyntheticEvent<HTMLButtonElement>) => {
    e.preventDefault();

    const { createNewMatch, setStatus } = this.props;
    const match = this.getMatch();

    const errors = validateMatch(match);

    if (!errors.valid) {
      this.setState({ errors });
    } else {
      try {
        await createNewMatch(match);
        await Router.push('/');
        setStatus('New match created');
      } catch (error) {
        // ignore
      }
    }
  };

  getMatch = (): NewMatch => {
    const {
      groupId,
      scoreTeam1,
      scoreTeam2,
      targetScore,
      player1,
      player2,
      player3,
      player4,
    } = this.state.match;

    return {
      groupId,
      player1Id: player1 ? player1.id : 0,
      player2Id: player2 ? player2.id : 0,
      player3Id: player3 ? player3.id : 0,
      player4Id: player4 ? player4.id : 0,
      scoreTeam1: Number.parseInt(scoreTeam1, 10) || 0,
      scoreTeam2: Number.parseInt(scoreTeam2, 10) || 0,
      targetScore: Number.parseInt(targetScore, 10) || 0,
    };
  };

  render() {
    const { players, classes } = this.props;
    const { teamsComplete, activeStep, match, errors } = this.state;

    const {
      scoreTeam1,
      scoreTeam2,
      player1,
      player2,
      player3,
      player4,
      targetScore,
    } = match;

    return (
      <Layout title={{ text: 'New Match', href: '' }}>
        <div>
          <MobileStepper
            position="static"
            className={classes.mobileStepper}
            steps={2}
            activeStep={activeStep}
            backButton={
              <Button size="small" onClick={this.onPrevious}>
                <BackIcon className={classes.button} />
                {activeStep === 0 ? 'Cancel' : 'Back'}
              </Button>
            }
            nextButton={
              <Button
                onClick={this.onSetScores}
                size="small"
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
                player1={player1}
                player2={player2}
                player3={player3}
                player4={player4}
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
                  onLoseFocus={this.onScoreLoseFocus}
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

export default withAuth(withStyles(styles)(CreateMatch));
