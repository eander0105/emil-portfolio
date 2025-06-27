<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Menu } from '@lucide/svelte';
	import { Separator } from '$components/ui/separator';
	import { Button } from '$components/ui/button';
	import * as Sheet from '$components/ui/sheet';

	let { children } = $props();
</script>

<Sheet.Root>
	<Sheet.Trigger>
		<Button size="icon">
			<Menu />
		</Button>
	</Sheet.Trigger>
	<Sheet.Content side="right" class="flex h-full flex-col justify-between">
		<Sheet.Header class="text-left">
			<Sheet.Title>
				<span class="text-xl">Emil Andersson</span>
			</Sheet.Title>
			<Sheet.Description>Software developer</Sheet.Description>
			<Separator />
		</Sheet.Header>
		<div class="my-4 flex h-full w-full flex-col">
			{@render children?.()}
		</div>
		<Sheet.Footer>
			<div>
				<!-- Wrapper need for some reason -->
				<Separator />

				<div class="flex w-full flex-row justify-around pt-6">
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
		</Sheet.Footer>
	</Sheet.Content>
</Sheet.Root>
