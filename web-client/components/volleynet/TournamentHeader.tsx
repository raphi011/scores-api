import React from 'react';

import TimeAgo from 'react-timeago';
import classNames from 'classnames';

import Button from '@material-ui/core/Button';
import External from '@material-ui/icons/OpenInNew';
import Typography from '@material-ui/core/Typography';
import SignupIcon from '@material-ui/icons/AssignmentTurnedIn';
import Fab from '@material-ui/core/Fab';
import {
  createStyles,
  withStyles,
  WithStyles,
  Theme,
} from '@material-ui/core/styles';

import { formatDate } from '../../utils/date';
import { fontPalette } from '../../styles/theme';
import { link } from '../../styles/shared';
import { Tournament } from '../../types';
import { withWidth } from '@material-ui/core';
import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';

const styles = (theme: Theme) =>
  createStyles({
    attrContainer: {
      marginRight: '40px',
    },
    attr: {
      fontSize: fontPalette[800],
      fontWeight: 500,
      marginRight: '10px',
    },
    attrValue: {
      fontSize: fontPalette[500],
      fontWeight: 300,
      color: theme.palette.grey[500],
    },
    externalIcon: {
      fontSize: '16px',
      marginLeft: '10px',
      verticalAlign: 'middle',
    },
    fab: {
      position: 'fixed',
      bottom: theme.spacing.unit * 2,
      right: theme.spacing.unit * 2,
    },
    signupButton: {
      width: '120px',
    },
    titleRow: {
      marginBottom: '20px',
      display: 'flex',
    },
    titleRowDesktop: {
      alignItems: 'flex-start',
      flexDirection: 'row',
      justifyContent: 'space-between',
    },
    titleRowMobile: {
      flexDirection: 'column',
      alignItems: 'stretch',
    },
    stats: {
      display: 'flex',
      flexDirection: 'row',
      [theme.breakpoints.down('sm')]: {
        flexDirection: 'column',
      },
    },
    volleynetLink: {
      ...link,
    },
  });

interface Props extends WithStyles<typeof styles> {
  tournament: Tournament;
  width: Breakpoint;

  onSignup?: (event: React.MouseEvent<HTMLElement>) => void;
}

function TournamentHeader({ tournament, width, classes, onSignup }: Props) {
  const isMobile = ['xs', 'sm'].includes(width);

  const titleRowClassName = classNames(classes.titleRow, {
    [classes.titleRowDesktop]: !isMobile,
    [classes.titleRowMobile]: isMobile,
  });

  let button = null;
  let fabButton = null;

  if (onSignup) {
    if (isMobile) {
      fabButton = (
        <>
          <div style={{ height: '30px' }} />
          <Fab className={classes.fab} color="primary">
            <SignupIcon />
          </Fab>
        </>
      );
    } else {
      button = (
        <Button
          variant="contained"
          className={classes.signupButton}
          color="primary"
          onClick={onSignup}
        >
          Signup
        </Button>
      );
    }
  }

  return (
    <>
      <div className={titleRowClassName}>
        <div>
          <a
            href={tournament.link}
            className={classes.volleynetLink}
            target="_blank"
            rel="noopener noreferrer"
          >
            <Typography variant="h1" inline>
              {tournament.name}
              <External className={classes.externalIcon} />
            </Typography>
          </a>
          <Typography variant="subtitle1">
            {tournament.subLeague} - {formatDate(tournament.start)}
          </Typography>
        </div>
        {button}
      </div>
      <div>
        <div className={classes.stats}>
          <span className={classes.attrContainer}>
            <TimeAgo
              date={tournament.start}
              formatter={(value, unit, suffix) => (
                <>
                  <Typography inline className={classes.attr}>
                    {value}
                  </Typography>
                  <Typography inline className={classes.attrValue}>
                    {unit}
                    {value > 1 ? 's' : ''} {suffix}
                  </Typography>
                </>
              )}
            />
          </span>
          <span className={classes.attrContainer}>
            <Typography inline className={classes.attr}>
              {`${tournament.signedupTeams}/${tournament.maxTeams}`}
            </Typography>
            <Typography inline className={classes.attrValue}>
              Teams signed up
            </Typography>
          </span>
          <span className={classes.attrContainer}>
            <Typography inline className={classes.attr}>
              {tournament.maxPoints}
            </Typography>
            <Typography inline className={classes.attrValue}>
              Max points
            </Typography>
          </span>
        </div>
      </div>
      {fabButton}
    </>
  );
}

export default withStyles(styles)(withWidth()(TournamentHeader));
