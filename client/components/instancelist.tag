import _ from 'lodash';
import { selectInstance } from '../store/action-creators';

<instancelist>
	<div class="list-group">
		<div role="button" each={ instances } onclick={ select } class={ list-group-item: true, active: selected, list-group-item-danger: status.invalid }>
			<h5 class="instance-title"><span title="invalid instance" show={ status.invalid }><i class="fa fa-exclamation-triangle"></i></span> { name }</h5>

			<div class="row stats" hide={ status.invalid }>
				<div class="col-md-4 text-center">
					MEM <br />{ memoryUsage }%
				</div>
				<div class="col-md-4 text-center">
					DISK <br />{ diskUsage } MB/sec
				</div>
				<div class="col-md-4 text-center"></div>
			</div>
		</div>
	</div>

	<script>
		this.mixin('redux');
		this.dispatchify({ selectInstance });

		this.subscribe(state => { 
			return { 
				instances: state.instances
			};
		});

		this.select = function () {
			if (!this.selected && !this.invalid) {
				this.selectInstance(this._item);
			}

			return false;
		};
	</script>
</instancelist>