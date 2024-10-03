<script>
	import '../app.css';
	import {
		Navbar,
		NavBrand,
		NavLi,
		NavUl,
		NavHamburger,
		Footer,
		FooterIcon
	} from 'flowbite-svelte';
	import { LinkedinSolid, GithubSolid, DiscordSolid, DribbbleSolid, EnvelopeSolid } from 'flowbite-svelte-icons';
</script>

<main class="dark flex min-h-screen flex-col bg-white dark:bg-gray-800">
	<div class="relative px-8 pt-20">
		<Navbar class="fixed start-0 top-0 z-20 w-full border-b px-2 py-2.5 sm:px-4">
			<NavBrand href="/">
				<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
					>Flowbite</span
				>
			</NavBrand>
			<NavHamburger />
			<NavUl>
				<NavLi href="/">Home</NavLi>
				<NavLi href="/about">Blog</NavLi>
				<NavLi href="/photography">Photography</NavLi>
				<NavLi href="/project">Projects</NavLi>
				<NavLi href="/docs/components/navbar">Resume</NavLi>
			</NavUl>
		</Navbar>
	</div>

	<slot class="flex-grow" />

	<Footer footerType="socialmedia" class="dark:bg-gray-900 mt-auto">
		<!-- <hr class="my-6 border-gray-200 sm:mx-auto lg:my-8 dark:border-gray-700" /> -->
		<div class="sm:flex sm:items-center sm:justify-between">
			<div class="mt-4 flex space-x-6 sm:mt-0 sm:justify-center rtl:space-x-reverse">
				<FooterIcon href="https://www.linkedin.com/in/emil-andersson-846b22203/" target="_blank">
					<LinkedinSolid
						class="h-5 w-5 text-gray-500 hover:text-gray-900 dark:text-gray-500 dark:hover:text-white"
					/>
				</FooterIcon>

				<FooterIcon href="httops://github.com/eander0105" target="_blank">
					<GithubSolid
					class="h-5 w-5 text-gray-500 hover:text-gray-900 dark:text-gray-500 dark:hover:text-white"
					/>
				</FooterIcon>

				<FooterIcon href="emilandersson.se" target="_blank">
					<DribbbleSolid />
				</FooterIcon>
				
				<FooterIcon id="discord-contact">
					<DiscordSolid
						class="h-5 w-5 text-gray-500 hover:text-gray-900 dark:text-gray-500 dark:hover:text-white"
					/>
				</FooterIcon>

				<FooterIcon id="email-contact">
					<EnvelopeSolid
						class="h-5 w-5 text-gray-500 hover:text-gray-900 dark:text-gray-500 dark:hover:text-white"
					/>
				</FooterIcon>
			</div>
		</div>
	</Footer>
</main>
