import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { error } from '@sveltejs/kit';
import { fetchFromBackend } from '../../../../util/backendService';

export const POST = (async ({ request }) => {
    const response = await fetchFromBackend(request, 'newProject', await request.json());

    if(response.status === 401) {
        throw error(401, "Your session has expired. Please sign in again.");
    } else if(response.status === 400) {
        throw error(400, "The data you provided is not valid.");
    }

	return new Response(null, {});
}) satisfies RequestHandler;