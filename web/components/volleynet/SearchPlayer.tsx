import { createStyles, withStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import SearchIcon from '@material-ui/icons/Search';
import React from 'react';
import { connect } from 'react-redux';

import PlayerList from '../../components/volleynet/PlayerList';
import { searchVolleynetPlayersAction } from '../../redux/actions/entities';
import { searchVolleynetplayerSelector } from '../../redux/reducers/entities';

import { Gender, VolleynetSearchPlayer } from '../../types';
import LoadingButton from '../LoadingButton';

const styles = createStyles({
  container: {
    padding: '0 10px',
  },
});

interface Props {
  gender: Gender;
  onSelectPlayer: (VolleynetSearchPlayer) => void;
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
  searching: boolean;
}

class SearchPlayer extends React.Component<Props, State> {
  state = {
    firstName: '',
    lastName: '',
    birthday: '',
    searching: false,
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

  onSearch = async e => {
    const { firstName: fname, lastName: lname, birthday: bday } = this.state;
    const { searchVolleynetPlayers } = this.props;

    this.setState({ searching: true });

    e.preventDefault();

    try {
      await searchVolleynetPlayers({ fname, lname, bday });
    } finally {
      this.setState({ searching: false });
    }
  };

  render() {
    const { onSelectPlayer, foundPlayers, classes } = this.props;
    const { firstName, lastName, birthday, searching } = this.state;

    return (
      <form onSubmit={this.onSearch} className={classes.container}>
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
        <LoadingButton loading={searching}>
          <SearchIcon />
          Search
        </LoadingButton>
      </form>
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
