<script lang="ts">
	import { burgerOpen } from "$s/Navbar";
	let { children } = $props();
</script>

<nav class="sticky top-0 flex flex-row justify-between bg-zinc-900 px-4 pt-6 pb-3 sm:py-6 text-white container mx-auto">
	<!-- Name, subtitle & Image -->
	<div class="flex flex-row items-center">
		<!-- Image -->
		<img src="https://github.com/eander0105.png" alt="Emil Andersson" class="h-16 w-16 sm:h-20 sm:w-20 rounded">
		<div class="ml-4 flex flex-col">
			<a href="/" class="text-3xl md:text-4xl font-bold">Emil Andersson</a>
			<!-- TODO: select and style a nice subtitle -->
			<!-- <span class="text-md md:text-xl text-zinc-400">Developer with a few to many interests</span> -->
			<!-- <span class="text-md md:text-xl text-zinc-400">Developer with a few to many interests</span> -->
			<span class="text-md md:text-xl text-zinc-400">Software developer</span>
		</div>
	</div>

	<!-- Burger menu -->
	<button
		class="group z-50 flex h-8 flex-col justify-evenly my-auto sm:hidden"
		aria-label="Toggle menu"
		onclick={() => ($burgerOpen = !$burgerOpen)}
	>
		<span
			id="s1"
			class={'translate h-0.5 w-8 rounded-xl bg-white ' +
				($burgerOpen ? 'translate-y-[0.54rem] rotate-45' : '')}
		>
			&nbsp;
		</span>
		<span
			id="s2"
			class={'h-0.5 w-8 rounded-xl bg-white ' + ($burgerOpen ? 'scale-50 opacity-0' : '')}
		>
			&nbsp;
		</span>
		<span
			id="s3"
			class={'translate h-0.5 w-8 rounded-xl bg-white ' +
				($burgerOpen ? '-translate-y-[0.54rem] -rotate-45' : '')}
		>
			&nbsp;
		</span>
	</button>

	<!-- TODO: think about a better looking way to display links on smaller screens(phones) -->
	<div
		class={'links absolute left-0 top-24 flex h-screen w-full flex-col p-5 transition-all duration-200 sm:static sm:left-auto sm:top-auto sm:h-auto sm:w-auto sm:flex-row sm:p-0 ' +
			($burgerOpen
				? 'visible translate-y-0 bg-zinc-900 opacity-100 ease-in-out'
				: 'invisible -translate-y-4 opacity-0 sm:visible sm:translate-y-0 sm:opacity-100')}
	>
		<!-- <div class={"flex flex-row absolute sm:top-auto sm:left-0"}> -->
		{@render children?.()}
	</div>
	<div class="border-zinc-200 border-t"></div>
</nav>

<style>
	.max-w {
		max-width: 1000px;
	}
	span {
		transition: all 0.1s;
		-webkit-transform-origin: center center;
		-moz-transform-origin: center center;
		-o-transform-origin: center center;
		transform-origin: center center;
	}
</style>
