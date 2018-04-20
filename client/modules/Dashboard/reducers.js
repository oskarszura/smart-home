import _ from 'lodash';
import * as actionTypes from './actionTypes';

const defaultState = {
  isAlerts: false,
  times: [],
  temperatures: [],
  motions: {},
  gas: false,
  error: '',
};

export default function reducer(state = defaultState, action) {
  switch (action.type) {
    case actionTypes.DATA_FETCH:
      return Object.assign({}, state);

    case actionTypes.DATA_FETCH_SUCCESS:
      const { times, temperatures, motions, gas } = action;

      return Object.assign({}, state, {
        times,
        temperatures,
        motions,
        gas,
        error: '',
      });

    case actionTypes.DATA_FETCH_ERROR:
      return Object.assign({}, state, { error: action.error });

    case actionTypes.SET_ALERTS:
      return Object.assign({}, state, { isAlerts: action.isAlerts });

    default:
      return state;
  }
}
