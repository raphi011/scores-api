import React from 'react';
// import Router from 'next/router';

import Card from '@material-ui/core/Card';
import PlayerList from '../../components/volleynet/PlayerList';
import withAuth from '../../containers/AuthContainer';
// import CenteredLoading from '../../components/CenteredLoading';
import Layout from '../../containers/LayoutContainer';
import { loadLadderAction } from '../../redux/actions/entities';
import { ladderVolleynetplayerSelector } from '../../redux/reducers/entities';

import { Player } from '../../types';

interface IProps {
  gender: "M" | "W";
  ladder: Player[];
  loadLadder: (gender: string) => void;
  classes: any;
}

const genderList = ["M", "W"];

class Ranking extends React.Component<IProps> {

  static mapDispatchToProps = {
    loadLadder: loadLadderAction,
  };
  static buildActions({ gender }: IProps) {
    return [loadLadderAction(gender)];
  }

  static getParameters(query) {
    let { gender } = query;

    if (!genderList.includes(gender)) {
      gender = "M"
    }

    return { gender };
  }

  static mapStateToProps(state, { gender }: IProps) {
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
          <PlayerList players={ladder} />
        </Card>
      </Layout>
    );
  }
}

export default withAuth(Ranking);

