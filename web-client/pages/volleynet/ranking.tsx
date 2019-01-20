import React from 'react';

import Card from '@material-ui/core/Card';

import Ladder from '../../components/volleynet/Ladder';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import { loadLadderAction } from '../../redux/entities/actions';
import { ladderVolleynetplayerSelector } from '../../redux/entities/selectors';

import { Player } from '../../types';

interface Props {
  gender: 'M' | 'W';
  ladder: Player[];

  loadLadder: (gender: string) => void;
}

const genderList = ['M', 'W'];

class Ranking extends React.Component<Props> {
  static mapDispatchToProps = {
    loadLadder: loadLadderAction,
  };

  static buildActions({ gender }: Props) {
    return [loadLadderAction(gender)];
  }

  static getParameters(query) {
    let { gender } = query;

    if (!genderList.includes(gender)) {
      gender = 'M';
    }

    return { gender };
  }

  static mapStateToProps(state, { gender }: Props) {
    const ladder = ladderVolleynetplayerSelector(state, gender);

    return { ladder };
  }

  componentDidUpdate(prevProps) {
    const { loadLadder, gender } = this.props;

    if (gender !== prevProps.gender) {
      loadLadder(gender);
    }
  }

  render() {
    const { ladder } = this.props;

    return (
      <Layout title={{ text: 'Rankings', href: '' }}>
        <Card>
          <Ladder players={ladder} />
        </Card>
      </Layout>
    );
  }
}

export default withAuth(Ranking);
