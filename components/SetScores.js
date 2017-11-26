import React from "react";
import TextField from 'material-ui/TextField';
import { withStyles } from 'material-ui/styles';

const styles = theme => ({
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: 200,
  },
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
      onChangeScore,
      classes
    } = this.props;

    const onChangeScoreTeam1 = e =>
      onChangeScore(1, Number.parseInt(e.target.value));
    const onChangeScoreTeam2 = e =>
      onChangeScore(2, Number.parseInt(e.target.value));

    return (
      <div>
        <p>
          {player1.Name} / {player2.Name}: {scoreTeam1}
        </p>
        <TextField
          id="number"
          label="Score"
          value={scoreTeam1}
          onChange={onChangeScoreTeam1}
          type="number"
          className={classes.textField}
          min={0}
          InputLabelProps={{
            shrink: true,
          }}
          margin="normal"
        />
        <p>
          {player3.Name} / {player4.Name}: {scoreTeam2}
        </p>
        <TextField
          id="number"
          label="Score"
          value={scoreTeam2}
          onChange={onChangeScoreTeam2}
          type="number"
          className={classes.textField}
          min={0}
          InputLabelProps={{
            shrink: true,
          }}
          margin="normal"
        />
      </div>
    );
  }
}


export default withStyles(styles)(SetScores);
