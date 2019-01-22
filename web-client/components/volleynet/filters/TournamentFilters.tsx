import React, { SyntheticEvent } from 'react';

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

const leagues = [
  { name: 'Junior Tour', key: 'junior-tour' },
  { name: 'Amateur Tour', key: 'amateur-tour' },
  { name: 'Pro Tour', key: 'pro-tour' },
];

const genders = [{ name: 'Female', key: 'W' }, { name: 'Male', key: 'M' }];

const season = [2019, 2018];

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

interface Props extends WithStyles<typeof styles> {
  selected: string[];
  leagues?: League[];

  onChange: (event: SyntheticEvent<{}>, league: string[]) => void;
  onSubmit: () => void;
}

const TournamentFilters = ({ onSubmit, classes }: Props) => (
  <form onSubmit={onSubmit} autoComplete="off" className={classes.form}>
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
      <Select value={2018}>
        {season.map(s => (
          <MenuItem value={s}>{s}</MenuItem>
        ))}
      </Select>
    </div>
    <div className={classes.filterGroup}>
      <Typography className={classes.filterHeader}>Gender</Typography>
      {genders.map(g => (
        <FormControlLabel
          control={
            <Checkbox
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
      {leagues.map(l => (
        <FormControlLabel
          control={
            <Checkbox
              className={`${classes.checkbox} ${classes.checkboxes}`}
              value={l.key}
            />
          }
          label={l.name}
        />
      ))}
    </div>
    <Button fullWidth variant="contained" color="primary">
    Search
    </Button>
  </form>
);

export default withStyles(styles)(TournamentFilters);
