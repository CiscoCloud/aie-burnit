<instancecontrols>
	<ol class="breadcrumb">
		<li><i class="fa fa-server"></i> instances</li>
		<li>{ instance.name }</li>
	</ol>

	<h4><i class="fa fa-braille"></i> Memory</h4>
	<hr />
	<memorycontrol usage={ instance.memoryUsage }></memorycontrol>

	<script>
		this.mixin('redux');

		this.subscribe(state => {
			return { instance: state.selected };
		})
	</script>
</instancecontrols>