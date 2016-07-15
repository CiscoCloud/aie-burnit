import request from 'browser-request';
import _ from 'lodash';
import { resetDisk, setResource } from '../store/action-creators';

<diskcontrol>
	<form onsubmit={ sendUpdate }>
		<div class="form-group">
			<label class="col-md-2 col-md-offset-2">Usage (Megabytes)</label>
			<div class="col-md-3">
				<input type="number" name="usage" min="0" max="100" step="10" value={ opts.usage } class="form-control" required />
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

		this.sendUpdate = (e) => {
			this.setResource('disk', this.usage.value);
			e.preventDefault();
			return false;
		};
	</script>
</diskcontrol>