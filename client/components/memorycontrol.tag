import request from 'browser-request';
import _ from 'lodash';
import { resetMemory, setResource } from '../store/action-creators';

<memorycontrol>
	<form onsubmit={ sendUpdate }>
		<div class="form-group">
			<label class="col-md-2 col-md-offset-2">Usage</label>
			<div class="col-md-4">
				<div class="input-group">
					<input type="number" name="usage" min="0" max="100" step="1" value={ opts.usage } class="form-control" required />
					<span class="input-group-addon"><i class="fa fa-percent"></i></span>
				</div>
				
			</div>
		</div>

		<div class="pull-right">
			<button type="submit" class="btn btn-success"><i class="fa fa-check"></i> Update</button>
			<button type="button" class="btn btn-danger" onclick={ resetMemory }><i class="fa fa-refresh"></i> Reset</button>
		</div>
	</form>

	<script>
		this.mixin('redux');
		this.dispatchify({ resetMemory, setResource });

		this.sendUpdate = (e) => {
			this.setResource('memory', this.usage.value);
			e.preventDefault();
			return false;
		};
	</script>
</memorycontrol>