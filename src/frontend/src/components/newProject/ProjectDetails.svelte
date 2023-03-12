<script lang="ts">
	import FormField from "../form/FormField.svelte";
	import FormTextField from "../form/FormTextField.svelte";
</script>

<h1 class="text-xl">Basic Details</h1>

<p class="text-sm">
	If you don't know what to put in here, don't worry. You can always change it later.
</p>
<p class="text-sm">
	<span class="text-red-600 font-semibold">*</span> = Mandatory
</p>

<FormTextField label="What is the project called?" placeholder="Project Name" required name="projectName"/>

<FormField label="When is the project set to be finished?" isFor="finishDate" required>
	<input type="date" name="finishDate" class="w-full border border-slate-300 rounded-lg p-2 mt-2"/>
</FormField>