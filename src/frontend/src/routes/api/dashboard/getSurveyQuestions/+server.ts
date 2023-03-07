import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { error } from '@sveltejs/kit';
import { fetchFromBackend } from '../../../../util/backendService';

export const GET = (async ({ url }) => {
    try {
		const destination = url.searchParams.get("type") === "team" ? "getTeamSurveyQuestions" : "getClientSurveyQuestions";
        const response = await fetchFromBackend(destination);
    
        return {
            status: 200,
            body: await response.json()
        }
    } catch(e) {
        throw error(500, "The request to the backend failed to resolve.");
    }
}) satisfies RequestHandler;