import { SET_INSTANCES, SET_RESOURCE, SELECT_INSTANCE, RESET_MEMORY } from './states';
import _ from 'lodash';

export function setInstances(instances) {
	instances = instances || [];
	return {
		type: SET_INSTANCES,
		instances: _.sortBy(instances.map(i => {
			var c = _.cloneDeep(i);
			c.memoryUsage = c.memory_usage;
			c.selected = false;
			delete c.memory_usage;
			return c;
		}), 'name')
	};
}

export function resetMemory() {
	return {
		type: RESET_MEMORY,
		remote: true,
		params: {
			resource: 'memory',
			value: '0'
		}
	};
}

export function setResource(resource, value) {
	return {
		type: SET_RESOURCE,
		remote: true,
		params: {
			resource: resource,
			value: value
		}
	};
}

export function selectInstance(instance) {
	return {
		type: SELECT_INSTANCE,
		instance: instance
	};
}