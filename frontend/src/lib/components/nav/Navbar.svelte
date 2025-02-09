<script lang="ts">
	import { burgerOpen } from "$s/Navbar";
	import BurgerMenu from "./BurgerMenu.svelte";
	import Icon from "@iconify/svelte";

	let { children } = $props();
</script>

<nav class="sticky top-0 flex flex-row justify-between bg-zinc-900 px-4 py-6 text-white container mx-auto">
	<!-- Header - Name, subtitle & Image -->
	<div class="flex flex-row items-center z-40 h-">
		<img src="https://github.com/eander0105.png" alt="Emil Andersson" class="h-16 w-16 sm:h-20 sm:w-20 rounded">
		<div class="ml-4 flex flex-col">
			<a href="/" class="text-3xl md:text-4xl font-bold">Emil Andersson</a>
			<!-- TODO: select and style a nice subtitle -->
			<!-- <span class="text-md md:text-xl text-zinc-400">Developer with a few to many interests</span> -->
			<span class="text-md md:text-xl text-zinc-400">Software developer</span>
		</div>
	</div>

	<!-- Desktop links -->
	<div class="hidden sm:flex left-auto h-auto w-auto flex-row p-0">
		{@render children?.()}
	</div>

	<!-- TODO: think about a better looking way to display links on smaller screens (phones) -->
	<!-- Mobile menu -->
	<BurgerMenu />
	<div
		class={'sm:hidden links absolute left-0 flex h-screen w-full ctransition-all duration-200 ' +
			($burgerOpen
				? 'visible translate-y-0 bg-zinc-900 opacity-100 ease-in-out'
				: 'invisible -translate-y-4 opacity-0')}
	>
		<div class="px-4 w-full flex flex-col mt-20">
			{@render children?.()}
		</div>

		<!-- Bottom links (mobile only) -->
		<div class="sm:hidden w-full px-4 absolute bottom-6">
			<hr class="my-2 border-gray-500">
			<div class="flex flex-row pb-4 pt-2 w-full justify-around">
				<a class="flex flex-row" href="https://github.com/eander0105" target="_blank" title="GitHub">
					<Icon icon="mdi:github" height="20" class="mr-2" /> Github
				</a>
				<a class="flex flex-row" href="https://www.linkedin.com/in/emil-andersson-846b22203/" target="_blank" title="LinkedIn">
					<Icon icon="mdi:linkedin" height="20" class="mr-2" /> LinkedIn
				</a>
			</div>
		</div>
	</div>
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
