import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import React from 'react';

const defaultLeagues = [
  { name: 'Junior Tour', key: 'JUNIOR TOUR' },
  { name: 'Amateur Tour', key: 'AMATEUR TOUR' },
  { name: 'Pro Tour', key: 'PRO TOUR' },
];

interface League {
  name: string;
  key: string;
}

interface Props {
  selected: string;
  onChange: (string) => void;
  leagues?: League[];
}

export default ({ selected, onChange, leagues = defaultLeagues }: Props) => (
  <Select value={selected} onChange={onChange}>
    {leagues.map(l => (
      <MenuItem key={l.key} value={l.key}>
        {l.name}
      </MenuItem>
    ))}
  </Select>
);
