import * as http from 'http';
import { IParams } from '../../api';
import * as actionNames from '../actionNames';
import { Action } from '../actions';

// eslint-disable-next-line import/prefer-default-export
export const multiApiAction = (actions: ApiAction[]): ApiActions => ({
  actions,
  type: actionNames.API_MULTI,
});

export interface ApiAction extends Action {
  type: 'API';
  method: string;
  url: string;
  success?: string;
  isServer?: boolean;
  params?: IParams;
  req?: http.IncomingMessage;
  res?: http.ServerResponse;
  headers?: { [key: string]: string };
  error?: string;
  body?: string;
  successStatus?: string;
  successParams?: object;
}

export interface ApiActions extends Action {
  type: 'API_MULTI';
  actions: ApiAction[];
  req?: http.IncomingMessage;
  res?: http.ServerResponse;
  isServer?: boolean;
}
