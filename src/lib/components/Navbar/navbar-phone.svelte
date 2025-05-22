<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Menu } from '@lucide/svelte'
	import { Separator } from '$c/ui/separator';
	import { Button } from '$c/ui/button'
	import * as Sheet from '$c/ui/sheet'

	let { children } = $props();
</script>

<Sheet.Root>
	<Sheet.Trigger>
		<Button size="icon">
			<Menu />
		</Button>
	</Sheet.Trigger>
	<Sheet.Content side="right" class="flex flex-col h-full justify-between">
		<Sheet.Header class="text-left">
			<Sheet.Title>
				Emil Andersson
			</Sheet.Title>
			<Sheet.Description>
				Software developer
			</Sheet.Description>
			<Separator />
		</Sheet.Header>
		<div class="flex flex-col h-full w-full my-4">
			{@render children?.()}
		</div>
		<Sheet.Footer>
			<div> <!-- Wrapper need for some reason -->
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
<!-- <div
	class={'links ctransition-all absolute left-0 flex h-screen w-full duration-200 sm:hidden ' +
		($burgerOpen
			? 'visible translate-y-0 opacity-100 ease-in-out'
			: 'invisible -translate-y-4 opacity-0')}
>
	<div class="mt-20 flex w-full flex-col px-4">
		{@render children?.()}
	</div>

	<div class="absolute bottom-6 w-full px-4 sm:hidden">

		<Separator class="my-2"/>

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
</div> -->
