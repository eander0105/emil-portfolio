import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	// Check if the user is authenticated
	// TODO: Implement real authentication check
	const isAuthenticated = !!event.cookies.get('auth_token');

	const { pathname } = event.url;

	// Redirect unauthenticated users trying to access /admin paths
	if (pathname.startsWith('/admin')) {
		if (!isAuthenticated) {
			if (pathname === '/admin/login') {
				// Allow access to the login page
				return resolve(event);
			}

			return new Response(null, {
				status: 302,
				headers: { Location: '/admin/login' }
			});
		}

		if (pathname === '/admin/login') {
			return new Response(null, {
				status: 302,
				headers: { Location: '/admin' }
			});
		}
	}

	// Proceed with the request
	return resolve(event);
};
