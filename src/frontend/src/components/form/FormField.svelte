<script lang="ts">
    export let label: string;
    export let required: boolean = false;
    export let isFor: string;
</script>

<label for={isFor} class="block mt-4 text-sm">{label}<span class="text-red-600 font-semibold" class:hidden={!required}>*</span></label>
<slot/>