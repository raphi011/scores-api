import ListItem from '@material-ui/core/ListItem';
import Typography from '@material-ui/core/Typography';
import { formatDate } from '../utils/dateFormat';

interface Props {
  appendix?: string;
  date: Date;
}

export default ({ date, appendix }: Props) => (
  <ListItem style={{ margin: '15px 0 3px 0' }}>
    <Typography variant="h6">
      {formatDate(date)} {appendix}
    </Typography>
  </ListItem>
);
