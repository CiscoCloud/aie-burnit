import request from 'browser-request';

import { trafficDone, recordTrafficHit, trafficStarted } from './store/action-creators';

export default class TrafficSimulator {
	constructor() {
		this.run = null;
	}

	init(store) {
		this.store = store;
	}

	start(config = {}) {
		if (this.run) {
			this.stop();
		}

		this.run = {
			req: null,
			completed: 0,
			hitCount: config.hitCount,
			statusCode: config.statusCode,
			delayMs: config.delayMs
		};

		this.nextRequest();
		this.store.dispatch(trafficStarted(config));
	}

	nextRequest() {
		var run = this.run;
		if (!run) {
			console.warn('TrafficSimulator:', 'no run parameters set');
			return;
		}

		run.completed++;
		run.req = request({
			url: '/traffic',
			qs: { 
				status: run.statusCode,
				delay: run.delayMs 
			}
		}, (err, resp) => {
			this.store.dispatch(recordTrafficHit(resp.statusCode));
			if (run.completed >= run.hitCount) {
				this.stop();
			} else {
				this.nextRequest();
			}
		});
	}

	stop() {
		if (this.run) {
			if (this.run.req) {
				this.run.req.abort();
			}

			this.run = null;
			this.store.dispatch(trafficDone());
		}
	}
}