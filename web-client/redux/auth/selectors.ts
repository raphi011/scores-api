import { Store } from '../store';

export const loginRouteSelector = (state: Store) => state.auth.loginRoute;

export const userSelector = (state: Store) => ({
  isLoggedIn: !!state.auth.user,
  user: state.auth.user,
});
