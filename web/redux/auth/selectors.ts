import { IStore } from '../store';

export const loginRouteSelector = (state: IStore) => state.auth.loginRoute;

export const userSelector = (state: IStore) => ({
  isLoggedIn: !!state.auth.user,
  user: state.auth.user,
});
