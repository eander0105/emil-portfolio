<script lang="ts">
	import * as Navbar from '$c/Navbar';
	import Footer from '$c/Footer.svelte';
	import DarkmodeToggle from '$c/DarkmodeToggle.svelte';

	let { children } = $props();
</script>

<!-- Header -->
<Navbar.Root>
	<Navbar.Link label="Home" link="/" />
	<Navbar.Link label="Blog" link="/blog" />
	<Navbar.Link label="Projects" link="/projects" />
	<Navbar.Link label="Photography" link="/photography" />
	<Navbar.Link label="About" link="/about" />
	<Navbar.Link label="Resume" link="/resume" />
</Navbar.Root>

<!-- Main content -->
<main class="min-h-screen">
	<div class="container mx-auto mt-4 px-4 sm:px-16">
		{@render children()}
	</div>
</main>

<!-- Footer -->
<Footer />
