<script lang="ts">
	import * as Navbar from '$components/Navbar';

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
	<div class="container max-w-[1024px] mx-auto mt-4 px-4 sm:px-16">
		{@render children()}
	</div>
</main>
