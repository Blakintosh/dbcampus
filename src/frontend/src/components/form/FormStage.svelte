<script lang="ts">
    export let complete: boolean;
    export let stage: number;
    export let label: string;
</script>

<div class="relative">
    <div class="absolute w-64 -top-8 text-sm font-medium left-[50%] translate-x-[-50%] text-center">
        {label}
    </div>
    {#if complete}
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="w-6 h-6 -mx-1 z-10 fill-green-500"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M256 512c141.4 0 256-114.6 256-256S397.4 0 256 0S0 114.6 0 256S114.6 512 256 512zM369 209L241 337c-9.4 9.4-24.6 9.4-33.9 0l-64-64c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l47 47L335 175c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9z"/></svg>
    {:else}
        <div class="w-6 h-6 -mx-1 z-10 fill-white rounded-full flex items-center justify-center bg-neutral-500 font-semibold text-white text-lg">
            {stage}
        </div>
    {/if}
</div>