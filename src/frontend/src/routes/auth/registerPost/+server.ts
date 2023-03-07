import { error, json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { fetchFromBackend } from '../../../util/backendService';

export const POST = (async ({ request }) => {
	await fetchFromBackend('register', await request.json());

	return json({
		status: 200
	});
}) satisfies RequestHandler;