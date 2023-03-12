<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import ButtonRoot from './ButtonRoot.svelte';

	export let type: string = 'primary';
	export let label: string = 'New';
    export let additionalClasses: string = '';

	export let loading: boolean = true;
    export let disabled: boolean = false;

	const dispatch = createEventDispatcher();
	const handleClick = () => {
        if(!disabled) {
		    dispatch('click');
        }
	};
</script>

<ButtonRoot
	{type}
	buttonClass="duration-75 rounded-md text-sm md:text-md font-medium py-1.5 text-slate-100 shadow-md flex items-center gap-2 justify-center {additionalClasses} {loading ? "opacity-80 cursor-wait" : ""} {disabled ? "cursor-not-allowed" : ""}"
	disabled={loading || disabled}
	on:click={handleClick}
>
	{#if loading}
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="fill-slate-100 w-4 h-4 animate-spin">
			<!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. -->
			<path
				d="M304 48c0-26.5-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48s48-21.5 48-48zm0 416c0-26.5-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48s48-21.5 48-48zM48 304c26.5 0 48-21.5 48-48s-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48zm464-48c0-26.5-21.5-48-48-48s-48 21.5-48 48s21.5 48 48 48s48-21.5 48-48zM142.9 437c18.7-18.7 18.7-49.1 0-67.9s-49.1-18.7-67.9 0s-18.7 49.1 0 67.9s49.1 18.7 67.9 0zm0-294.2c18.7-18.7 18.7-49.1 0-67.9S93.7 56.2 75 75s-18.7 49.1 0 67.9s49.1 18.7 67.9 0zM369.1 437c18.7 18.7 49.1 18.7 67.9 0s18.7-49.1 0-67.9s-49.1-18.7-67.9 0s-18.7 49.1 0 67.9z"
			/>
		</svg>
	{/if}
	<span />
	{label}
</ButtonRoot>
