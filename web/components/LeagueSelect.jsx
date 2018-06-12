// @flow

import React from 'react';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';

const defaultLeagues = [
  { name: 'Junior Tour', key: 'JUNIOR TOUR' },
  { name: 'Amateur Tour', key: 'AMATEUR TOUR' },
  { name: 'Pro Tour', key: 'PRO TOUR' },
];

type League = {
  name: string,
  key: string,
};

type Props = {
  selected: string,
  onChange: string => void,
  leagues: Array<League>,
};

export default ({ selected, onChange, leagues = defaultLeagues }: Props) => (
  <Select value={selected} onChange={onChange}>
    {leagues.map(l => (
      <MenuItem key={l.key} value={l.key}>
        {l.name}
      </MenuItem>
    ))}
  </Select>
);
