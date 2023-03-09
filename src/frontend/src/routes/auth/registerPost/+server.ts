import { error, json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { fetchFromBackend } from '../../../util/backendService';

export const POST = (async ({ request }) => {
	await fetchFromBackend(request, 'register', await request.json());

	return new Response();
}) satisfies RequestHandler;