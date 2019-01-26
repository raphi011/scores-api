import React from 'react';

import Checkbox from '@material-ui/core/Checkbox';
import FormControl from '@material-ui/core/FormControl';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Input from '@material-ui/core/Input';
import InputAdornment from '@material-ui/core/InputAdornment';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import {
  createStyles,
  Theme,
  withStyles,
  WithStyles,
} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import SearchIcon from '@material-ui/icons/Search';
import { Gender } from '../../../types';
import LoadingButton from '../../LoadingButton';

const availableLeagues = [
  { name: 'Junior Tour', key: 'junior-tour' },
  { name: 'Amateur Tour', key: 'amateur-tour' },
  { name: 'Pro Tour', key: 'pro-tour' },
];
const availableGenders: Array<{name: string, key: Gender }> = [{ name: 'Female', key: 'W' }, { name: 'Male', key: 'M' }];
const availableSeasons = [2018, 2019];

const styles = (theme: Theme) =>
  createStyles({
    checkbox: {
      color: theme.palette.grey[700],
      padding: '4px',
    },
    checkboxes: {
      marginLeft: '9px',
    },
    filterGroup: {
      display: 'flex',
      flexDirection: 'column',
      marginBottom: '15px',
    },
    filterHeader: {
      color: theme.palette.grey[400],
      fontSize: '18px',
      marginBottom: '3px',
    },
    form: {
      maxWidth: '200px',
    },
  });

export interface Filters {
  league: string[];
  gender: Gender[];
  season: number;
}

interface Props extends WithStyles<typeof styles> {
  // availableLeagues: League[];
  // availableGenders: Gender[];
  // availableSeasons: number[];

  league: string[];
  gender: Gender[];
  season: number;
  
  loading: boolean;

  onFilter: (filters: Filters) => void;
}

type State = Filters;

class TournamentFilters extends React.Component<Props, State> {
  constructor(props) {
    super(props);

    const { league, gender, season } = this.props;

    this.state = {
      gender,
      league,
      season,
    };
  }

  onSelectSeason = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const { gender, season, league } = this.state;

    const selectedSeason = Number(event.target.value);

    if (selectedSeason === season) {
      return;
    }

    this.setState({
      gender,
      league,
      season: selectedSeason,
    });
  }

  onSelectLeague = (selected: string) => {
    const { gender, season, league } = this.state;

    if (league.length === 1 && league[0] === selected) {
      return;
    }

    let newSelected = league;

    if (league.includes(selected)) {
      newSelected = newSelected.filter(l => l !== selected);
    } else {
      newSelected.push(selected);
    }

    this.setState({
      gender,
      league: newSelected,
      season,
    });
  }

  onSelectGenders = (selected: Gender) => {
    const { gender, season, league } = this.state;

    if (gender.length === 1 && gender[0] === selected) {
      return;
    }

    let newSelected = gender;

    if (gender.includes(selected)) {
      newSelected = newSelected.filter(g => g !== selected);
    } else {
      newSelected.push(selected);
    }

    this.setState({
      gender: newSelected,
      league,
      season,
    });
  }

  onSubmit = (e) => {
    e.preventDefault();

    const { onFilter } = this.props;

    onFilter(this.state); 
  }

  render() {
    const { classes, loading = false } = this.props;
    const { gender, league, season } = this.state;

    return (
      <form onSubmit={this.onSubmit} autoComplete="off" className={classes.form}>
        <div className={classes.filterGroup}>
          <Typography className={classes.filterHeader}>Name</Typography>
          <FormControl>
            <InputLabel htmlFor="input-with-icon-adornment">Name</InputLabel>
            <Input
              id="input-with-icon-adornment"
              endAdornment={
                <InputAdornment position="end">
                  <SearchIcon />
                </InputAdornment>
              }
            />
          </FormControl>
        </div>

        <div className={classes.filterGroup}>
          <Typography className={classes.filterHeader}>Season</Typography>
          <Select value={season} onChange={this.onSelectSeason}>
            {availableSeasons.map(s => (
              <MenuItem key={s} value={s}>{s}</MenuItem>
            ))}
          </Select>
        </div>
        <div className={classes.filterGroup}>
          <Typography className={classes.filterHeader}>Gender</Typography>
          {availableGenders.map(g => (
            <FormControlLabel
              key={g.key}
              control={
                <Checkbox
                  checked={gender.includes(g.key)}
                  onChange={() => this.onSelectGenders(g.key)}
                  className={`${classes.checkbox} ${classes.checkboxes}`}
                  value={g.key}
                />
              }
              label={g.name}
            />
          ))}
        </div>
        <div className={classes.filterGroup}>
          <Typography className={classes.filterHeader}>Tour</Typography>
          {availableLeagues.map(l => (
            <FormControlLabel
              key={l.key}
              control={
                <Checkbox
                  checked={league.includes(l.key)}
                  onChange={() => this.onSelectLeague(l.key)}
                  className={`${classes.checkbox} ${classes.checkboxes}`}
                  value={l.key}
                />
              }
              label={l.name}
            />
          ))}
        </div>
        <LoadingButton loading={loading}>
          Search
        </LoadingButton>
      </form>
    );
  }
}

export default withStyles(styles)(TournamentFilters);
