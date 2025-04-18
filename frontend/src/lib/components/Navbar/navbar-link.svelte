<!-- 
 TODO: add functionality to add sub links
 TODO: show current active link (and expand if sub links)
-->
<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';

	import { burgerOpen } from '$s/Navbar';

	let { label, link = '' } = $props();

	const navigate = (e: any) => {
		e.preventDefault();
		goto(link);
		$burgerOpen = !$burgerOpen;
	};
</script>

<a
	href={link}
	class={'transistion-all px-3 py-2 duration-100 sm:my-auto sm:rounded ' +
		(page.url.pathname === link ? 'underline' : '')}
	onclick={navigate}
>
	{label}
</a>
