<!-- TODO: look through this file and improve it =) -->
<script lang="ts">
	let fileInput: HTMLInputElement;
	export let endpoint: string = '/api/upload';
	let uploadResult: string = '';

	async function uploadFile() {
		const file = fileInput?.files?.[0];
		if (!file) {
			uploadResult = 'Please select a file.';
			return;
		}

		const formData = new FormData();
		formData.append('file', file);

		const res = await fetch(endpoint, {
			method: 'POST',
			body: formData
		});

		const data: { success?: boolean; path?: string; error?: string } = await res.json();
		uploadResult = data.success
			? `File uploaded: ${data.path}`
			: `Error: ${data.error}`;
	}
</script>

<div class="max-w-md mx-auto mt-6 p-6 bg-white rounded-lg shadow">
	<form class="flex flex-col gap-4" on:submit|preventDefault={uploadFile}>
		<input type="file" bind:this={fileInput} class="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded file:border-0 file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100" />
		<button type="submit" class="bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700 transition font-semibold">
			Upload
		</button>
	</form>
	{#if uploadResult}
		<div class="mt-4 text-gray-700">{uploadResult}</div>
	{/if}
</div>
