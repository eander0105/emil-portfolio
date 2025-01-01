<script lang="ts">
	let { link, children } = $props();
</script>

<a
	href={link}
	class="px-3 py-2 md:rounded transistion-all duration-100 hover:bg-zinc-800"
>
	{@render children?.()}
</a>