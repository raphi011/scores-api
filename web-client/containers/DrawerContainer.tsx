import React from 'react';

import { withWidth } from '@material-ui/core';
import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';

import Drawer from '../components/Drawer';

interface Props {
  width: Breakpoint;
  open: boolean;
}

interface State {
  open: boolean;
}

class DrawerContainer extends React.Component<Props, State> {
  constructor(props) {
    super(props);

    this.state = {
      open: props.open,
    };
  }

  componentDidUpdate(prevProps: Props) {
    if (prevProps.width === 'xs' && this.props.width !== 'xs') {
      this.setState({ open: false });
    } else if (prevProps.open !== this.props.open) {
      this.setState({
        open: !this.state.open,
      });
    }
  }

  onToggleDrawer = () => {
    this.setState({ open: !this.state.open });
  };

  onOpenDrawer = () => {
    this.setState({ open: true });
  };

  onClose = () => {
    this.setState({ open: false });
  };

  render() {
    const { width } = this.props;
    const { open } = this.state;

    return (
      <Drawer
        mobile={['xs', 'sm'].includes(width)}
        open={open}
        onClose={this.onClose}
      />
    );
  }
}

export default withWidth()(DrawerContainer);
