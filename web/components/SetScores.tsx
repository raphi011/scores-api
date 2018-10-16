import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
import React, { SyntheticEvent } from 'react';

import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormLabel from '@material-ui/core/FormLabel';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import TextField from '@material-ui/core/TextField';
import DoneIcon from '@material-ui/icons/Done';

import { Player } from '../types';
import { IMatchValidation } from '../validation/match';

const styles = (theme: Theme) =>
  createStyles({
    button: {
      bottom: '24px',
      margin: theme.spacing.unit,
      position: 'fixed',
      right: '24px',
    },
    container: {
      display: 'flex',
      flexDirection: 'column',
    },
    formControl: {
      marginBottom: '20px',
    },
    group: {
      display: 'flex',
      flexDirection: 'row',
    },
    root: {
      display: 'flex',
    },
  });

interface IProps {
  onChangeTargetScore: (event: Event, score: number) => void;
  onChangeScore: (teamNr: number, event: string) => void;
  onLoseFocus: (
    teamNr: number,
    event: SyntheticEvent<HTMLButtonElement>,
  ) => void;
  player1: Player;
  player2: Player;
  player3: Player;
  player4: Player;
  scoreTeam1: string;
  scoreTeam2: string;
  targetScore: string;
  onCreateMatch: (Event) => void;
  errors: IMatchValidation;
  classes: any;
}

class SetScores extends React.PureComponent<IProps> {
  onChangeScoreTeam1 = e => {
    const { onChangeScore } = this.props;
    onChangeScore(1, e.target.value);
  };

  onChangeScoreTeam2 = e => {
    const { onChangeScore } = this.props;
    onChangeScore(2, e.target.value);
  };

  onLoseFocusScoreTeam1 = e => {
    const { onLoseFocus } = this.props;
    onLoseFocus(1, e);
  };

  onLoseFocusScoreTeam2 = e => {
    const { onLoseFocus } = this.props;
    onLoseFocus(2, e);
  };

  render() {
    const {
      player1,
      player2,
      player3,
      player4,
      scoreTeam1,
      scoreTeam2,
      targetScore,
      onChangeTargetScore,
      onCreateMatch,
      errors,
      classes,
    } = this.props;

    const score1Error = errors.scoreTeam1;
    const score2Error = errors.scoreTeam2;
    const allError = errors.all;

    return (
      <form onSubmit={onCreateMatch} className={classes.container}>
        <FormControl component="fieldset" className={classes.formControl}>
          <FormLabel component="legend">Target Score</FormLabel>
          <RadioGroup
            aria-label="Score"
            name="targetScore"
            className={classes.group}
            value={targetScore}
            onChange={onChangeTargetScore}
          >
            <FormControlLabel
              value="15"
              control={<Radio />}
              label="15 Points"
            />
            <FormControlLabel
              value="21"
              control={<Radio />}
              label="21 Points"
            />
          </RadioGroup>
        </FormControl>
        <FormControl className={classes.formControl}>
          <FormLabel>
            {player1.name} / {player2.name}
          </FormLabel>
          <TextField
            id="scoreTeam1"
            label={score1Error}
            error={!!score1Error}
            value={scoreTeam1}
            onChange={this.onChangeScoreTeam1}
            onBlur={this.onLoseFocusScoreTeam1}
            type="number"
            className={classes.textField}
            InputLabelProps={{
              shrink: true,
            }}
            margin="normal"
          />
        </FormControl>
        <FormControl className={classes.formControl}>
          <FormLabel>
            {player3.name} / {player4.name}
          </FormLabel>
          <TextField
            id="scoreTeam1"
            label={score2Error}
            error={!!score2Error}
            value={scoreTeam2}
            onChange={this.onChangeScoreTeam2}
            onBlur={this.onLoseFocusScoreTeam2}
            type="number"
            className={classes.textField}
            InputLabelProps={{
              shrink: true,
            }}
            margin="normal"
          />
          {allError}
        </FormControl>
        <Button
          className={classes.submitButton}
          color="primary"
          type="submit"
          variant="contained"
        >
          <DoneIcon className={classes.leftIcon} />
          Submit
        </Button>
      </form>
    );
  }
}

export default withStyles(styles)(SetScores);
