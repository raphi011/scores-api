import React from "react";
import TextField from "material-ui/TextField";
import Radio from 'material-ui/Radio';
import { withStyles } from "material-ui/styles";

const styles = theme => ({
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: 200
  }
});

class SetScores extends React.Component {
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
      onChangeScore,
      errors,
      classes
    } = this.props;

    const score1Error = errors.scoreTeam1;
    const score2Error = errors.scoreTeam2;
    const allError = errors.all;

    const onChangeScoreTeam1 = e =>
      onChangeScore(1, Number.parseInt(e.target.value));
    const onChangeScoreTeam2 = e =>
      onChangeScore(2, Number.parseInt(e.target.value));

    return (
      <div>
        <div>
          <Radio
            checked={targetScore === 15}
            onChange={onChangeTargetScore}
            label="15 point game"
            value="15"
            name="targetScore"
            aria-label={targetScore}
          />

          <Radio
            checked={targetScore === 21}
            onChange={onChangeTargetScore}
            label="21 point game"
            value="21"
            name="targetScore"
            aria-label={targetScore}
          />
        </div>
        <p>
          {player1.Name} / {player2.Name}
        </p>
        <TextField
          id="number"
          label={score1Error || "Score"}
          error={!!score1Error}
          value={scoreTeam1}
          onChange={onChangeScoreTeam1}
          type="number"
          className={classes.textField}
          min={0}
          InputLabelProps={{
            shrink: true
          }}
          margin="normal"
        />
        <p>
          {player3.Name} / {player4.Name}
        </p>
        <TextField
          id="number"
          label={score2Error || "Score"}
          error={!!score2Error}
          value={scoreTeam2}
          onChange={onChangeScoreTeam2}
          type="number"
          className={classes.textField}
          min={0}
          InputLabelProps={{
            shrink: true
          }}
          margin="normal"
        />
        {allError}
      </div>
    );
  }
}

export default withStyles(styles)(SetScores);
