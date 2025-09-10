<script lang="ts">
	import FileUpload from "$components/input/FileUpload.svelte";
</script>

<p>Hello!</p>

<FileUpload endpoint="/api/image/upload" />

<img src="/uploads/ekb60-finished.jpg" alt="">
