import _ from 'lodash';
import { createStore, applyMiddleware } from 'redux';

import reducer from './reducers';
import * as wares from './middleware';

const defaultState = {
	instances: [],
	selected: null
};

var middleware = applyMiddleware(wares.logging, wares.relay);
export default createStore(reducer, defaultState, middleware);
