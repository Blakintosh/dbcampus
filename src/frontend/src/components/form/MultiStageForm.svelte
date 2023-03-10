<script lang="ts">
	import FormStage from "./FormStage.svelte";

    export let stage: number;
    export let stages: Array<string>;
</script>

{#each stages as stageName, index}
    <FormStage stage={index + 1} complete={stage > index + 1} label={stageName}/>
    {#if index < stages.length - 1}
        <div class="w-32 h-1 bg-neutral-500"></div>
    {/if}
{/each}
