import React from 'react';
import { connect } from 'react-redux';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import { withStyles, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import SearchIcon from '@material-ui/icons/Search';

import PlayerList from '../../components/volleynet/PlayerList';
import { searchVolleynetplayerSelector } from '../../redux/reducers/entities';
import { searchVolleynetPlayersAction } from '../../redux/actions/entities';

import { Gender, VolleynetSearchPlayer } from '../../types';

const styles = createStyles({
  container: {
    padding: '0 10px',
  },
});

interface Props {
  gender: Gender;
  onSelectPlayer: VolleynetSearchPlayer;
  searchVolleynetPlayers: (
    params: { fname: string; lname: string; bday: string },
  ) => void;
  foundPlayers: VolleynetSearchPlayer[];
  classes: any;
}

interface State {
  firstName: string;
  lastName: string;
  birthday: string;
}

class SearchPlayer extends React.Component<Props, State> {
  state = {
    firstName: '',
    lastName: '',
    birthday: '',
  };

  onChangeFirstname = event => {
    const firstName = event.target.value;

    this.setState({ firstName });
  };

  onChangeLastname = event => {
    const lastName = event.target.value;

    this.setState({ lastName });
  };
  onChangeBirthday = event => {
    const birthday = event.target.value;

    this.setState({ birthday });
  };

  onSearch = async () => {
    const { firstName: fname, lastName: lname, birthday: bday } = this.state;
    const { searchVolleynetPlayers } = this.props;

    searchVolleynetPlayers({ fname, lname, bday });
  };

  render() {
    const { onSelectPlayer, foundPlayers, classes } = this.props;
    const { firstName, lastName, birthday } = this.state;

    return (
      <div className={classes.container}>
        <Typography variant="title" style={{ margin: '20px 0' }}>
          Search for your partner
        </Typography>
        <TextField
          label="Firstname"
          type="search"
          margin="normal"
          fullWidth
          onChange={this.onChangeFirstname}
          value={firstName}
        />
        <TextField
          label="Lastname"
          type="search"
          margin="normal"
          fullWidth
          onChange={this.onChangeLastname}
          value={lastName}
        />
        <TextField
          label="Birthday"
          type="search"
          margin="normal"
          fullWidth
          onChange={this.onChangeBirthday}
          value={birthday}
        />
        <PlayerList players={foundPlayers} onPlayerClick={onSelectPlayer} />
        <Button
          color="primary"
          type="submit"
          variant="raised"
          onClick={this.onSearch}
          fullWidth
        >
          <SearchIcon />
          Search
        </Button>
      </div>
    );
  }
}

const mapDispatchToProps = {
  searchVolleynetPlayers: searchVolleynetPlayersAction,
};

function mapStateToProps(state) {
  const foundPlayers = searchVolleynetplayerSelector(state);

  return { foundPlayers };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(SearchPlayer));
