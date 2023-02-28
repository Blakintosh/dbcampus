<script lang="ts">
	export let type: string = "primary";
	export let label: string = "New";

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
				typeColour = "bg-slate-600";
				hoverColour = "hover:bg-slate-700";
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
</script>

<button class="{typeColour} {hoverColour} duration-75 rounded-sm text-xs font-medium px-3 py-1.5 text-slate-100 mx-4 shadow-md">
	{label}
</button>