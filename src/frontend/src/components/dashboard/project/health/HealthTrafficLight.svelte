<script lang="ts">
	export let status: string = "unknown";
	export let horizontal: boolean = false;

    export let sizeClass: string = "w-10 h-10 m-1";
</script>

<div class=" bg-slate-800 rounded-lg border-slate-700 border inline-flex p-1 {horizontal ? "" : "flex-col"}">
	<div class="rounded-full {sizeClass} border-slate-900 border" class:bg-red-600="{status === 'danger'}" class:bg-black="{status !== 'danger'}"></div>
	<div class="rounded-full {sizeClass} border-slate-900 border" class:bg-amber-500="{status === 'warning'}" class:animate-blinking="{status === 'unknown'}" class:bg-black="{status !== 'warning' && status !== 'unknown'}"></div>
	<div class="rounded-full {sizeClass} border-slate-900 border" class:bg-green-500="{status === 'normal'}" class:bg-black="{status !== 'normal'}"></div>
</div>