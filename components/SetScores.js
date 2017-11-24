import React from "react";

class SetScores extends React.Component {
  render() {
    const {
      player1,
      player2,
      player3,
      player4,
      scoreTeam1,
      scoreTeam2,
      onChangeScore
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
        <input
          type="range"
          min={0}
          max={40}
          step={1}
          value={scoreTeam1}
          onChange={onChangeScoreTeam1}
        />
        <p>
          {player3.Name} / {player4.Name}: {scoreTeam2}
        </p>
        <input
          type="range"
          min={0}
          max={40}
          step={1}
          value={scoreTeam2}
          onChange={onChangeScoreTeam2}
        />
      </div>
    );
  }
}

export default SetScores;
