import React, { ChangeEvent } from 'react';

import { connect } from 'react-redux';
import EditUserDialog from '../../components/admin/EditUserDialog';
import { updateUserAction } from '../../redux/admin/actions';
import { User } from '../../types';

interface Props {
  user?: User;
  open: boolean;

  onClose: () => void;
  updateUser: (email: string, password: string) => Promise<null>;
}

interface State {
  email: string;
  password: string;
  submitting: boolean;
}

class EditUserDialogContainer extends React.Component<Props, State> {
  state = {
    email: '',
    password: '',
    submitting: false,
  };

  componentDidUpdate(prevProps) {
    const { user } = this.props;

    if (prevProps.user === user) {
      return;
    }

    if (user) {
      this.setState({ 
        email: user.email,
        password: "",
      });
    }
  }

  onChangeEmail = (event: ChangeEvent<HTMLInputElement>) => {
    this.setState({ email: event.currentTarget.value });
  };

  onChangePassword = (event: ChangeEvent<HTMLInputElement>) => {
    this.setState({ password: event.currentTarget.value });
  };

  onSubmit = async () => {
    const { onClose, updateUser } = this.props;
    const { email, password } = this.state;

    await updateUser(email, password);

    onClose();
  };

  canSubmit = () => {
    const { email, password, submitting } = this.state;

    return !!email && !!password && !submitting;
  };

  render() {
    const { user, onClose, open } = this.props;

    const isNew = !user || !user.id;

    return (
      <EditUserDialog
        canSubmit={this.canSubmit()}
        email={this.state.email}
        isNew={isNew}
        open={open}
        password={this.state.password}
        user={user}
        onSubmit={this.onSubmit}
        onClose={onClose}
        onChangeEmail={this.onChangeEmail}
        onChangePassword={this.onChangePassword}
      />
    );
  }
}

const mapDispatchToProps = {
  updateUser: updateUserAction,
};

export default connect(
  null,
  mapDispatchToProps,
  // @ts-ignore
)(EditUserDialogContainer);
