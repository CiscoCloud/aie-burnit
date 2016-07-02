import _ from 'lodash';
import { selectInstance } from '../store/action-creators';

<instancelist>
	<div class="list-group">
		<a role="button" each={ instances } onclick={ select } class={ list-group-item: true, active: selected }>{ name } <span class="badge">RAM { memoryUsage }%</span></a>
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
			if (!this.selected) {
				this.selectInstance(this._item);
			}

			return false;
		};
	</script>
</instancelist>