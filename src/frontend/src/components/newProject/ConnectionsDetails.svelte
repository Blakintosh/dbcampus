<script lang="ts">
	import FormField from "../form/FormField.svelte";
	import FormTextField from "../form/FormTextField.svelte";
</script>

<h1 class="text-xl">Connections</h1>

<p class="text-sm">
	This step is optional. You can either skip it, or fill out all the fields below.
</p>
<p class="text-sm">
	<span class="text-red-600 font-semibold">*</span> = Mandatory
</p>

<FormTextField label="Enter the URL to your Jira board." placeholder="Project Name" required name="projectName"/>

<FormTextField label="Enter your jira API key." placeholder="API Key" required name="apiKey"/>
