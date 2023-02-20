<script lang="ts">
	export let status: string = "unknown";
</script>

<div class=" bg-slate-800 rounded-lg border-slate-700 border inline-block">
	<div class="rounded-full w-12 h-12 border-slate-900 border m-2" class:bg-red-600="{status === 'danger'}" class:bg-black="{status !== 'danger'}"></div>
	<div class="rounded-full w-12 h-12 border-slate-900 border m-2" class:bg-amber-500="{status === 'warning'}" class:animate-blinking="{status === 'unknown'}" class:bg-black="{status !== 'warning' && status !== 'unknown'}"></div>
	<div class="rounded-full w-12 h-12 border-slate-900 border m-2" class:bg-green-500="{status === 'normal'}" class:bg-black="{status !== 'normal'}"></div>
</div>