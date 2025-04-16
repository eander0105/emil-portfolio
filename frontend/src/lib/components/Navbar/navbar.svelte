<script lang="ts">
	import { burgerOpen } from '$s/Navbar';
	import BurgerMenu from './navbar-burgermenu.svelte';
	import Icon from '@iconify/svelte';

	let { children } = $props();
</script>

<nav
	class="container sticky top-0 mx-auto flex flex-row justify-between bg-zinc-900 px-4 py-6 text-white"
>
	<!-- Header - Name, subtitle & Image -->
	<div class="h- z-40 flex flex-row items-center">
		<img
			src="https://github.com/eander0105.png"
			alt="Emil Andersson"
			class="h-16 w-16 rounded sm:h-20 sm:w-20"
		/>
		<div class="ml-4 flex flex-col">
			<a href="/" class="text-3xl font-bold md:text-4xl">Emil Andersson</a>
			<!-- TODO: select and style a nice subtitle -->
			<!-- <span class="text-md md:text-xl text-zinc-400">Developer with a few to many interests</span> -->
			<span class="text-md text-zinc-400 md:text-xl">Software developer</span>
		</div>
	</div>

	<!-- Desktop links -->
	<div class="left-auto hidden h-auto w-auto flex-row p-0 sm:flex">
		{@render children?.()}
	</div>

	<!-- TODO: think about a better looking way to display links on smaller screens (phones) -->
	<!-- Mobile menu -->
	<BurgerMenu />
	<div
		class={'links ctransition-all absolute left-0 flex h-screen w-full duration-200 sm:hidden ' +
			($burgerOpen
				? 'visible translate-y-0 bg-zinc-900 opacity-100 ease-in-out'
				: 'invisible -translate-y-4 opacity-0')}
	>
		<div class="mt-20 flex w-full flex-col px-4">
			{@render children?.()}
		</div>

		<!-- Bottom links (mobile only) -->
		<div class="absolute bottom-6 w-full px-4 sm:hidden">
			<hr class="my-2 border-gray-500" />
			<div class="flex w-full flex-row justify-around pb-4 pt-2">
				<a
					class="flex flex-row"
					href="https://github.com/eander0105"
					target="_blank"
					title="GitHub"
				>
					<Icon icon="mdi:github" height="20" class="mr-2" /> Github
				</a>
				<a
					class="flex flex-row"
					href="https://www.linkedin.com/in/emil-andersson-846b22203/"
					target="_blank"
					title="LinkedIn"
				>
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
