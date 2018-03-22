// @flow

import * as actionNames from '../actionNames';
import type { ApiAction, ApiActions } from '../../types';

// eslint-disable-next-line import/prefer-default-export
export const multiApiAction = (actions: Array<ApiAction>): ApiActions => ({
  type: actionNames.API_MULTI,
  actions,
});
