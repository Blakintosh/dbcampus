<script lang="ts">
	import RagCircle from "../common/RagCircle.svelte";
	
	export let name: string;
	export let description: string;
	export let satisfaction: number;
	export let status: string;
</script>

<li class="flex justify-between w-full items-center border-b bg-slate-100 px-4">
	<div class="text-lg font-medium py-2">{name}</div>
	<div class="flex flex-row items-center">
		<div class="text-xl font-medium">{
			satisfaction.toLocaleString(undefined,{
				style: 'percent', 
				minimumFractionDigits: 1
			})
		}
		</div>
		<div class="text-sm font-normal mx-2">satisfied</div>
		<div class="text-sm">
			<RagCircle {status}/>
		</div>
	</div>
</li>