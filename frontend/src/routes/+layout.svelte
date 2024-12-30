<script lang="ts">
	import '../app.css';
	import Nav from '$c/nav/Nav.svelte';
	import Footer from '$c/Footer.svelte';

	let { children } = $props();
</script>

<!-- Header -->
<Nav>

</Nav>

<main class="bg-zinc-900 text-white min-h-screen">
	<div class="container mx-auto py-4">
		{@render children()}
	</div>
</main>

<Footer />
