<script lang="ts">
    export let title: string;
    export let description: string = "";
</script>

<div class="my-4">
    <h3 class="text-lg font-semibold">{title}</h3>
    <p class="mb-1">{description}</p>

    <slot/>
</div>