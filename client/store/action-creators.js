import * as states from './states';
import _ from 'lodash';

export function setInstances(instances) {
	instances = instances || [];
	return {
		type: states.SET_INSTANCES,
		instances: _.sortBy(instances.map(i => {
			var c = _.cloneDeep(i);
			c.memoryUsage = c.memory_usage;
			c.diskUsage = c.disk_usage;
			c.selected = false;
			delete c.memory_usage;
			delete c.disk_usage;
			return c;
		}), 'status.name')
	};
}

export function resetMemory() {
	return {
		type: states.RESET_MEMORY,
		remote: true,
		params: {
			resource: 'memory',
			value: '0'
		}
	};
}

export function resetDisk() {
	return {
		type: states.RESET_DISK,
		remote: true,
		params: {
			resource: 'disk',
			value: '0'
		}
	};
}

export function setResource(resource, value) {
	return {
		type: states.SET_RESOURCE,
		remote: true,
		params: {
			resource: resource,
			value: value
		}
	};
}

export function selectInstance(instance) {
	return {
		type: states.SELECT_INSTANCE,
		instance: instance
	};
}

export function nextInstance() {
	return {
		type: states.NEXT_INSTANCE
	};
}

export function recordTrafficHit(statusCode) {
	return {
		type: states.TRAFFIC_HIT,
		timestamp: Date.now(),
		statusCode
	};
}

export function startTraffic(config) {
	return {
		type: states.TRAFFIC_GO,
		config: _.cloneDeep(config)
	};
}

export function trafficStarted(config) {
	return {
		type: states.TRAFFIC_STARTED,
		config: _.cloneDeep(config)
	}
}

export function stopTraffic() {
	return {
		type: states.TRAFFIC_STOP
	};
}

export function trafficDone() {
	return {
		type: states.TRAFFIC_DONE
	};
}

export function toggleTraffic(enabled) {
	return {
		type: states.TRAFFIC_TOGGLE,
		enabled: enabled
	};
}