<script lang="ts">
	import PhoneNav from './navbar-phone.svelte';

	let { children } = $props();
</script>

<header class="sticky top-0 z-50 bg-background">
	<nav class="mx-auto flex flex-row justify-between border-b px-4 py-4 md:pb-4">
		<!-- Header - Name, subtitle & Image -->
		<div class="h- z-40 flex flex-row items-center">
			<img src="/profile.jpeg" alt="Emil Andersson" class="h-16 w-16 rounded sm:h-20 sm:w-20" />
			<div class="ml-4 flex flex-col">
				<a href="/" class="text-3xl font-bold md:text-4xl">Emil Andersson</a>
				<!-- TODO: select and style a nice subtitle -->
				<!-- <span class="text-md md:text-xl text-zinc-400">Developer with a few to many interests</span> -->
				<span class="text-md text-muted-foreground md:text-xl">Software developer</span>
			</div>
		</div>

		<!-- Desktop links -->
		<div class="left-auto hidden h-auto w-auto flex-row p-0 md:flex">
			{@render children?.()}
		</div>

		<!-- Mobile links/menu -->
		<div class="flex h-auto w-auto items-center justify-end md:hidden">
			<PhoneNav>
				{@render children?.()}
			</PhoneNav>
		</div>
	</nav>
</header>
