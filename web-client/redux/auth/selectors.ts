import { Store } from '../store';

export const loginRouteSelector = (state: Store) => state.auth.loginRoute;

export const userSelector = (state: Store) => state.auth.user;
