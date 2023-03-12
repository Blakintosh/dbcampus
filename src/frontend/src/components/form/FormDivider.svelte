<script lang="ts">
    export let complete: boolean;
</script>

<div class="w-20 lg:w-32 h-1 bg-neutral-500">
    <div class="h-1 duration-1000 bg-green-500 {complete ? "w-full" : "w-0"}"></div>
</div>