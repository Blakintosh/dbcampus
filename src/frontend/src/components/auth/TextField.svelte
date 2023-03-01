<script lang="ts">
	export let label: string;
    export let name: string;
    export let value: string = "";
    export let isError: boolean = false;
    export let errorMessage: string = "";

    const resetErrorStatus = () => {
        isError = false;
        errorMessage = "";
    }
</script>

<div class="flex flex-col w-full mt-4">
    <label for={name} class="text-sm font-medium flex items-center gap-2" class:text-slate-600={!isError} class:text-red-600={isError}>
        {#if isError}
            <span class="fill-red-600">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="w-4 h-4"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M256 32c14.2 0 27.3 7.5 34.5 19.8l216 368c7.3 12.4 7.3 27.7 .2 40.1S486.3 480 472 480H40c-14.3 0-27.6-7.7-34.7-20.1s-7-27.8 .2-40.1l216-368C228.7 39.5 241.8 32 256 32zm0 128c-13.3 0-24 10.7-24 24V296c0 13.3 10.7 24 24 24s24-10.7 24-24V184c0-13.3-10.7-24-24-24zm32 224a32 32 0 1 0 -64 0 32 32 0 1 0 64 0z"/></svg>
            </span>
        {/if}
        {label}
    </label>
    <input type="text" id={name} bind:value={value} class="border rounded-md p-2 mt-1 w-full {isError ? "bg-red-50 border-red-600" : "border-slate-300"}" on:change={resetErrorStatus} />
    <div class="text-xs text-red-600 py-2 h-1">
        {errorMessage}
    </div>
</div>