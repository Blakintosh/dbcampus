<script lang="ts">
	import Tile from "./Tile.svelte";

    export let projectName: string = "";
    export let categoryName: string = "";
    export let tileClass: string = "";
</script>

<Tile {tileClass}>
    <div class="flex items-center gap-4 w-full mb-2">
        <span class="w-6 fill-amber-400">
            <slot name="categoryIcon"/>
        </span>
        <span class="text-lg font-medium text-red-600">
            {categoryName}
        </span>
    </div>
    <h1 class="text-3xl font-serif mb-4">{projectName}</h1>
		
    <div class="flex justify-between items-center">
        <p class="text-md font-semibold mx-2 uppercase">
            Actions
        </p>
        <p class="flex-grow flex justify-center">
            <slot name="actionButtons"/>
        </p>
    </div>
</Tile>