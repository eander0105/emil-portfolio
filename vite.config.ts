import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig(({command}) => {
	const production = command === 'build';
	
	// Allow *.emilandersson.se during development
	const allowedHosts = production ? ['emilandersson.se'] : ['.emilandersson.se'];

	return {
		plugins: [sveltekit()],
		server: {
			host: '0.0.0.0',
			port: 5173,
			allowedHosts,
		}
	};
});
