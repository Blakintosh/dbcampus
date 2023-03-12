import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { error } from '@sveltejs/kit';
import { fetchFromBackend } from '../../../../util/backendService';

export const POST = (async ({ request }) => {
    const response = await fetchFromBackend(request, 'newProject', await request.json());

	return new Response(null, {});
}) satisfies RequestHandler;