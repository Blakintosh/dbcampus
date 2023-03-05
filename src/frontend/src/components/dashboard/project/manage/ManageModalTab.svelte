<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let label: string;
    export let selected: boolean = false;

    const dispatch = createEventDispatcher();
    const handleClick = () => {
        dispatch("click");
    }
</script>

<button class="bg-slate-800 text-slate-100 py-2 w-full text-left text-sm hover:bg-slate-700/50 {selected ? "bg-slate-700/20 font-semibold" : "font-normal"}" on:click={handleClick}>
    <li class="flex items-center gap-4 relative">
        <div class="absolute -left-4 top-[50%] translate-y-[-50%] h-6 w-1 {selected ? "block" : "hidden"} rounded-sm bg-amber-300/90">
        </div>
        <span class="{selected ? "fill-amber-300" : "fill-amber-200"}">
            <slot/>
        </span>
        <span>
            {label}
        </span>
    </li>
</button>