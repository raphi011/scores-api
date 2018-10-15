import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import React from 'react';

const defaultLeagues = [
  { name: 'Junior Tour', key: 'JUNIOR TOUR' },
  { name: 'Amateur Tour', key: 'AMATEUR TOUR' },
  { name: 'Pro Tour', key: 'PRO TOUR' },
];

interface ILeague {
  name: string;
  key: string;
}

interface IProps {
  selected: string;
  onChange: (league: string) => void;
  leagues?: ILeague[];
}

export default ({ selected, onChange, leagues = defaultLeagues }: IProps) => (
  <Select value={selected} onChange={onChange}>
    {leagues.map(l => (
      <MenuItem key={l.key} value={l.key}>
        {l.name}
      </MenuItem>
    ))}
  </Select>
);
