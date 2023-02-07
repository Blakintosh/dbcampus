<script lang="ts">
	export let status: string = "unknown";
</script>

<div class=" bg-slate-800 rounded-lg border-slate-700 border inline-block">
	<div class="rounded-full w-12 h-12 border-slate-900 border bg-black m-2" class:bg-red-600="{status === 'danger'}"></div>
	<div class="rounded-full w-12 h-12 border-slate-900 border bg-black m-2" class:bg-amber-500="{status === 'warning' || status === 'unknown'}" class:animate-blinking="{status === 'unknown'}"></div>
	<div class="rounded-full w-12 h-12 border-slate-900 border bg-black m-2" class:bg-green-500="{status === 'normal'}"></div>
</div>