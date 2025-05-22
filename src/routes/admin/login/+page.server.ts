import type { Actions, PageServerLoad } from './$types'
import { superValidate } from 'sveltekit-superforms'
import { loginSchema } from './schema'
import { zod } from 'sveltekit-superforms/adapters'

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(zod(loginSchema))
	}
};

export const actions: Actions = {
	default: async (event) => {
		
	}
}
