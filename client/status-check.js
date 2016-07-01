import request from 'browser-request';
import { setInstances } from './store/action-creators';

export default class StatusChecker {
	constructor(store) {
		this.timeoutId = null;
		this.store = store;
	}

	start() {
		if (this.timeoutId) {
			clearTimeout(this.timeoutId);
			this.timeoutId = null;
		}

		this.update();
	}

	update() {
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