import _ from 'lodash';
import { createStore, applyMiddleware } from 'redux';

import reducer from './reducers';
import * as wares from './middleware';

const defaultState = {
	instances: [],
	selected: null
};

export default function storeFactory(...middlewares) {
	return createStore(reducer, defaultState, applyMiddleware(...middlewares));
}
