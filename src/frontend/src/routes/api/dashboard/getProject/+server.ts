import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { error } from '@sveltejs/kit';
import { fetchFromBackend } from '../../../../util/backendService';

export const POST = (async ({ request }) => {
    const data = await request.json();
    const response = await fetchFromBackend(request, 'dashboard/project', data);

    if(response.status === 401) {
        throw error(401, "Your session has expired. Please sign in again.");
    }

	return new Response(response.body);
}) satisfies RequestHandler;