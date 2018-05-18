// @flow

import React from 'react';
import fetch from 'isomorphic-unfetch';
import TextField from '@material-ui/core/TextField';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import SearchIcon from '@material-ui/icons/Search';
import { BACKEND_URL } from '../../utils/env';

import PlayerList from '../../components/volleynet/PlayerList';
import type { Gender, VolleynetSearchPlayer } from '../../types';

const styles = () => ({
  container: {
    padding: '0 10px',
  },
});

function buildUrl(endpoint: string, params: Params = {}) {
  let paramUrl = '';

  const paramList = Object.keys(params)
    .filter(key => params[key])
    .map(key => `${key}=${params[key]}`);

  paramUrl = paramList.length ? `?${paramList.join('&')}` : '';

  const url = `${BACKEND_URL}/api/${endpoint}${paramUrl}`;

  return encodeURI(url);
}

type Props = {
  gender: Gender,
  onSelectPlayer: VolleynetSearchPlayer,
};

type State = {
  firstName: string,
  lastName: string,
  birthday: string,
  foundPlayers: Array<VolleynetSearchPlayer>,
};

class SearchPlayer extends React.Component<Props, State> {
  state = {
    firstName: '',
    lastName: '',
    birthday: '',
    foundPlayers: [],
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

    const response = await fetch(
      buildUrl('volleynet/players/search', { fname, lname, bday }),
    );

    const foundPlayers = await response.json();

    this.setState({ foundPlayers });
  };

  render() {
    const { gender, onSelectPlayer, classes } = this.props;
    const { firstName, lastName, birthday, foundPlayers } = this.state;

    return (
      <div className={classes.container}>
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

export default withStyles(styles)(SearchPlayer);
