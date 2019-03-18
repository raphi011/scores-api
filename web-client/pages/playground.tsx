import React from 'react';

import {
  createStyles,
  Theme,
  withStyles,
  WithStyles,
} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import withAuth from '../hoc/next/withAuth';
import Layout from '../containers/LayoutContainer';
import Divider from '@material-ui/core/Divider';

const styles = (theme: Theme) =>
  createStyles({
    example: {
      '&> *': {
        marginBottom: '20px',
      },
    },
    contrastText: {
      backgroundColor: theme.palette.primary.contrastText,
    },
    dark: {
      backgroundColor: theme.palette.primary.dark,
    },
    light: {
      backgroundColor: theme.palette.primary.light,
    },
    main: {
      backgroundColor: theme.palette.primary.main,
    },
  });

type Props = WithStyles<typeof styles>;

class Administration extends React.Component<Props> {
  render() {
    const { classes } = this.props;

    return (
      <Layout title={{ text: 'Administration', href: '' }}>
        <div className={classes.example}>
          <Typography variant="h1">h1 Example</Typography>
          <Typography variant="h2">h2 Example</Typography>
          <Typography variant="h3">h3 Example</Typography>
          <Typography variant="h4">h4 Example</Typography>
          <Typography variant="subtitle1">subtitle1 Example</Typography>
          <Typography variant="subtitle2">subtitle2 Example</Typography>
          <Typography variant="body1">body1 Example</Typography>
          <Typography variant="body2">body2 Example</Typography>
        </div>
        <Divider />
        <div>
          <div className={classes.contrastText}>primary.contrastText</div>
          <div className={classes.dark}>primary.dark</div>
          <div className={classes.light}>primary.light</div>
          <div className={classes.main}>primary.main</div>
        </div>
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Administration));
