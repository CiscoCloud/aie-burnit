import _ from 'lodash';
import { selectInstance } from '../store/action-creators';

<instancelist>
	<div class="list-group">
		<a role="button" each={ instances } onclick={ select } class={ list-group-item: true, active: selected, list-group-item-danger: status.invalid }>
		<span title="invalid instance" show={ status.invalid }><i class="fa fa-exclamation-triangle"></i></span> { name } 
			<span hide={ status.invalid } class="badge">RAM { memoryUsage }%</span>
			<span hide={ status.invalid } class="badge">DISK { diskUsage } MB/sec</span>
		</a>
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