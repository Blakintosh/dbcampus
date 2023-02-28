<script lang="ts">
	export let satisfaction: number = 0;

	let status: string = "R";
	$: {
		if (satisfaction >= 0.8) {
			status = "G";
		} else if (satisfaction >= 0.6) {
			status = "a";
		} else if (satisfaction >= 0.4) {
			status = "A";
		} else if (satisfaction >= 0.2) {
			status = "r";
		} else {
			status = "R";
		}
	}
</script>

<div class="flex items-center gap-4">
	<div class="flex bg-slate-800 rounded-sm p-0.5">
		<div class="rounded-full h-3 w-3 m-0.5" class:bg-red-600="{status === 'R'}" class:bg-black="{status !== 'R'}"></div>
		<div class="rounded-full h-3 w-3 m-0.5" class:bg-red-400="{status === 'r'}" class:bg-black="{status !== 'r'}"></div>
		<div class="rounded-full h-3 w-3 m-0.5" class:bg-amber-500="{status === 'A'}" class:bg-black="{status !== 'A'}"></div>
		<div class="rounded-full h-3 w-3 m-0.5" class:bg-amber-300="{status === 'a'}" class:bg-black="{status !== 'a'}"></div>
		<div class="rounded-full h-3 w-3 m-0.5" class:bg-green-500="{status === 'G'}" class:bg-black="{status !== 'G'}"></div>
	</div>
	<span class="font-normal">
		<span class="text-lg font-medium">{
			satisfaction.toLocaleString(undefined,{
				style: 'percent', 
				minimumFractionDigits: 1
			})
		}</span>
		satisfied
	</span>
</div>