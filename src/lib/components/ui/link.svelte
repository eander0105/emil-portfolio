<script>
	let { children, href } = $props();
</script>
<a class="underline hover:text-blue-400"
href={href}>
	{@render children()}
</a>
