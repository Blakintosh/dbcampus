<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let type: string = "primary";
    export let buttonClass: string = "";
    export let disabled: boolean = false;

	// Map the button type to a colour
	let typeColour = "bg-slate-500";
	let hoverColour = "hover:bg-slate-600";
	$: {
		switch(type)
		{
			case "primary":
				typeColour = "bg-slate-500";
				hoverColour = "hover:bg-slate-600";
				break;
			case "secondary":
				typeColour = "bg-slate-400";
				hoverColour = "hover:bg-slate-500";
				break;
			case "danger":
				typeColour = "bg-red-500";
				hoverColour = "hover:bg-red-600";
				break;
			case "warning":
				typeColour = "bg-amber-500";
				hoverColour = "hover:bg-amber-600";
				break;
			case "success":
				typeColour = "bg-green-500";
				hoverColour = "hover:bg-green-600";
				break;
			default:
				typeColour = "bg-slate-500";
				hoverColour = "hover:bg-slate-600";
				break;
		}
	}

	const dispatch = createEventDispatcher();
	const handleClick = () => {
        if(disabled) {
            return;
        }
		dispatch("click");
	}
</script>

<button class="{typeColour} duration-75 {buttonClass} {disabled ? "opacity-80" : hoverColour}" on:click={handleClick}>
	<slot/>
</button>