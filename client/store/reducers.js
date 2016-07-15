import _ from 'lodash';

import { SET_INSTANCES, SELECT_INSTANCE, NEXT_INSTANCE } from './states';

export default reducer = (state, action) => {
	state = _.cloneDeep(state);
	if (action.type === SET_INSTANCES) {
		state.instances = action.instances;
	} else if (action.type === SELECT_INSTANCE) {
		state.selected = action.instance;
	} else if (action.type === NEXT_INSTANCE) {
		if (!state.selected) {
			state.selected = state.instances[0];
		} else {
			var index = _.findIndex(state.instances, { name: state.selected.name });
			if (index >= 0) {
				state.selected = state.instances[index+1] || null;
			}
		}
	}

	updateSelection(state);
	return state;
};

function updateSelection(state) {
	if (state.selected) {
		var selected = _.find(state.instances, { name: state.selected.name }) || null;
		if (!selected) {
			state.selected.invalid = true;
			state.selected.memoryUsage = '??';
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