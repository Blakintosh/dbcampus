<script lang="ts">
	export let status: string = "unknown";
</script>

<div class="rounded-full h-3 w-3 border shrink-0" class:bg-red-600="{status === 'r'}" class:bg-amber-500="{status === 'a'}" class:bg-green-500="{status === 'g'}" class:bg-gray-300="{status === 'unknown'}">
</div>