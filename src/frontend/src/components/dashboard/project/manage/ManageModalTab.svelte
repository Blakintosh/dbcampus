<script lang="ts">
    export let label: string;
</script>

<button class="bg-slate-800 text-slate-100 py-2 w-full text-left font-normal text-sm hover:bg-slate-700/50">
    <li class="flex items-center gap-4">
        <span class="fill-amber-200">
            <slot/>
        </span>
        <span>
            {label}
        </span>
    </li>
</button>