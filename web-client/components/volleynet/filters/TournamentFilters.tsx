import React from 'react';

import Button from '@material-ui/core/Button';
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

const availableLeagues = [
  { name: 'Junior Tour', key: 'junior-tour' },
  { name: 'Amateur Tour', key: 'amateur-tour' },
  { name: 'Pro Tour', key: 'pro-tour' },
];
const availableGenders = [{ name: 'Female', key: 'W' }, { name: 'Male', key: 'M' }];
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

interface League {
  name: string;
  key: string;
}

interface Filters {
  leagues: League[];
  genders: Gender[];
  season: number;
}

interface Props extends WithStyles<typeof styles> {
  // leagues: League[];
  // genders: Gender[];
  // seasons: number[];

  leagues: League[];
  genders: Gender[];
  season: number;

  onChange: (filters: Filters) => void;
  onSubmit: () => void;
}

class TournamentFilters extends React.Component<Props> {
  onSelectSeason = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const { genders, season, leagues, onChange } = this.props;

    const selectedSeason = Number(event.target.value);

    if (selectedSeason === season) {
      return;
    }

    onChange({
      genders,
      leagues,
      season: selectedSeason,
    });
  }

  onSelectLeague = (league: League) => {
    const { genders, season, leagues, onChange } = this.props;

    if (leagues.length === 1 && leagues[0] === league) {
      return;
    }

    let newSelected = leagues;

    if (leagues.includes(league)) {
      newSelected = newSelected.filter(l => l !== league);
    } else {
      newSelected.push(league);
    }

    onChange({
      genders,
      leagues: newSelected,
      season,
    });
  }

  onSelectGenders = (gender: Gender) => {
    const { genders, season, leagues, onChange } = this.props;

    if (genders.length === 1 && genders[0] === gender) {
      return;
    }

    let newSelected = genders;

    if (genders.includes(gender)) {
      newSelected = newSelected.filter(g => g !== gender);
    } else {
      newSelected.push(gender);
    }

    onChange({
      genders: newSelected,
      leagues,
      season,
    });
  }

  onSubmit = (e) => {
    e.preventDefault();

    const { onSubmit } = this.props;

    onSubmit(); 
  }

  render() {
    const { classes, genders, leagues, season } = this.props;

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
              <MenuItem value={s}>{s}</MenuItem>
            ))}
          </Select>
        </div>
        <div className={classes.filterGroup}>
          <Typography className={classes.filterHeader}>Gender</Typography>
          {availableGenders.map(g => (
            <FormControlLabel
              control={
                <Checkbox
                  checked={genders.includes(g.key)}
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
              control={
                <Checkbox
                  checked={leagues.includes(l.key)}
                  onChange={() => this.onSelectLeague(l.key)}
                  className={`${classes.checkbox} ${classes.checkboxes}`}
                  value={l.key}
                />
              }
              label={l.name}
            />
          ))}
        </div>
        <Button fullWidth type="submit" variant="contained" color="primary">
          Search
        </Button>
      </form>
    );
  }
}

export default withStyles(styles)(TournamentFilters);
