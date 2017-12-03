import React from "react";
import fetch from "isomorphic-unfetch";
import { withStyles } from "material-ui/styles";
import Badge from "material-ui/Badge";
import PersonIcon from "material-ui-icons/Person";
import AddCircleIcon from "material-ui-icons/AddCircle";
import withRedux from "next-redux-wrapper";
import Button from "material-ui/Button";
import AddIcon from "material-ui-icons/Add";
import Router from "next/router";
import Tooltip from "material-ui/Tooltip";

import withRoot from "../components/withRoot";
import Layout from "../components/Layout";
import MatchOptionsDialog from "../components/MatchOptionsDialog";
import MatchList from "../components/MatchList";
import initStore, { dispatchActions } from "../redux/store";
import { matchesSelector } from "../redux/reducers/reducer";
import {
  loadMatchesAction,
  setStatusAction,
  deleteMatchAction,
  userOrLoginRouteAction,
} from "../redux/actions/action";

const styles = theme => ({
  matchListContainer: {
    marginBottom: "70px"
  },
  button: {
    margin: theme.spacing.unit,
    position: "fixed",
    right: "24px",
    bottom: "24px"
  }
});

class Index extends React.Component {
  state = {
    selectedMatch: null
  };

  static async getInitialProps({ store, req, res, isServer }) {
    const actions = [loadMatchesAction(), userOrLoginRouteAction()];

    await dispatchActions(store.dispatch, isServer, req, res, actions);
  }

  onCloseDialog = () => {
    this.setState({ selectedMatch: null });
  };

  onOpenDialog = selectedMatch => {
    this.setState({ selectedMatch });
  };

  onCreateMatch = () => {
    Router.replace("/newMatch");
  };

  onDeleteMatch = () => {
    const { matches, deleteMatch } = this.props;
    const { selectedMatch } = this.state;

    deleteMatch(selectedMatch);

    this.setState({ selectedMatch: null });
  };

  onRematch = () => {
    const { setStatus } = this.props;
    const { selectedMatch } = this.state;

    Router.push(`/newMatch?rematchID=${selectedMatch.ID}`)
  };

  render() {
    const { matches, error, classes } = this.props;
    const { selectedMatch } = this.state;

    return (
      <Layout title="Matches">
        <div className={classes.matchListContainer}>
          <MatchList matches={matches} onMatchClick={this.onOpenDialog} />
        </div>
        <MatchOptionsDialog
          open={selectedMatch != null}
          match={selectedMatch}
          onClose={this.onCloseDialog}
          onRematch={this.onRematch}
          onDelete={this.onDeleteMatch}
        />
        <Tooltip title="Create new Match" className={classes.button}>
          <Button
            fab
            color="primary"
            aria-label="add"
            onClick={this.onCreateMatch}
          >
            <AddIcon />
          </Button>
        </Tooltip>
      </Layout>
    );
  }
}

function mapStateToProps(state) {
  const matches = matchesSelector(state);

  return {
    matches,
  };
}

const mapDispatchToProps = {
  loadMatches: loadMatchesAction,
  setStatus: setStatusAction,
  deleteMatch: deleteMatchAction
};

export default withRedux(initStore, mapStateToProps, mapDispatchToProps)(
  withRoot(withStyles(styles)(Index))
);
