import React from 'react';

import Checkbox from '@material-ui/core/Checkbox';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import {
  createStyles,
  Theme,
  withStyles,
  WithStyles,
} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Calendar from '@material-ui/icons/CalendarTodayTwoTone';

import { Gender } from '../../../types';
import LoadingButton from '../../LoadingButton';

const availableLeagues = [
  { name: 'Junior Tour', key: 'junior-tour' },
  { name: 'Amateur Tour', key: 'amateur-tour' },
  { name: 'Pro Tour', key: 'pro-tour' },
];
const availableGenders: { name: string; key: Gender }[] = [
  { name: 'Female', key: 'W' },
  { name: 'Male', key: 'M' },
];
const availableSeasons = [2018, 2019];

const styles = (theme: Theme) =>
  createStyles({
    calendarIcon: {
      margin: '8px 8px 8px 0',
    },
    checkbox: {
      color: theme.palette.grey[200],
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
      marginBottom: '3px',
    },
    font: {
      color: theme.palette.grey[700],
    },
    form: {
      width: '100%',
    },
    seasonFilterRow: {
      alignItems: 'center',
      display: 'flex',
      flexDirection: 'row',
    },
  });

export interface Filters {
  league: string[];
  gender: Gender[];
  season: number;
}

interface Props extends WithStyles<typeof styles> {
  league: string[];
  gender: Gender[];
  season: number;

  loading: boolean;

  onFilter: (filters: Filters) => void;
}

type State = Filters;

class TournamentFilters extends React.Component<Props, State> {
  constructor(props: Props) {
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
  };

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
  };

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
  };

  onSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const { onFilter } = this.props;

    onFilter(this.state);
  };

  render() {
    const { classes, loading = false } = this.props;
    const { gender, league, season } = this.state;

    return (
      <form
        onSubmit={this.onSubmit}
        autoComplete="off"
        className={classes.form}
      >
        <div className={classes.filterGroup}>
          <Typography variant="subtitle2" className={classes.filterHeader}>
            Season
          </Typography>
          <div className={classes.seasonFilterRow}>
            <Calendar
              className={classes.calendarIcon}
              color="primary"
              fontSize="small"
            />
            <Select
              style={{ marginTop: '3px' }}
              value={season}
              onChange={this.onSelectSeason}
              fullWidth
            >
              {availableSeasons.map(s => (
                <MenuItem classes={{ root: classes.font }} key={s} value={s}>
                  <Typography className={classes.font}>{s}</Typography>
                </MenuItem>
              ))}
            </Select>
          </div>
        </div>
        <div className={classes.filterGroup}>
          <Typography variant="subtitle2" className={classes.filterHeader}>
            Gender
          </Typography>
          {availableGenders.map(g => (
            <FormControlLabel
              key={g.key}
              classes={{ label: classes.font }}
              control={
                <Checkbox
                  checked={gender.includes(g.key)}
                  onChange={() => this.onSelectGenders(g.key)}
                  className={`${classes.checkbox} ${classes.checkboxes}`}
                  value={g.key}
                  color="primary"
                />
              }
              label={g.name}
            />
          ))}
        </div>
        <div className={classes.filterGroup}>
          <Typography variant="subtitle2" className={classes.filterHeader}>
            Tour
          </Typography>
          {availableLeagues.map(l => (
            <FormControlLabel
              key={l.key}
              classes={{ label: classes.font }}
              control={
                <Checkbox
                  checked={league.includes(l.key)}
                  onChange={() => this.onSelectLeague(l.key)}
                  className={`${classes.checkbox} ${classes.checkboxes}`}
                  value={l.key}
                  color="primary"
                />
              }
              label={l.name}
            />
          ))}
        </div>
        <LoadingButton loading={loading}>
          <span>Search</span>
        </LoadingButton>
      </form>
    );
  }
}

export default withStyles(styles)(TournamentFilters);
