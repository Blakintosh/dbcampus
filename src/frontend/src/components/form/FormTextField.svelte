<script lang="ts">
	import FormField from "./FormField.svelte";

    export let label: string;
    export let required: boolean = false;
    export let name: string;
    export let placeholder: string;
</script>

<FormField {label} {required} isFor={name}>
    <input type="text" {name} {placeholder} class="w-full border border-slate-300 rounded-lg p-2 mt-2"/>
</FormField>