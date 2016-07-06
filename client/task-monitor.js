import request from 'browser-request';
import { setInstances } from './store/action-creators';

export default class TaskMonitor {
	constructor() {
		this.timeoutId = null;
	}

	init(store) {
		this.store = store;
	}

	start() {
		this.update();
	}

	update() {
		if (this.timeoutId) {
			clearTimeout(this.timeoutId);
			this.timeoutId = null;
		}

		request({
			method: 'GET',
			url: '/status/all',
			json: true
		}, (err, res, data) => {
			if (err) {
				console.error('couldn\'t get statuses: ', err);
			}

			this.store.dispatch(setInstances(data));
			this.timeoutId = setTimeout(() => {
				this.update();
			}, 15000);
		});
	}
}