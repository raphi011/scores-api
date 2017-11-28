import React from "react";
import fetch from "isomorphic-unfetch";
import Badge from "material-ui/Badge";
import PersonIcon from "material-ui-icons/Person";
import AddCircleIcon from "material-ui-icons/AddCircle";

import withRoot from "../components/withRoot";
import Layout from "../components/Layout";
import MatchOptionsDialog from "../components/MatchOptionsDialog";
import MatchList from "../components/MatchList";

class Index extends React.Component {
  state = {
    loginRoute: "",
    selectedMatch: null,
    status: "",
  };

  static async getInitialProps() {
    const matchResponse = await fetch(`${process.env.BACKEND_URL}/api/matches`);
    const matches = await matchResponse.json();

    return { matches };
  }

  async componentDidMount() {
    const routeResponse = await fetch(
      `${process.env.BACKEND_URL}/api/loginRoute`,
      { credentials: "same-origin" }
    );

    const loginRoute = await routeResponse.json();

    this.setState({ loginRoute });
  }

  onCloseDialog = () => {
    this.setState({ selectedMatch: null });
  };

  onOpenDialog = (selectedMatch) => {
    this.setState({ selectedMatch });
  };

  onDeleteMatch = async () => {
    try {
      const { matches } = this.props;
      const { selectedMatch } = this.state;

      await fetch(`${process.env.BACKEND_URL}/api/matches/${selectedMatch.ID}`, {
        method: "DELETE"
      });

      // BAD: maybe put matches into state instead of props
      matches.splice(matches.indexOf(selectedMatch), 1);
      this.setState({ selectedMatch: null, status: "Deleted Match"});
    } catch (e) {
      this.setState({ status: "Error deleting Match"});
    }
  };

  onCloneMatch = () => {
    this.setState({ selectedMatch: null, status: "Not implemented yet"});
  };

  render() {
    const { matches, error } = this.props;
    const { loginRoute, selectedMatch } = this.state;

    return (
      <Layout status={this.state.status} title="Matches" loginRoute={loginRoute}>
        <MatchList matches={matches} onMatchClick={this.onOpenDialog} />
        <MatchOptionsDialog
          open={selectedMatch != null}
          match={selectedMatch}
          onClose={this.onCloseDialog}
          onClone={this.onCloneMatch}
          onDelete={this.onDeleteMatch}
        />
      </Layout>
    );
  }
}

export default withRoot(Index);
