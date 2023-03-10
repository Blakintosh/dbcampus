import type { Handle } from "@sveltejs/kit";
import { SvelteKitAuth } from "sk-auth";

export const handle: Handle = async ({ event, resolve }) => {
	// TODO https://github.com/sveltejs/kit/issues/1046

	//   if (event.request.query.has("_method")) {
	//     event.request.method = event.request.query.get("_method").toUpperCase();
	//   }

	const response = await resolve(event);

	return response;
};

export const { getSession } = new SvelteKitAuth({
	providers: [
	],
	callbacks: {
		jwt(token, profile) {
			if (profile?.provider) {
				const { provider, ...account } = profile;
				token = {
					...token,
					user: {
						...(token.user ?? {}),
						connections: { ...(token.user?.connections ?? {}), [provider]: account },
					},
				};
			}

			return token;
		},
	},
	jwtSecret: import.meta.env.VITE_JWT_SECRET_KEY,
});
