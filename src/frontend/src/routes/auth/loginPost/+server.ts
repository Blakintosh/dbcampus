import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { fetchFromBackend } from '../../../util/backendService';

const BACKEND_PORT = 8081;

const BACKEND_URL = `http://localhost:${BACKEND_PORT}`;

export const POST = (async ({ request }) => {
    const response = await fetchFromBackend('auth/login', await request.json());

    return json({
        status: 200,
        data: await response.json()
    });
}) satisfies RequestHandler;