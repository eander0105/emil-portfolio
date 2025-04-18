<script lang="ts">
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
	import * as Navbar from '$c/Navbar';
	import DarkmodeToggle from '$c/DarkmodeToggle.svelte';
	import Footer from '$c/Footer.svelte';

	let { children } = $props();
</script>

<ModeWatcher defaultMode={"dark"}/>

<!-- Header -->
<Navbar.Root>
	<Navbar.Link label="Home" link="/" />
	<Navbar.Link label="Blog" link="/blog" />
	<Navbar.Link label="Projects" link="/projects" />
	<Navbar.Link label="Photography" link="/photography" />
	<Navbar.Link label="About" link="/about" />
	<Navbar.Link label="Resume" link="/resume" />
</Navbar.Root>

<!-- <DarkmodeToggle /> -->

<!-- Main content -->
<main class="min-h-screen">
	<div class="container mx-auto px-4 sm:px-16">
		{@render children()}
	</div>
</main>

<!-- Footer -->
<Footer />
