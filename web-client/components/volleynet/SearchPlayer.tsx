import React from 'react';

import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import SearchIcon from '@material-ui/icons/Search';
import { connect } from 'react-redux';

import { searchPlayersAction } from '../../redux/entities/actions';

import LoadingButton from '../LoadingButton';

const styles = createStyles({
  container: {
    padding: '0 10px',
  },
});

interface Props extends WithStyles<typeof styles> {
  gender: string;

  searchVolleynetPlayers: (params: {
    fname: string;
    lname: string;
    gender: string;
  }) => void;
}

interface State {
  firstName: string;
  lastName: string;
  searching: boolean;
}

class PlayerSearch extends React.Component<Props, State> {
  state = {
    firstName: '',
    lastName: '',
    searching: false,
  };

  onChangeFirstname = (event: React.ChangeEvent<HTMLInputElement>) => {
    const firstName = event.target.value;

    this.setState({ firstName });
  };

  onChangeLastname = (event: React.ChangeEvent<HTMLInputElement>) => {
    const lastName = event.target.value;

    this.setState({ lastName });
  };

  onSearch = async (e: React.FormEvent<HTMLFormElement>) => {
    const { firstName: fname, lastName: lname } = this.state;
    const { gender, searchVolleynetPlayers } = this.props;

    this.setState({ searching: true });

    e.preventDefault();

    try {
      await searchVolleynetPlayers({ fname, lname, gender });
    } finally {
      this.setState({ searching: false });
    }
  };

  render() {
    const { classes } = this.props;
    const { firstName, lastName, searching } = this.state;

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
        <LoadingButton loading={searching}>
          <SearchIcon />
          <span>Search</span>
        </LoadingButton>
      </form>
    );
  }
}

const mapDispatchToProps = {
  searchVolleynetPlayers: searchPlayersAction,
};

export default connect(
  null,
  mapDispatchToProps,
)(withStyles(styles)(PlayerSearch));
