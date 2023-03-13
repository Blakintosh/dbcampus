<script lang="ts">
    export let section: string;
    export let title: string;
</script>

<span class="text-lg font-medium flex items-center gap-3 text-slate-900">
    <img src="/pitstop.png" alt="Back" class="h-5 inline-block" />
    <span>
        <span class="text-red-600">
            {section}
        </span>
        <span class="mx-1 text-sm">
            /
        </span>
        <span class="text-sm">
            {title}
        </span>
    </span>
</span>