import * as http from 'http';
import { Params } from '../../api';
import * as actionNames from '../actionNames';
import { Action } from '../actions';
import { ReceiveEntitiesAction } from '../entities/actions';

export const multiApiAction = (actions: ApiAction[]): ApiActions => ({
  actions,
  type: actionNames.API_MULTI,
});

export interface ApiResponse<T> {
  status: number;
  message: string;
  payload: T;
}

export type ApiActionType = ApiAction | ApiActions;

export interface ApiAction extends Action {
  type: 'API';
  method: string;
  url: string;
  success?: string;
  isServer?: boolean;
  params?: Params;
  req?: http.IncomingMessage;
  res?: http.ServerResponse;
  headers?: { [key: string]: string };
  error?: string;
  body?: string;
  successStatus?: string;
  successParams?: ReceiveEntitiesAction;
}

export interface ApiActions extends Action {
  type: 'API_MULTI';
  actions: ApiAction[];
  req?: http.IncomingMessage;
  res?: http.ServerResponse;
  isServer?: boolean;
}
