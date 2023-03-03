import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { error } from '@sveltejs/kit';
import { fetchFromBackend } from '../../../util/backendService';

export const POST = (async ({ request }) => {
    try {
        const response = await fetchFromBackend('login', await request.json());
    
        console.log(response);
    
        return {
            status: 200,
            headers: {
                'set-cookie': response.headers.get('set-cookie')
            },
            body: await response.json()
        }
    } catch(e) {
        throw error(500, "The request to the backend failed to resolve.");
    }
}) satisfies RequestHandler;