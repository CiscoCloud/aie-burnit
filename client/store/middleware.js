import request from 'browser-request';
import _ from 'lodash';

export const logging = store => next => action => {
    console.info('[store]', action.type, action);
    return next(action);
};

export const relay = store => next => action => {
	if (action.remote) {
		var req = action.params;
		req.host = req.host || store.getState().selected.host;
		request({
			method: 'POST',
			url: '/update',
			json: req
		}, _.noop);		
	}

	return next(action);
};