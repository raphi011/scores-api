import React from 'react';
import { connect } from 'react-redux';

import { Dialog, Typography } from '@material-ui/core';

import Login from '../components/volleynet/Login';
import { Store } from '../redux/store';
import { userSelector } from '../redux/auth/selectors';
import { volleynetLoginAction } from '../redux/entities/actions';
import { User } from '../types';

interface Props {
  user: User | null;

  volleynetLogin: (username: string, password: string) => Promise<void>;
}

class MandatoryVolleynetLogin extends React.Component<Props> {
  volleynetSignup = async (username: string, password: string) => {
    const { volleynetLogin } = this.props;

    await volleynetLogin(username, password);
  };

  render() {
    const { open } = this.props;
    return (
      <Dialog open={open}>
        <div style={{ padding: '20px' }}>
          <Typography>
            Please login using your <i>volleynet.at</i> credentials so that we
            know who you are :). The password is
            <strong> not</strong> saved and has to be reentered when signing up
            for a tournament.
          </Typography>
          <Login showRememberMe={false} onLogin={this.volleynetSignup} />
        </div>
      </Dialog>
    );
  }
}

function mapStateToProps(state: Store) {
  const user = userSelector(state);

  return {
    open: user && !user.volleynetUserId,
  };
}

const mapDispatchToProps = {
  volleynetLogin: volleynetLoginAction,
};

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(MandatoryVolleynetLogin);
