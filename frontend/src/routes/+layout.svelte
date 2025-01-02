<script lang="ts">
	import '../app.css';
	import Navbar from '$c/nav/Navbar.svelte';
	import Navlink from '$c/nav/Navlink.svelte';
	import Footer from '$c/Footer.svelte';

	let { children } = $props();
</script>

<!-- Header -->
<Navbar>
	<Navlink label="Home" link="/" />
	<Navlink label="About" link="/about" />
	<Navlink label="Blog" link="/blog" />
	<Navlink label="Projects" link="/projects" />
	<Navlink label="Resume" link="/resume" />
</Navbar>

<main class="bg-zinc-900 text-white min-h-screen">
	<div class="container mx-auto p-4">
		{@render children()}
	</div>
</main>

<Footer />
