import { resetDisk, setResource } from '../store/action-creators';

<diskcontrol>
	<form onsubmit={ sendUpdate }>
		<div class="form-group">
			<label class="col-md-2 col-md-offset-2">Usage</label>
			<div class="col-md-4">
				<div class="input-group">
					<input type="number" name="usage" min="1" max="100" step="1" value={ opts.usage } class="form-control" required />
					<span class="input-group-addon"><strong>MBs</strong></span>
				</div>
			</div>
		</div>

		<div class="pull-right">
			<button type="submit" class="btn btn-success"><i class="fa fa-check"></i> Update</button>
			<button type="button" class="btn btn-danger" onclick={ resetDisk }><i class="fa fa-refresh"></i> Reset</button>
		</div>
	</form>

	<script>
		this.mixin('redux');
		this.dispatchify({ resetDisk, setResource });

		this.sendUpdate = () => {
			this.setResource('disk', this.usage.value);
		};
	</script>
</diskcontrol>