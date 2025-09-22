<script lang=ts>
	import type { Snippet } from "svelte";

	interface Props {
		children?: Snippet;
		href: string;
	};

	let { children, href } : Props = $props();
	let link: string;

	if (!href.startsWith('/') && !href.startsWith('https://')) {
		link = `https://${href}`;
	} else {
		link = href;
	}
</script>

<a class="underline hover:text-blue-400"
href={link}>
	{#if children}
		{@render children()}
	{:else}
		{href}
	{/if}
</a>
