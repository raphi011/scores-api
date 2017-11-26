import React from "react";
import fetch from "isomorphic-unfetch";
import Badge from "material-ui/Badge";
import PersonIcon from "material-ui-icons/Person";
import AddCircleIcon from "material-ui-icons/AddCircle";

import Layout from '../components/Layout';
import MatchList from '../components/MatchList';


export default class extends React.Component {
  static async getInitialProps() {
    try {
      const matchResponse = await fetch("http://localhost:3000/api/matches");
      const matches = await matchResponse.json();

      return { matches };
    } catch (error) {
      return { error };
    }
  }

  render() {
    const { matches, error } = this.props;

    return (
      <Layout>
        <h1>Matches</h1>
        <MatchList matches={matches} />
      </Layout>
    );
  }
}
