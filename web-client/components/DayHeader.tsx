import ListItem from '@material-ui/core/ListItem';
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import { formatDate } from '../utils/dateFormat';

const styles = (theme: Theme) =>
  createStyles({
    header: {
      color: theme.palette.grey[500],
      fontSize: '16px',
    },
  });

interface Props extends WithStyles<typeof styles> {
  appendix?: string;
  date: Date;
}

const DayHeader = ({ date, appendix, classes }: Props) => (
  <ListItem disableGutters>
    <Typography className={classes.header} variant="h6">
      {formatDate(date)} {appendix}
    </Typography>
  </ListItem>
);

export default withStyles(styles)(DayHeader);
