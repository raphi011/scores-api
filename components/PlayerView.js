import React from "react";
import { withStyles } from "material-ui/styles";

const styles = theme => ({
});

function PlayerView({ player }) {
  return (
    <div>
      {player.Name}
    </div>
  );
}

export default withStyles(styles)(PlayerView);
