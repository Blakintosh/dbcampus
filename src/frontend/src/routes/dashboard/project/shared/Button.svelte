<script lang="ts">
	export let type: string = "primary";
	export let label: string = "New";

	// Map the button type to a colour
	let typeColour = "bg-slate-500";
	$: {
		switch(type)
		{
			case "primary":
				typeColour = "bg-slate-500";
				break;
			case "secondary":
				typeColour = "bg-slate-600";
				break;
			case "danger":
				typeColour = "bg-red-500";
				break;
			case "warning":
				typeColour = "bg-amber-500";
				break;
			case "success":
				typeColour = "bg-green-500";
				break;
			default:
				typeColour = "bg-slate-500";
				break;
		}
	}
</script>

<button class="{typeColour} rounded-sm text-sm px-2 py-1 text-slate-100 mx-4 shadow-sm">
	{label}
</button>