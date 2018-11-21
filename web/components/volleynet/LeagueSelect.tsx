import ToggleButton from '@material-ui/lab/ToggleButton';
import ToggleButtonGroup from '@material-ui/lab/ToggleButtonGroup';
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
  selected: string[];
  onChange: (event, league: string[]) => void;
  leagues?: ILeague[];
}

export default ({ selected, onChange, leagues = defaultLeagues }: IProps) => (
  <ToggleButtonGroup value={selected} onChange={onChange}>
    {leagues.map(l => (
      <ToggleButton key={l.key} value={l.key}>
        {l.name}
      </ToggleButton>
    ))}
  </ToggleButtonGroup>
);
