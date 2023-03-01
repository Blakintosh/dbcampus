<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import ButtonRoot from './ButtonRoot.svelte';

	export let type: string = "primary";
	export let label: string = "New";

	const dispatch = createEventDispatcher();
	const handleClick = () => {
		dispatch("click");
	}
</script>

<ButtonRoot type={type} buttonClass="duration-75 rounded-sm text-xs font-medium px-3 py-1.5 text-slate-100 mx-4 shadow-md" on:click={handleClick}>
	{label}
</ButtonRoot>