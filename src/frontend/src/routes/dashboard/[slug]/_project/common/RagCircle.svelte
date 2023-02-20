<script lang="ts">
	export let status: string = "unknown";
</script>

<div class="rounded-full h-3 w-3 border" class:bg-red-600="{status === 'danger'}" class:bg-amber-500="{status === 'warning'}" class:bg-green-500="{status === 'normal'}" class:bg-gray-300="{status === 'unknown'}">
</div>