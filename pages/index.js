import React from "react";
import fetch from "isomorphic-unfetch";
import Badge from "material-ui/Badge";
import PersonIcon from "material-ui-icons/Person";
import AddCircleIcon from "material-ui-icons/AddCircle";
import withRedux from "next-redux-wrapper";

import withRoot from "../components/withRoot";
import Layout from "../components/Layout";
import MatchOptionsDialog from "../components/MatchOptionsDialog";
import MatchList from "../components/MatchList";
import initStore from "../redux/store";
import { matchesSelector } from "../redux/reducers/reducer";
import { loadMatchesAction, setStatusAction, deleteMatchAction } from "../redux/actions/action";

class Index extends React.Component {
  state = {
    loginRoute: "",
    selectedMatch: null
  };

  static async getInitialProps({ store }) {
    await store.dispatch(loadMatchesAction());
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

  onOpenDialog = selectedMatch => {
    this.setState({ selectedMatch });
  };

  onDeleteMatch = () => {
      const { matches, deleteMatch } = this.props;
      const { selectedMatch } = this.state;

      deleteMatch(selectedMatch);

      this.setState({ selectedMatch: null });
  };

  onCloneMatch = () => {
    const { setStatus } = this.props;
    setStatus("Not implemented yet");
    this.setState({ selectedMatch: null });
  };

  render() {
    const { matches, error } = this.props;
    const { loginRoute, selectedMatch } = this.state;

    return (
      <Layout title="Matches" loginRoute={loginRoute}>
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

function mapStateToProps(state) {
  const matches = matchesSelector(state);

  return {
    matches
  };
}

const mapDispatchToProps = {
  loadMatches: loadMatchesAction,
  setStatus: setStatusAction,
  deleteMatch: deleteMatchAction,
};

export default withRedux(initStore, mapStateToProps, mapDispatchToProps)(
  withRoot(Index)
);
