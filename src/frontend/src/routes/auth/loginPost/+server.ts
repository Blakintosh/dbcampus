import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { fetchFromBackend } from '../../../util/backendService';

export const POST = (async ({ request }) => {
    const response = await fetchFromBackend('login', await request.json());

    console.log(response);

    return json({
        status: 200,
        data: await response.json()
    });
}) satisfies RequestHandler;