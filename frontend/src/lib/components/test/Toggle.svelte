<script lang="ts">
	import { Toggle } from 'melt/components';

	let value = $state(false);
</script>

<Toggle bind:value>
	{#snippet children(toggle)}
		<button {...toggle.trigger}>
			{toggle.value ? 'On' : 'Off'}
		</button>
	{/snippet}
</Toggle>
