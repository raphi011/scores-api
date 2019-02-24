// import React from 'react';

// import classNames from 'classnames';

// import {
//   createStyles,
//   Theme,
//   WithStyles,
//   withStyles,
// } from '@material-ui/core/styles';

// const statusMarkerStyles = (theme: Theme) =>
//   createStyles({
//     marker: {
//       minWidth: '7px',
//       width: '7px',
//     },
//     statusCanceled: { background: theme.palette.secondary[700] },
//     statusClosed: { background: theme.palette.secondary[800] },
//     statusOpen: { background: theme.palette.secondary[900] },
//   });

// interface StatusMarkerProps extends WithStyles<typeof statusMarkerStyles> {
//   status: string;
// }

// export default withStyles(statusMarkerStyles)(
//   ({ classes, status }: StatusMarkerProps) => {
//     const className = classNames(classes.marker, {
//       [classes.statusCanceled]: status === 'canceled',
//       [classes.statusClosed]: status === 'done',
//       [classes.statusOpen]: status === 'upcoming',
//     });

//     return <div className={className} />;
//   },
// );
