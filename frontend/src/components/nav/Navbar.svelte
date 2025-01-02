<script lang="ts">
	import { burgerOpen } from "$l/client/Navbar";
	let { children } = $props();
</script>

<nav class={'sticky top-0 flex flex-row justify-between bg-zinc-900 px-8 py-6 text-white'}>
	<span>
		<a href="/" class="text-xl font-bold">Emil Andersson</a>
	</span>

	<!-- Burger menu -->
	<button
		class="group z-50 flex h-8 flex-col justify-evenly sm:hidden"
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

	<div
		class={'links absolute left-0 top-20 flex h-screen w-full flex-col p-5 transition-all duration-200 sm:static sm:left-auto sm:top-auto sm:h-auto sm:w-auto sm:flex-row sm:p-0 ' +
			($burgerOpen
				? 'visible translate-y-0 bg-zinc-900 opacity-100 ease-in-out'
				: 'invisible -translate-y-4 opacity-0 sm:visible sm:translate-y-0 sm:opacity-100')}
	>
		<!-- <div class={"flex flex-row absolute sm:top-auto sm:left-0"}> -->
		{@render children?.()}
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
