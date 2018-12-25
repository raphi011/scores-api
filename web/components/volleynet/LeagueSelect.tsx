import React, { SyntheticEvent } from 'react';

import ToggleButton from '@material-ui/lab/ToggleButton';
import ToggleButtonGroup from '@material-ui/lab/ToggleButtonGroup';

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
  selected: string[];
  leagues?: League[];

  onChange: (event: SyntheticEvent<{}>, league: string[]) => void;
}

export default ({ selected, onChange, leagues = defaultLeagues }: Props) => (
  <ToggleButtonGroup value={selected} onChange={onChange}>
    {leagues.map(l => (
      <ToggleButton key={l.key} value={l.key}>
        {l.name}
      </ToggleButton>
    ))}
  </ToggleButtonGroup>
);
