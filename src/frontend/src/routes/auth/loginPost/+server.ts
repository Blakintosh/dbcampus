import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { error } from '@sveltejs/kit';
import { fetchFromBackend } from '../../../util/backendService';

export const POST = (async ({ request }) => {
	const cookie1 = request.headers.get('cookie');

	if(!cookie1) {
		throw error(500, "NO COOKIE");
	}
    const response = await fetchFromBackend('login', await request.json());

	const cookie = response.headers.get('set-cookie');

	if(!cookie) {
		throw error(500, "The backend did not return a cookie.");
	}

	return new Response(null, {
		headers: [
			['set-cookie', cookie]
		]
	});
}) satisfies RequestHandler;