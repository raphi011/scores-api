import Chip from '@material-ui/core/Chip';
import ListItem from '@material-ui/core/ListItem';
import { formatDate } from '../utils/dateFormat';

interface Props {
  date: Date;
}

export default ({ date }: Props) => (
  <ListItem dense style={{ justifyContent: 'center' }}>
    <Chip label={formatDate(date)} />
  </ListItem>
);
