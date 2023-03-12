<script lang="ts">
    export let label: string;
    export let required: boolean = false;
    export let isFor: string;
    export let error: string = "";
</script>

<label for={isFor} class="block mt-4 text-sm">{label}<span class="text-red-600 font-semibold" class:hidden={!required}>*</span></label>
<slot/>
{#if error}
    <p class="text-red-600 text-xs mt-1">{error}</p>
{/if}