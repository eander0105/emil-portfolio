import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	// Check if the user is authenticated
	const isAuthenticated = !!event.cookies.get('auth_token'); // Simplified boolean conversion

	const { pathname } = event.url;

	// Redirect unauthenticated users trying to access /admin paths
	if (pathname.startsWith('/admin') && !isAuthenticated) {
		if (pathname === '/admin/login') {
			// Allow access to the login page
			return resolve(event);
		}

		return new Response(null, {
			status: 302,
			headers: { Location: '/admin/login' }
		});
	}

	// Proceed with the request
	return resolve(event);
};
