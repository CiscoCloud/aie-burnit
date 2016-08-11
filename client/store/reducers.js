import _ from 'lodash';

import * as states from './states';

export default reducer = (state, action) => {
	state = _.cloneDeep(state);

	if (action.type === states.SET_INSTANCES) {
		state.instances = action.instances;
	} else if (action.type === states.SELECT_INSTANCE) {
		state.selected = action.instance;
	} else if (action.type === states.NEXT_INSTANCE) {
		if (!state.selected) {
			state.selected = state.instances[0];
		} else {
			var index = _.findIndex(state.instances, { name: state.selected.name });
			if (index >= 0) {
				state.selected = state.instances[index+1] || null;
			}
		}
	} else if (action.type === states.TRAFFIC_STARTED) {
		state.traffic = {
			enabled: true,
			active: true,
			config: action.config,
			completed: 0,
			completed_pct: 0.0,
			hits: []
		};
	} else if (action.type === states.TRAFFIC_HIT) {
		let goal = state.traffic.config.hitCount;
		let completed = state.traffic.completed + 1;
		let completed_pct = (goal > 0 ? (completed/goal) : 0.0) * 100.0;
		let hits = state.traffic.hits.concat([{ statusCode: action.statusCode }]);
		state.traffic = _.merge(state.traffic, {
			hits,
			completed,
			completed_pct
		});
	} else if (action.type === states.TRAFFIC_DONE) {
		state.traffic.active = false;
	} else if (action.type === states.TRAFFIC_TOGGLE) {
		state.traffic.enabled = action.enabled;
	}

	updateSelection(state);
	return state;
};

function updateSelection(state) {
	if (state.selected) {
		var selected = _.find(state.instances, { name: state.selected.name }) || null;
		if (!selected) {
			state.selected.status = {
				invalid: true,
				name: 'missing',
				message: 'does not exist'
			};
			
			state.instances.unshift(_.clone(state.selected));
		} else {
			selected.selected = true;
			state.selected = _.clone(selected);
		}
	} else if (state.instances.length > 0) {
		state.instances[0].selected = true;
		state.selected = _.clone(state.instances[0]);
	}

	state.instances.forEach(i => {
		if (!state.selected || state.selected.name !== i.name) {
			i.selected = false;
		}
	});
}