import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	// Check if the user is authenticated
	const isAuthenticated = !!event.cookies.get('auth_token'); // Simplified boolean conversion

	const { pathname } = event.url;

	// Redirect unauthenticated users trying to access /admin paths
	if (pathname.startsWith('/admin') && pathname !== '/admin/login' && !isAuthenticated) {
		return new Response(null, {
			status: 302,
			headers: { Location: '/admin/login' }
		});
	}

	// Proceed with the request
	return resolve(event);
};
