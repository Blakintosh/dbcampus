import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { error } from '@sveltejs/kit';
import { fetchFromBackend } from '../../../util/backendService';

export const POST = (async ({ request }) => {
    const response = await fetchFromBackend(request, 'login', await request.json());

	const cookie = response.headers.get('set-cookie');

	const responseHeaders: HeadersInit = [];
	if(cookie) {
		responseHeaders.push(['set-cookie', cookie]);
	}

	console.log("OK");
	return new Response(null, {
		headers: responseHeaders
	});
}) satisfies RequestHandler;