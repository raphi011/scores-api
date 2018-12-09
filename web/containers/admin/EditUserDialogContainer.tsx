import React, { ChangeEvent } from 'react';

import { connect } from 'react-redux';
import EditUserDialog from '../../components/admin/EditUserDialog';
import { User } from '../../types';

type Props = {
  user?: User;
  open: boolean;
  onClose: () => void;
};

type State = {
  email: string;
  password: string;
};

class EditUserDialogContainer extends React.Component<Props, State> {
  state = {
    email: '',
    password: '',
  };

  componentDidUpdate(prevProps) {
    const { user } = this.props;

    if (prevProps.user === user) {
      return;
    }

    const newState = { email: '', password: '' };

    if (user) {
      newState.email = user.email;
    }

    this.setState(newState);
  }

  onChangeEmail = (event: ChangeEvent<HTMLInputElement>) => {
    this.setState({ email: event.currentTarget.value });
  };

  onChangePassword = (event: ChangeEvent<HTMLInputElement>) => {
    this.setState({ password: event.currentTarget.value });
  };

  render() {
    const { user, onClose, open } = this.props;

    const isNew = !user || !user.id;

    return (
      <EditUserDialog
        onClose={onClose}
        isNew={isNew}
        user={user}
        open={open}
        email={this.state.email}
        onChangeEmail={this.onChangeEmail}
        password={this.state.password}
        onChangePassword={this.onChangePassword}
      />
    );
  }
}

function mapStateToProps(state) {
  // const { isLoggedIn } = userSelector(state);
  // const loginRoute = loginRouteSelector(state);

  return {
    // isLoggedIn,
    // loginRoute,
  };
}

const mapDispatchToProps = {
  // logout: logoutAction,
};

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(EditUserDialogContainer);
