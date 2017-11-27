import React from "react";
import fetch from "isomorphic-unfetch";
import Badge from "material-ui/Badge";
import PersonIcon from "material-ui-icons/Person";
import AddCircleIcon from "material-ui-icons/AddCircle";

import withRoot from '../components/withRoot';
import Layout from "../components/Layout";
import MatchList from "../components/MatchList";

class Index extends React.Component {
  static async getInitialProps() {
    const matchResponse = await fetch(`${process.env.BACKEND_URL}/api/matches`);
    const matches = await matchResponse.json();

    return { matches };
  }

  render() {
    const { matches, error } = this.props;

    return (
      <Layout title="Matches" >
        <MatchList matches={matches} />
      </Layout>
    );
  }
}

export default withRoot(Index)
