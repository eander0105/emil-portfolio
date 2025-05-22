<script lang="ts">
	import PhoneNav from './navbar-phone.svelte'

	let { children } = $props();
</script>

<nav
	class="container sticky top-0 z-50 mx-auto flex flex-row justify-between px-4 py-6"
>
	<!-- Header - Name, subtitle & Image -->
	<div class="h- z-40 flex flex-row items-center">
		<img
			src="/profile.jpeg"
			alt="Emil Andersson"
			class="h-16 w-16 rounded sm:h-20 sm:w-20"
		/>
		<div class="ml-4 flex flex-col">
			<a href="/" class="text-3xl font-bold md:text-4xl">Emil Andersson</a>
			<!-- TODO: select and style a nice subtitle -->
			<!-- <span class="text-md md:text-xl text-zinc-400">Developer with a few to many interests</span> -->
			<span class="text-md text-muted-foreground md:text-xl">Software developer</span>
		</div>
	</div>

	<!-- Desktop links -->
	<div class="left-auto hidden h-auto w-auto flex-row p-0 sm:flex">
		{@render children?.()}
	</div>

	<!-- Mobile links/menu -->
	<div class="md:hidden">
		<PhoneNav>
			{@render children?.()}
		</PhoneNav>
	</div>
</nav>

<style>
	span {
		transition: all 0.1s;
		-webkit-transform-origin: center center;
		-moz-transform-origin: center center;
		-o-transform-origin: center center;
		transform-origin: center center;
	}
</style>
