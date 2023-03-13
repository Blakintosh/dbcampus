<script lang="ts">
	import FormDivider from "./FormDivider.svelte";
    import FormStage from "./FormStage.svelte";

    export let stage: number;
    export let stages: Array<string>;
</script>

{#each stages as stageName, index}
    <FormStage stage={index + 1} complete={stage > index + 1} label={stageName} current={stage === index + 1}/>
    {#if index < stages.length - 1}
        <FormDivider complete={stage > index + 1}/>
    {/if}
{/each}
