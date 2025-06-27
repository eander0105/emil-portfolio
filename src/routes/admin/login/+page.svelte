<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';

	import { type Infer, type SuperValidated, superForm } from 'sveltekit-superforms';
	import { type LoginSchema, loginSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	let { data }: { data: { form: SuperValidated<Infer<LoginSchema>> } } = $props();

	const form = superForm(data.form, {
		validators: zodClient(loginSchema),
	});

	const { form: formData, enhance } = form;
</script>

<div class="h-dvh">
	<Card.Root class="mx-4 my-8 md:mx-auto md:w-[400px]">
		<Card.Header>
			<Card.Title class="text-center">Admin login</Card.Title>
		</Card.Header>
		<Card.Content>
			<form use:enhance>
				<Form.Field {form} name="email">
					<Form.Control>
						{#snippet children({ props })}
							<Form.Label>Email</Form.Label>
							<Input type="email" {...props} bind:value={$formData.email} />
						{/snippet}
					</Form.Control>
				</Form.Field>
				<Form.Field {form} name="password">
					<Form.Control>
						{#snippet children({ props })}
							<Form.Label>Password</Form.Label>
							<Input type="password" {...props} bind:value={$formData.password} />
						{/snippet}
					</Form.Control>
				</Form.Field>
				<!-- <Form.Field {form} name="password">
					<Form.Control>
						<Form.Label>Password</Form.Label>
					</Form.Control>
				</Form.Field> -->
				<!-- <div class="grid gap-4">
					<div class="flex flex-col gap-2">
						<Label for="email">Email</Label>
						<Input id="email" type="email" />
					</div>
					<div class="flex flex-col gap-2">
						<Label for="password">Password</Label>
						<Input id="password" type="password" />
					</div>
				</div> -->
			</form>
		</Card.Content>
		<Card.Footer>
			<Button class="w-full">Login</Button>
		</Card.Footer>
	</Card.Root>
</div>
