import { ApiAction, ApiActions } from '../../types';
import * as actionNames from '../actionNames';

// eslint-disable-next-line import/prefer-default-export
export const multiApiAction = (actions: ApiAction[]): ApiActions => ({
  type: actionNames.API_MULTI,
  actions,
});
