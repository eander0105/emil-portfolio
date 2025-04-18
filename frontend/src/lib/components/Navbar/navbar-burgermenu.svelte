<script>
	import { burgerOpen } from '$s/Navbar';
</script>

<!-- Burger menu -->
<button
	class="group z-50 my-auto flex h-8 flex-col justify-evenly sm:hidden"
	aria-label="Toggle menu"
	onclick={() => ($burgerOpen = !$burgerOpen)}
>
	<span
		id="s1"
		class={'translate h-0.5 w-8 rounded-xl ' +
			($burgerOpen ? 'translate-y-[0.54rem] rotate-45' : '') +
			' bg-black dark:bg-white'}
	>
		&nbsp;
	</span>
	<span
		id="s2"
		class={'h-0.5 w-8 rounded-xl ' + ($burgerOpen ? 'scale-50 opacity-0' : '') + 
		' bg-black dark:bg-white'}
	>
		&nbsp;
	</span>
	<span
		id="s3"
		class={'translate h-0.5 w-8 rounded-xl ' +
			($burgerOpen ? '-translate-y-[0.54rem] -rotate-45' : '') +
			' bg-black dark:bg-white'}
	>
		&nbsp;
	</span>
</button>
