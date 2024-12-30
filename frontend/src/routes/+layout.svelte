<script lang="ts">
	import '../app.css';
	import Header from '../components/Header.svelte';
	import Footer from '../components/Footer.svelte';

	let { children } = $props();
</script>

<Header />

<main class="bg-zinc-900 text-white min-h-screen">
	<div class="container mx-auto py-4">
		{@render children()}
	</div>
</main>

<Footer />
