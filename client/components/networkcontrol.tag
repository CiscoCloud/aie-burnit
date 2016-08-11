import { toggleTraffic, startTraffic, stopTraffic } from '../store/action-creators';

<networkcontrol>
	<form class="form-horizontal" onsubmit={ startSim } show={ sim.enabled && !sim.active }>
		<div class="form-group">
			<label class="col-md-6 text-right">Hits</label>
			<div class="col-md-6">
				<input type="number" name="hits" min="0" value="0" class="form-control" required />
			</div>
		</div>

		<div class="form-group">
			<label class="col-md-6 text-right">Delay (ms)</label>
			<div class="col-md-6">
				<input type="number" name="delay" min="0" value="0" class="form-control" required />
			</div>
		</div>

		<div class="form-group">
			<label class="col-md-6 text-right">Status</label>
			<div class="col-md-6">
				<input type="number" name="statusCode" min="0" value="0" class="form-control" required />
			</div>
		</div>

		<button type="button" class="btn btn-success" onclick={ startSim }><i class="fa fa-check"></i> Start</button>
	</form>

	<section hide={ sim.enabled } class="text-center">
		<button type="button" onclick={ setEnabled } class="btn btn-primary">Enable</button>
	</section>

	<section show={ sim.active }>
		Executing request(s)
		<div class="progress">
			<div class="progress-bar" role="progressbar" style="width:{sim.completed_pct}%">
				<span class="sr-only">{sim.completed_pct}% Complete</span>
			</div>
		</div>

		<button type="button" class="btn btn-danger" onclick={ stopSim }><i class="fa fa-refresh"></i> Stop</button>
	</section>

	<script>
		this.mixin('redux');
		this.dispatchify({ toggleTraffic, startTraffic, stopTraffic });

		this.subscribe(state => {
			let traffic = state.traffic;

			return {
				sim: {
					enabled: traffic.enabled,
					active: traffic.active,
					goal: traffic.config.hitCount,
					completed: traffic.completed,
					completed_pct: traffic.completed_pct
				}
			};
		});

		this.setEnabled = () => {
			this.toggleTraffic(true);
		};

		this.startSim = (e) => {
			this.startTraffic({
				statusCode: parseInt(this.statusCode.value),
				delayMs: parseInt(this.delay.value),
				hitCount: parseInt(this.hits.value)
			});

			e.preventDefault();
			return false;
		};

		this.stopSim = () => {
			this.stopTraffic();
		};
	</script>
</networkcontrol>