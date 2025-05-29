<!-- 
 TODO: add functionality to add sub links
 TODO: better SEO for links
-->
<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';

	let { label, link = '' } = $props();

	const navigate = (e: any) => {
		e.preventDefault();
		goto(link);
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
