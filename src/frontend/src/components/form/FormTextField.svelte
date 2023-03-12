<script lang="ts">
	import FormField from "./FormField.svelte";

    export let label: string;
    export let required: boolean = false;
    export let name: string;
    export let placeholder: string;
    export let value: string;
    export let error: string = "";
</script>

<FormField {label} {required} isFor={name} {error}>
    <input type="text" {name} {placeholder} class="w-full border duration-75 {error ? "border-red-400 bg-red-100/80" : "border-slate-300"} rounded-lg p-2 mt-2" bind:value={value}/>
</FormField>