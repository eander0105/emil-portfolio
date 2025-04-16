<script lang="ts">
	import '../app.css';
	import * as Navbar from '$c/Navbar';
	import Footer from '$c/Footer.svelte';

	let { children } = $props();
</script>

<!-- Header -->
<Navbar.Root>
	<Navbar.Link label="Home" link="/" />
	<Navbar.Link label="About" link="/about" />
	<Navbar.Link label="Blog" link="/blog" />
	<Navbar.Link label="Projects" link="/projects" />
	<Navbar.Link label="Resume" link="/resume" />
</Navbar.Root>

<!-- Main content -->
<main class="min-h-screen bg-zinc-900 text-white">
	<div class="container mx-auto px-4 sm:px-16">
		{@render children()}
	</div>
</main>

<!-- Footer -->
<Footer />
