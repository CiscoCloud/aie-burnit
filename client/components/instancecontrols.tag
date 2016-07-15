import { nextInstance } from '../store/action-creators';

<instancecontrols>
	<ol class="breadcrumb">
		<li><i class="fa fa-server"></i> instances</li>
		<li>{ instance.name }</li>
	</ol>

	<div class="alert alert-block alert-warning" show={ instance.invalid }>
		<i class="fa fa-exclamation-triangle"></i> Sorry, <strong>{ instance.name }</strong> is no longer a valid instance. Select the <a onclick={ nextInstance } role="button">next available</a>.
	</div>

	<div hide={ instance.invalid }>
		<h4><i class="fa fa-braille"></i> Memory</h4>
		<hr />
		<memorycontrol usage={ instance.memoryUsage }></memorycontrol>
	</div>

	<div class="clearfix"></div>

	<div hide={ instance.invalid }>
		<h4><i class="fa fa-braille"></i> Disk</h4>
		<hr />
		<diskcontrol usage={ instance.diskUsage }></diskcontrol>
	</div>

	<script>
		this.mixin('redux');
		this.dispatchify({ nextInstance });

		this.subscribe(state => {
			return { instance: state.selected };
		})
	</script>
</instancecontrols>