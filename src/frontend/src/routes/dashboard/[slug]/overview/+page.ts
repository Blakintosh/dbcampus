import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';
 
export const load = (({ params }) => {
 
  throw redirect(301, `/dashboard/${params.slug}/`);
}) satisfies PageLoad;