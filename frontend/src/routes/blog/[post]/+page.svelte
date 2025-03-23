<script lang="ts">
	import type { PageData } from './$types';
	let props : PageData = $props();
</script>

Hello! {props.data.title}
